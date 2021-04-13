package file

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/properties"
)

type FileHashes struct {
	Sha512 string `json:"SHA-512,omitempty"`
	Sha256 string `json:"SHA-256,omitempty"`
	Sha1   string `json:"SHA-1,omitempty"`
	MD5    string `json:"MD5,omitempty"`
}

type File struct {
	objects.CommonObjectProperties
	properties.ExtensionsProperty
	properties.GranularMarking

	Hashes             *FileHashes `json:"hashes,omitempty"`
	Size               int64       `json:"size,omitempty"`
	Name               string      `json:"name,omitempty"`
	NameEnc            string      `json:"name_enc,omitempty"`
	MagicNumberHex     string      `json:"magic_number_hex,omitempty"`
	MimeType           string      `json:"mime_type,omitempty"`
	CTime              string      `json:"ctime,omitempty"`
	MTime              string      `json:"mtime,omitempty"`
	ATime              string      `json:"atime,omitempty"`
	ParentDirectoryRef string      `json:"parent_directory_ref,omitempty"`
	ContainsRef        []string    `json:"contains_ref,omitempty"`
	ContentRef         string      `json:"content_ref,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *File) GetPropertyList() []string {
	return []string{
		"lang", "marking_ref", "selectors", "extensions", "hashes", "size", "name", "name_enc", "magic_number_hex", "mime_type",
		"ctime", "mtime", "atime", "parent_directory_ref", "contains_ref", "content_ref",
	}
}

func (o *File) IdContributingProperties() []string {
	return []string{
		"hashes", "name", "extensions", "parent_directory_ref",
	}
}

func (o *File) FixupIdContributingProperties(properties map[string]interface{}) {
	hashes, ok := properties["hashes"].(map[string]interface{})
	if !ok {
		return
	}

	for _, algo := range [...]string{"MD5", "SHA-1", "SHA-256", "SHA-512"} {
		if val, prs := hashes[algo]; prs {
			properties["hashes"] = map[string]interface{}{
				algo: val,
			}
			return
		}
	}

	delete(properties, "hashes")
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Domain Name SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *File {
	var obj File
	obj.InitSCO("file")
	return &obj
}
