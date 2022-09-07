// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package observeddata

import (
	"github.com/nextpart/libstix2/datatypes/timestamp"

	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
ObservedData - This type implements the STIX 2 Observed Data SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type ObservedData struct {
	common.CommonObjectProperties
	FirstObserved  timestamp.Timestamp `json:"first_observed"`
	LastObserved   timestamp.Timestamp `json:"last_observed"`
	NumberObserved int                 `json:"number_observed"`
	properties.ObjectRefsProperty

	// Deprecated: Objects is deprecated in Stix2.1
	Objects map[string]interface{} `json:"objects,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeObservedData, func() common.STIXObject {
		return New()
	})
}

func New() *ObservedData {
	var obj ObservedData
	obj.InitSDO(objects.TypeObservedData)
	return &obj
}

func (o *ObservedData) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if o.NumberObserved <= 0 {
		errors = append(errors, objects.PropertyInvalid("number_observed", o.NumberObserved, "must be > 0"))
	}

	if len(o.ObjectRefs) != 0 && len(o.Objects) != 0 {
		errors = append(errors, objects.PropertyInvalid("object", o, "only one of 'objects' and 'object_refs' can be present"))
	} else if len(o.ObjectRefs) == 0 && len(o.Objects) == 0 {
		errors = append(errors, objects.PropertyInvalid("object_refs", o, "one of 'objects' and 'object_refs' must be present"))
	}

	if o.FirstObserved.IsZero() {
		errors = append(errors, objects.PropertyMissing("first_observed"))
	}

	if o.LastObserved.IsZero() {
		errors = append(errors, objects.PropertyMissing("last_observed"))
	}

	if err := timestamp.CheckRange(&o.FirstObserved, &o.LastObserved, "first_observed", "last_observed"); err != nil {
		errors = append(errors, err)
	}

	return errors
}
