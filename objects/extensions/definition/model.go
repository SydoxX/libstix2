package definition

import (
	"fmt"

	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/factory"

	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/properties"
	"github.com/avast/libstix2/vocabs"
)

// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_32j232tfvtly

type Definition struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty

	Schema              string                 `json:"schema"`
	Version             string                 `json:"version"`
	ExtensionTypes      []vocabs.ExtensionType `json:"extension_types"`
	ExtensionProperties []string               `json:"extension_properties,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeExtensionDefinition, func() common.STIXObject {
		return New()
	})
}

func New() *Definition {
	var obj Definition
	obj.InitSDO(objects.TypeExtensionDefinition)
	return &obj
}

func (o *Definition) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if o.Schema == "" {
		errors = append(errors, objects.PropertyMissing("schema"))
	}

	if o.Version == "" {
		errors = append(errors, objects.PropertyMissing("version"))
	}

	if len(o.ExtensionTypes) == 0 {
		errors = append(errors, objects.PropertyMissing("extension_types"))
	}

	hasTopLevelProps := false
	for _, t := range o.ExtensionTypes {
		if !t.IsValid() {
			errors = append(errors,
				objects.PropertyInvalid("extension_types", t, "the value must come from extension_types enum"))
		}
		if t == vocabs.ExtensionPropertyTopLevel {
			hasTopLevelProps = true
		}
	}

	if len(o.ExtensionProperties) != 0 && !hasTopLevelProps {
		errors = append(errors,
			objects.PropertyInvalid("extension_types", o.ExtensionProperties,
				fmt.Sprintf("extension_properties can only be used in combination with '%s' extension_type",
					vocabs.ExtensionPropertyTopLevel)))
	}

	return errors
}
