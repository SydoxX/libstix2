// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package autonomoussystem

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
AutonomousSystem - This type implements the STIX 2 Autonomous System SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type AutonomousSystem struct {
	common.CommonObjectProperties
	Number int `json:"number,omitempty" bson:"number,omitempty"`
	properties.NameProperty
	Rir string `json:"rir,omitempty" bson:"rir,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeAutonomousSystem, func() common.STIXObject {
		return New()
	})
}

func New() *AutonomousSystem {
	var obj AutonomousSystem
	obj.InitSCO(objects.TypeAutonomousSystem)
	return &obj
}

func (o *AutonomousSystem) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}
