// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv4addr

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
IPv4Addr - This type implements the STIX 2 IPv4 Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type IPv4Addr struct {
	common.CommonObjectProperties
	properties.ValueProperty `idcontrib:"1"`
	properties.ExtensionsProperty

	// Deprecated
	properties.ResolvesToRefsProperty
	// Deprecated
	properties.BelongsToRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeIPv4Address, func() common.STIXObject {
		return New()
	})
}

func New() *IPv4Addr {
	var obj IPv4Addr
	obj.InitSCO(objects.TypeIPv4Address)
	return &obj
}

func (o *IPv4Addr) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ValueProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
