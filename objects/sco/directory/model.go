// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package directory

import (
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Directory - This type implements the STIX 2 Directory SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Directory struct {
	common.CommonObjectProperties
	Path        string   `json:"path,omitempty" bson:"path,omitempty"`
	PathEnc     string   `json:"path_enc,omitempty" bson:"path_enc,omitempty"`
	Ctime       string   `json:"ctime,omitempty" bson:"ctime,omitempty"`
	Mtime       string   `json:"mtime,omitempty" bson:"mtime,omitempty"`
	Atime       string   `json:"atime,omitempty" bson:"atime,omitempty"`
	ContainsRef []string `json:"contains_ref,omitempty" bson:"contains_ref,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeDirectory, func() common.STIXObject {
		return New()
	})
}

func New() *Directory {
	var obj Directory
	obj.InitSCO(objects.TypeDirectory)
	return &obj
}

func (o *Directory) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}
