package properties

import (
	"encoding/json"
	"fmt"

	"github.com/nextpart/libstix2/datatypes/stixid"
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/vocabs"
)

type ExtensionsProperty struct {
	Extensions map[string]any `json:"extensions,omitempty" idcontrib:"1"`
}

func (p *ExtensionsProperty) NewExtensionGeneric(name string) map[string]any {
	if p.Extensions == nil {
		p.Extensions = make(map[string]any)
	}

	extension := map[string]any{}
	p.Extensions[name] = extension
	return extension
}

func (p *ExtensionsProperty) AddExtension(name string, val any) {
	if p.Extensions == nil {
		p.Extensions = make(map[string]any)
	}
	p.Extensions[name] = val
}

func (p *ExtensionsProperty) Valid() error {
	type extensionValue struct {
		ExtensionType vocabs.ExtensionType `json:"extension_type"`
	}

	for name, val := range p.Extensions {
		typ, _, ok := stixid.SplitStixId(name)
		if !ok || typ != objects.TypeExtensionDefinition {
			continue
		}

		// the val might be anything, not just map[string]any...
		valMarshalled, err := json.Marshal(&val)
		if err != nil {
			return fmt.Errorf("failed to marshal %s value: %s", name, val)
		}

		var extVal extensionValue
		if err := json.Unmarshal(valMarshalled, &extVal); err != nil {
			return fmt.Errorf("failed to unmarshal %s value: %s", name, val)
		}

		if !extVal.ExtensionType.IsValid() {
			return fmt.Errorf("extension %s has invalid type: %s", name, extVal.ExtensionType)
		}
	}

	return nil
}
