package objects

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/google/uuid"
	"github.com/gowebpki/jcs"
)

// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_q6l05xzpcdf
var Stix21DeterminableNamespace = uuid.MustParse("00abedb4-aa42-466c-9c01-fed23315a9b7")

type WithIdContributingProperties interface {
	IdContributingProperties() []string
}

type IdContributingPropertyFixups interface {
	FixupIdContributingProperties(properties map[string]interface{})
}

// https://github.com/oasis-open/cti-python-stix2/blob/2743b90fc0afe45193863d3f43f69b5f95cda6a3/stix2/base.py
// Returns random ID + os.ErrNotExist if no id determining properties are present
func DetermineId(obj STIXObject) (string, error) {
	withProps, ok := obj.(WithIdContributingProperties)
	if !ok {
		panic(fmt.Sprintf("Type %T does not implement objects.WithIdContributingProperties!", obj))
	}

	encoded, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	var generalized map[string]interface{}
	if err := json.Unmarshal(encoded, &generalized); err != nil {
		return "", err
	}

	idDeterminingObject := map[string]interface{}{}
	for _, c := range withProps.IdContributingProperties() {
		val := generalized[c]
		if val != nil && !reflect.ValueOf(val).IsZero() {
			idDeterminingObject[c] = val
		}
	}

	if len(idDeterminingObject) == 0 {
		rndUuid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s--%s", obj.GetCommonProperties().ObjectType, rndUuid.String()), os.ErrNotExist
	}

	if fixup, ok := obj.(IdContributingPropertyFixups); ok {
		fixup.FixupIdContributingProperties(idDeterminingObject)
	}

	encodedDeterminingUnsorted, err := json.Marshal(idDeterminingObject)
	if err != nil {
		return "", err
	}

	encodedDeterminingCanonical, err := jcs.Transform(encodedDeterminingUnsorted)
	if err != nil {
		return "", err
	}

	objId := uuid.NewSHA1(Stix21DeterminableNamespace, encodedDeterminingCanonical)
	return fmt.Sprintf("%s--%s", obj.GetCommonProperties().ObjectType, objId.String()), nil
}
