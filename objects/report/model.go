// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package report

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
Report - This type implements the STIX 2 Report SDO and defines all of the
properties and methods needed to create and work with this object. All of the
methods not defined local to this type are inherited from the individual
properties.
*/
type Report struct {
	common.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	ReportTypes []vocabs.ReportType `json:"report_types,omitempty"`
	Published   timestamp.Timestamp `json:"published,omitempty"`
	properties.ObjectRefsProperty
}

func init() {
	factory.RegisterObjectCreator(objects.TypeReport, func() common.STIXObject {
		return New()
	})
}

func New() *Report {
	var obj Report
	obj.InitSDO(objects.TypeReport)
	return &obj
}

func (o *Report) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	if err := o.NameProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	if len(o.ReportTypes) == 0 {
		errors = append(errors, objects.PropertyMissing("report_types"))
	}

	if o.Published.IsZero() {
		errors = append(errors, objects.PropertyMissing("published"))
	}

	if err := o.ObjectRefsProperty.VerifyExists(); err != nil {
		errors = append(errors, err)
	}

	return errors
}
