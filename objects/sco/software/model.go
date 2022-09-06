// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package software

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
Software - This type implements the STIX 2 Software SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Software struct {
	common.CommonObjectProperties
	properties.NameProperty
	Cpe       string   `json:"cpe,omitempty" bson:"cpe,omitempty"`
	Languages []string `json:"languages,omitempty" bson:"languages,omitempty"`
	Vendor    string   `json:"vendor,omitempty" bson:"vendor,omitempty"`
	Version   string   `json:"version,omitempty" bson:"version,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeSoftware, func() common.STIXObject {
		return New()
	})
}

func New() *Software {
	var obj Software
	obj.InitSCO(objects.TypeSoftware)
	return &obj
}

func (o *Software) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}
