// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package intrusionset

import (
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
IntrusionSet - This type implements the STIX 2 Intrusion Set SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type IntrusionSet struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.SeenProperties
	properties.GoalsProperty
	properties.ResourceLevelProperty
	properties.MotivationProperties
}

func init() {
	factory.RegisterObjectCreator(objects.TypeIntrusionSet, func() common.STIXObject {
		return New()
	})
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Intrusion Set object and return
it as a pointer. It will also initialize the object by setting all of the basic
properties.
*/
func New() *IntrusionSet {
	var obj IntrusionSet
	obj.InitSDO(objects.TypeIntrusionSet)
	return &obj
}

func (o *IntrusionSet) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}
	return errors
}
