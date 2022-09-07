// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package urlobject

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
URLObject - This type implements the STIX 2 URL SCO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type URLObject struct {
	common.CommonObjectProperties
	properties.ValueProperty `idcontrib:"1"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeURL, func() common.STIXObject {
		return New()
	})
}

func New() *URLObject {
	var obj URLObject
	obj.InitSCO(objects.TypeURL)
	return &obj
}

func (o *URLObject) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ValueProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
