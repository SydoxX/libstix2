// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package infrastructure

import (
	"github.com/avast/libstix2/vocabs"

	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Infrastructure - This type implements the STIX 2 Infrastructure SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Infrastructure struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	InfrastructureTypes []vocabs.InfrastructureType `json:"infrastructure_types"`
	properties.AliasesProperty
	properties.KillChainPhasesProperty
	properties.SeenProperties
}

func init() {
	factory.RegisterObjectCreator(objects.TypeInfrastructure, func() common.STIXObject {
		return New()
	})
}

func New() *Infrastructure {
	var obj Infrastructure
	obj.InitSDO(objects.TypeInfrastructure)
	return &obj
}

func (o *Infrastructure) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if len(o.InfrastructureTypes) == 0 {
		errors = append(errors, objects.PropertyMissing("infrastructure_types"))
	}

	return errors
}
