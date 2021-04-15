// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package attackpattern

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
AttackPattern - This type implements the STIX 2 Attack Pattern SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type AttackPattern struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	properties.AliasesProperty
	properties.KillChainPhasesProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeAttackPattern, func() common.STIXObject {
		return New()
	})
}

func New() *AttackPattern {
	var obj AttackPattern
	obj.InitSDO(objects.TypeAttackPattern)
	return &obj
}

func (o *AttackPattern) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
