// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package tool

import (
	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/common"
	"github.com/nextpart/libstix2/objects/factory"
	"github.com/nextpart/libstix2/objects/properties"
	"github.com/nextpart/libstix2/vocabs"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
Tool - This type implements the STIX 2 Tool SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type Tool struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ToolTypes []vocabs.ToolType `json:"tool_types,omitempty"`
	properties.AliasesProperty
	properties.KillChainPhasesProperty
	ToolVersion string `json:"tool_version,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeTool, func() common.STIXObject {
		return New()
	})
}

func New() *Tool {
	var obj Tool
	obj.InitSDO(objects.TypeTool)
	return &obj
}

func (o *Tool) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if len(o.ToolTypes) == 0 {
		errors = append(errors, objects.PropertyMissing("threat_actor_types"))
	}

	return errors
}
