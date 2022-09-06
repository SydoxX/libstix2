// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package opinion

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
Opinion - This type implements the STIX 2 Opinion SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type Opinion struct {
	common.CommonObjectProperties
	Explanation string `json:"explanation,omitempty"`
	properties.AuthorsProperty
	Opinion string `json:"opinion"`
	properties.ObjectRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeOpinion, func() common.STIXObject {
		return New()
	})
}

func New() *Opinion {
	var obj Opinion
	obj.InitSDO(objects.TypeOpinion)
	return &obj
}

func (o *Opinion) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if o.Opinion == "" {
		errors = append(errors, objects.PropertyMissing("opinion"))
	}

	if err := o.ObjectRefsProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
