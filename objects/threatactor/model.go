// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package threatactor

import (
	"github.com/avast/libstix2/datatypes/timestamp"
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
	"github.com/avast/libstix2/vocabs"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
ThreatActor - This type implements the STIX 2 Threat Actor SDO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type ThreatActor struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ThreatActorTypes []vocabs.ThreatActorType `json:"threat_actor_types"`
	properties.AliasesProperty
	properties.SeenProperties
	properties.RolesProperty
	properties.GoalsProperty
	Sophistication vocabs.ThreatActorSophistication `json:"sophistication,omitempty"`
	properties.ResourceLevelProperty
	properties.MotivationProperties
	PersonalMotivations []vocabs.AttackMotivation `json:"personal_motivations,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeThreatActor, func() common.STIXObject {
		return New()
	})
}

func New() *ThreatActor {
	var obj ThreatActor
	obj.InitSDO(objects.TypeThreatActor)
	return &obj
}

func (o *ThreatActor) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if len(o.ThreatActorTypes) == 0 {
		errors = append(errors, objects.PropertyMissing("threat_actor_types"))
	}

	if err := timestamp.CheckRange(o.FirstSeen, o.LastSeen, "fist_seen", "last_seen"); err != nil {
		errors = append(errors, err)
	}

	return errors
}
