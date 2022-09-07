// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package emailaddr

import (
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
EmailAddr - This type implements the STIX 2 Email Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type EmailAddr struct {
	common.CommonObjectProperties
	properties.ValueProperty
	DisplayName string `json:"display_name,omitempty" bson:"display_name,omitempty"`
	properties.BelongsToRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeEmailAddress, func() common.STIXObject {
		return New()
	})
}

func New() *EmailAddr {
	var obj EmailAddr
	obj.InitSCO(objects.TypeEmailAddress)
	return &obj
}

func (o *EmailAddr) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ValueProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
