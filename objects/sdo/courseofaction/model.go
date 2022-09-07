// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package courseofaction

import (
	"fmt"

	"github.com/nextpart/libstix2/vocabs"

	"github.com/nextpart/libstix2/datatypes/binary"
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
CourseOfAction - This type implements the STIX 2 Course Of Action SDO and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type CourseOfAction struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty

	ActionType      vocabs.CourseOfActionType `json:"action_type,omitempty"`
	OsExecutionEnvs []string                  `json:"os_execution_envs,omitempty"`
	ActionBin       binary.Binary             `json:"action_bin,omitempty"`
	ActionReference string                    `json:"action_reference,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeCourseOfAction, func() common.STIXObject {
		return New()
	})
}

/*
New - This function will create a new STIX Course of Action object and return
it as a pointer.
*/
func New() *CourseOfAction {
	var obj CourseOfAction
	obj.InitSDO(objects.TypeCourseOfAction)
	return &obj
}

func (o *CourseOfAction) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if len(o.ActionBin) != 0 && o.ActionReference != "" {
		errors = append(errors, fmt.Errorf("Only one of [action_bin, action_reference] can be specified."))
	}

	return errors
}
