package file

import (
	"strings"

	"github.com/nextpart/libstix2/datatypes/hex"
	"github.com/nextpart/libstix2/datatypes/timestamp"
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
)

type FileHashes struct {
	Sha512 string `json:"SHA-512,omitempty"`
	Sha256 string `json:"SHA-256,omitempty"`
	Sha1   string `json:"SHA-1,omitempty"`
	MD5    string `json:"MD5,omitempty"`
}

type File struct {
	common.CommonObjectProperties
	properties.GranularMarking
	properties.NameProperty `idcontrib:"1"`

	Hashes             *FileHashes          `json:"hashes,omitempty" idcontrib:"1"`
	Size               int64                `json:"size,omitempty"`
	NameEnc            string               `json:"name_enc,omitempty"`
	MagicNumberHex     hex.Hex              `json:"magic_number_hex,omitempty"`
	MimeType           string               `json:"mime_type,omitempty"`
	CTime              *timestamp.Timestamp `json:"ctime,omitempty"`
	MTime              *timestamp.Timestamp `json:"mtime,omitempty"`
	ATime              *timestamp.Timestamp `json:"atime,omitempty"`
	ParentDirectoryRef string               `json:"parent_directory_ref,omitempty"`
	ContainsRef        []string             `json:"contains_ref,omitempty"`
	ContentRef         string               `json:"content_ref,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeFile, func() common.STIXObject {
		return New()
	})
}

func (o *File) FixupIdContributingProperties(properties map[string]any) {
	hashes, ok := properties["hashes"].(map[string]any)
	if !ok {
		return
	}

	for _, algo := range [...]string{"MD5", "SHA-1", "SHA-256", "SHA-512"} {
		if val, prs := hashes[algo]; prs {
			properties["hashes"] = map[string]any{
				algo: val,
			}
			return
		}
	}

	delete(properties, "hashes")
}

func New() *File {
	var obj File
	obj.InitSCO(objects.TypeFile)
	return &obj
}

func (o *File) Valid() []error {
	var errors []error

	if err := o.TypeProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if err := o.IDProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if o.Size < 0 {
		errors = append(errors, objects.PropertyInvalid("size", o.Size, "Must be >= 0"))
	}

	if o.ParentDirectoryRef != "" && !strings.HasPrefix(o.ParentDirectoryRef, objects.TypeDirectory) {
		errors = append(errors, objects.PropertyInvalid("parent_directory_ref", o.ParentDirectoryRef, "must be reference to a directory object."))
	}

	if o.ContentRef != "" && !strings.HasPrefix(o.ContentRef, objects.TypeArtifact) {
		errors = append(errors, objects.PropertyInvalid("content_ref", o.ContentRef, "must be reference to an artifact object."))
	}

	return errors
}
