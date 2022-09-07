// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package ipv6addr

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
IPv6Addr - This type implements the STIX 2 IPv6 Address SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type IPv6Addr struct {
	common.CommonObjectProperties
	properties.ValueProperty
	properties.ResolvesToRefsProperty
	properties.BelongsToRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeIPv6Address, func() common.STIXObject {
		return New()
	})
}

func New() *IPv6Addr {
	var obj IPv6Addr
	obj.InitSCO(objects.TypeIPv6Address)
	return &obj
}

func (o *IPv6Addr) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ValueProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
