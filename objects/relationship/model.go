// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package relationship

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
Relationship - This type implements the STIX 2 Relationship SRO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Relationship struct {
	common.CommonObjectProperties
	RelationshipType string `json:"relationship_type"`
	properties.DescriptionProperty
	SourceRef string              `json:"source_ref"`
	TargetRef string              `json:"target_ref"`
	StartTime timestamp.Timestamp `json:"start_time,omitempty"`
	StopTime  timestamp.Timestamp `json:"stop_time,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeRelationship, func() common.STIXObject {
		return New()
	})
}

func New() *Relationship {
	var obj Relationship
	obj.InitSRO(objects.TypeRelationship)
	return &obj
}

func (o *Relationship) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if o.RelationshipType == "" {
		errors = append(errors, objects.PropertyMissing("relationship_type"))
	}

	if o.SourceRef == "" {
		errors = append(errors, objects.PropertyMissing("source_ref"))
	}

	if o.TargetRef == "" {
		errors = append(errors, objects.PropertyMissing("target_ref"))
	}

	if err := timestamp.CheckRange(o.StartTime, o.StopTime, "start_time", "stop_time"); err != nil {
		errors = append(errors, err)
	}

	return errors
}
