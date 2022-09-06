// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package note

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Note - This type implements the STIX 2 Note SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Note struct {
	common.CommonObjectProperties
	Abstract string `json:"abstract,omitempty"`
	Content  string `json:"content"`
	properties.AuthorsProperty
	properties.ObjectRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeNote, func() common.STIXObject {
		return New()
	})
}

func New() *Note {
	var obj Note
	obj.InitSDO(objects.TypeNote)
	return &obj
}

func (o *Note) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if len(o.Content) == 0 {
		errors = append(errors, objects.PropertyMissing("content"))
	}

	if err := o.ObjectRefsProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
