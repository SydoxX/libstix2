package properties

type ExtensionsProperty struct {
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

func (p *ExtensionsProperty) NewExtensionGeneric(name string) map[string]interface{} {
	if p.Extensions == nil {
		p.Extensions = make(map[string]interface{})
	}

	extension := map[string]interface{}{}
	p.Extensions[name] = extension
	return extension
}

func (p *ExtensionsProperty) AddExtension(name string, val interface{}) {
	if p.Extensions == nil {
		p.Extensions = make(map[string]interface{})
	}
	p.Extensions[name] = val
}
