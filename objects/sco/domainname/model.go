// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package domainname

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
DomainName - This type implements the STIX 2 Domain Name SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type DomainName struct {
	common.CommonObjectProperties
	properties.ExtensionsProperty

	properties.ValueProperty `idcontrib:"1"`

	// Deprecated
	properties.ResolvesToRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeDomainName, func() common.STIXObject {
		return New()
	})
}

func New() *DomainName {
	var obj DomainName
	obj.InitSCO(objects.TypeDomainName)
	return &obj
}

func (o *DomainName) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.ValueProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
