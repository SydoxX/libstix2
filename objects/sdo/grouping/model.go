// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package grouping

import (
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
	"github.com/nextpart/libstix2/vocabs"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Grouping - This type implements the STIX 2 Grouping SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Grouping struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.ObjectRefsProperty

	Context vocabs.GroupingContext `json:"context"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeGrouping, func() common.STIXObject {
		return New()
	})
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Grouping object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *Grouping {
	var obj Grouping
	obj.InitSDO(objects.TypeGrouping)
	return &obj
}

func (o *Grouping) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ObjectRefsProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if o.Context == "" {
		errors = append(errors, objects.PropertyMissing("context"))
	}

	return errors
}
