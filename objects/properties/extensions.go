package properties

type ExtensionsProperty struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
}

func (p *ExtensionsProperty) NewExtension(name string) map[string]interface{} {
	extension := map[string]interface{}{}
	p.Extensions[name] = extension
	return extension
}
