package vocabs

type ExtensionType string

const (
	ExtensionNewSdo           ExtensionType = "new-sdo"
	ExtensionNewSco           ExtensionType = "new-sco"
	ExtensionNewSro           ExtensionType = "new-sro"
	ExtensionProperty         ExtensionType = "property-extension"
	ExtensionPropertyTopLevel ExtensionType = "toplevel-property-extension"
)

func (ex ExtensionType) IsValid() bool {
	switch ex {
	case ExtensionNewSdo, ExtensionNewSco, ExtensionNewSro, ExtensionProperty, ExtensionPropertyTopLevel:
		return true
	default:
		return false
	}
}
