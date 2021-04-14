package definition

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/properties"
)

// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_32j232tfvtly

type ExtensionType string

const (
	ExtensionNewSdo           ExtensionType = "new-sdo"
	ExtensionNewSco           ExtensionType = "new-sco"
	ExtensionNewSro           ExtensionType = "new-sro"
	ExtensionProperty         ExtensionType = "property-extension"
	ExtensionPropertyTopLevel ExtensionType = "toplevel-property-extension"
)

type Definition struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty

	Schema              string          `json:"schema"`
	Version             string          `json:"version"`
	ExtensionTypes      []ExtensionType `json:"extension_types"`
	ExtensionProperties []string        `json:"extension_properties,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Definition) GetPropertyList() []string {
	return []string{
		"name", "description", "schema", "version", "extension_types", "extension_properties",
	}
}

func New() *Definition {
	var obj Definition
	obj.InitSDO("extension-definition")
	return &obj
}
