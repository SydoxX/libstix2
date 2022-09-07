// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package process

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
Process - This type implements the STIX 2 Process SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type Process struct {
	common.CommonObjectProperties
	properties.ExtensionsProperty
	IsHidden             bool              `json:"is_hidden,omitempty" bson:"is_hidden,omitempty"`
	Pid                  int               `json:"pid,omitempty" bson:"pid,omitempty"`
	CreatedTime          string            `json:"created_time,omitempty" bson:"created_time,omitempty"`
	Cwd                  string            `json:"cwd,omitempty" bson:"cwd,omitempty"`
	CommandLine          string            `json:"command_line,omitempty" bson:"command_line,omitempty"`
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty" bson:"environment_variables,omitempty"`
	OpenedConnectionRefs []string          `json:"opened_connection_refs,omitempty" bson:"opened_connection_refs,omitempty"`
	CreatorUserRef       string            `json:"creator_user_ref,omitempty" bson:"creator_user_ref,omitempty"`
	ImageRef             string            `json:"image_ref,omitempty" bson:"image_ref,omitempty"`
	ParentRef            string            `json:"parent_ref,omitempty" bson:"parent_ref,omitempty"`
	ChildRefs            []string          `json:"child_refs,omitempty" bson:"child_refs,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeProcess, func() common.STIXObject {
		return New()
	})
}

func New() *Process {
	var obj Process
	obj.InitSCO(objects.TypeProcess)
	return &obj
}

func (o *Process) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}
