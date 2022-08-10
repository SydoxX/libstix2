// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package artifact

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Artifact - This type implements the STIX 2 Artifact SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Artifact struct {
	objects.CommonObjectProperties
	MimeType   string `json:"mime_type,omitempty" bson:"mime_type,omitempty"`
	PayloadBin string `json:"payload_bin,omitempty" bson:"payload_bin,omitempty"`
	Url        string `json:"url,omitempty" bson:"url,omitempty"`
	objects.HashesProperty
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty" bson:"encryption_algorithm,omitempty"`
	DecryptionKey       string `json:"decryption_key,omitempty" bson:"decryption_key,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *Artifact) GetPropertyList() []string {
	return []string{
		"mime_type",
		"payload_bin",
		"url",
		"hashes",
		"encryption_algorithm",
		"decryption_key",
	}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX File SCO and return it as a
pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Artifact {
	var obj Artifact
	obj.InitSCO("artifact")
	return &obj
}
