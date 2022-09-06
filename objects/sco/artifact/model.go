// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package artifact

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/sco/file"
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
	common.CommonObjectProperties
	MimeType            string           `json:"mime_type,omitempty" bson:"mime_type,omitempty"`
	PayloadBin          string           `json:"payload_bin,omitempty" bson:"payload_bin,omitempty"`
	Url                 string           `json:"url,omitempty" bson:"url,omitempty"`
	Hashes              *file.FileHashes `json:"hashes,omitempty" idcontrib:"1"`
	EncryptionAlgorithm string           `json:"encryption_algorithm,omitempty" bson:"encryption_algorithm,omitempty"`
	DecryptionKey       string           `json:"decryption_key,omitempty" bson:"decryption_key,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeArtifact, func() common.STIXObject {
		return New()
	})
}

func New() *Artifact {
	var obj Artifact
	obj.InitSCO(objects.TypeArtifact)
	return &obj
}

func (o *Artifact) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}
