// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"github.com/avast/libstix2/datatypes/timestamp"

	"github.com/avast/libstix2/vocabs"

	"github.com/avast/libstix2/objects"
	"github.com/avast/libstix2/objects/common"
	"github.com/avast/libstix2/objects/factory"
	"github.com/avast/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/*
Indicator - This type implements the STIX 2 Indicator SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties.
*/
type Indicator struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty

	IndicatorTypes []vocabs.IndicatorType `json:"indicator_types"`
	Pattern        string                 `json:"pattern"`
	PatternType    vocabs.PatternType     `json:"pattern_type"`
	PatternVersion string                 `json:"pattern_version,omitempty"`
	ValidFrom      timestamp.Timestamp    `json:"valid_from"`
	ValidUntil     timestamp.Timestamp    `json:"valid_until,omitempty"`
	properties.KillChainPhasesProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeIndicator, func() common.STIXObject {
		return New()
	})
}

func New() *Indicator {
	var obj Indicator
	obj.InitSDO(objects.TypeIndicator)
	return &obj
}

func (o *Indicator) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if len(o.IndicatorTypes) == 0 {
		errors = append(errors, objects.PropertyMissing("indicator_types"))
	}

	if len(o.Pattern) == 0 {
		errors = append(errors, objects.PropertyMissing("pattern"))
	}

	if len(o.PatternType) == 0 {
		errors = append(errors, objects.PropertyMissing("pattern_type"))
	}

	if o.ValidFrom.IsZero() {
		errors = append(errors, objects.PropertyMissing("valid_from"))
	}

	if err := timestamp.CheckRange(o.ValidFrom, o.ValidUntil, "valid_from", "valid_to"); err != nil {
		errors = append(errors, err)
	}

	return errors
}
