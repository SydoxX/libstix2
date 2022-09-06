package factory

import (
	"encoding/json"
	"fmt"

	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
)

var objectCreators = map[objects.ObjectType]func() common.STIXObject{}

func RegisterObjectCreator(typ objects.ObjectType, init func() common.STIXObject) {
	if _, prs := objectCreators[typ]; prs {
		panic(fmt.Sprintf("Object type %s initializer registered twice!", typ))
	}
	objectCreators[typ] = init
}

func Create(typ objects.ObjectType) common.STIXObject {
	factory := objectCreators[typ]
	if factory == nil {
		return nil
	}
	return factory()
}

func Decode(data []byte) (common.STIXObject, error) {
	var d Decoder
	return d.Decode(data)
}

type Decoder struct {
	KeepRawData                    bool
	DecodeCustomTopLevelProperties bool
}

func (d *Decoder) Decode(data []byte) (common.STIXObject, error) {
	typ, err := common.DecodeType(data)
	if err != nil {
		return nil, err
	}

	obj := Create(typ)
	if obj == nil {
		return nil, fmt.Errorf("unimplemented STIX type: %s", typ)
	}

	if err := json.Unmarshal(data, obj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal STIX type: %s %w", typ, err)
	}

	if d.DecodeCustomTopLevelProperties {
		var genericDecode map[string]*json.RawMessage
		if err := json.Unmarshal(data, &genericDecode); err != nil {
			return nil, fmt.Errorf("failed to unmarshal STIX type: %s %w", typ, err)
		}

		props := GetJsonPropertyNames(typ)
		for _, name := range props {
			delete(genericDecode, name.Name)
		}

		if len(genericDecode) != 0 {
			obj.SetCustomProperties(genericDecode)
		}
	}

	if d.KeepRawData {
		obj.SetRawData(data)
	}
	return obj, nil
}
