// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package sighting

import (
	"github.com/avast/libstix2/datatypes/timestamp"
	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Sighting - This type implements the STIX 2 Sighting SRO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Sighting struct {
	common.CommonObjectProperties
	properties.DescriptionProperty
	properties.SeenProperties
	Count            int      `json:"count,omitempty"`
	SightingOfRef    string   `json:"sighting_of_ref"`
	ObservedDataRefs []string `json:"observed_data_refs,omitempty"`
	WhereSightedRefs []string `json:"where_sighted_refs,omitempty"`
	Summary          bool     `json:"summary,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeSighting, func() common.STIXObject {
		return New()
	})
}

func New() *Sighting {
	var obj Sighting
	obj.InitSRO(objects.TypeSighting)
	return &obj
}

func (o *Sighting) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if len(o.SightingOfRef) == 0 {
		errors = append(errors, objects.PropertyMissing("sighting_of_ref"))
	}

	if o.Count <= 0 {
		errors = append(errors, objects.PropertyInvalid("count", o.Count, "Must be >= 0"))
	}

	if err := timestamp.CheckRange(o.FirstSeen, o.LastSeen, "fist_seen", "last_seen"); err != nil {
		errors = append(errors, err)
	}

	return errors
}
