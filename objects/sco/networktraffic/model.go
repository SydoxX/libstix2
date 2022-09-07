// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package networktraffic

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
NetworkTraffic - This type implements the STIX 2 Network Traffic SCO and defines
all of the properties and methods needed to create and work with this object.
All of the methods not defined local to this type are inherited from the
individual properties.
*/
type NetworkTraffic struct {
	common.CommonObjectProperties
	properties.ExtensionsProperty
	Start             string            `json:"start,omitempty" bson:"start,omitempty"`
	End               string            `json:"end,omitempty" bson:"end,omitempty"`
	IsActive          bool              `json:"is_active,omitempty" bson:"is_active,omitempty"`
	SrcRef            string            `json:"src_ref,omitempty" bson:"src_ref,omitempty"`
	DstRef            string            `json:"dst_ref,omitempty" bson:"dst_ref,omitempty"`
	SrcPort           int               `json:"src_port,omitempty" bson:"src_port,omitempty"`
	DstPort           int               `json:"dst_port,omitempty" bson:"dst_port,omitempty"`
	Protocols         []string          `json:"protocols,omitempty" bson:"protocols,omitempty"`
	SrcByteCount      int               `json:"src_byte_count,omitempty" bson:"src_byte_count,omitempty"`
	DstByteCount      int               `json:"dst_byte_count,omitempty" bson:"dst_byte_count,omitempty"`
	SrcPackets        int               `json:"src_packets,omitempty" bson:"src_packets,omitempty"`
	DstPackets        int               `json:"dst_packets,omitempty" bson:"dst_packets,omitempty"`
	Ipfix             map[string]string `json:"ipfix,omitempty" bson:"ipfix,omitempty"`
	SrcPayloadRef     string            `json:"src_payload_ref,omitempty" bson:"src_payload_ref,omitempty"`
	DstPayloadRef     string            `json:"dst_payload_ref,omitempty" bson:"dst_payload_ref,omitempty"`
	EncapsulatesRefs  []string          `json:"encapsulates_refs,omitempty" bson:"encapsulates_refs,omitempty"`
	EncapsulatedByRef string            `json:"encapsulated_by_ref,omitempty" bson:"encapsulated_by_ref,omitempty"`
}

func init() {
	factory.RegisterObjectCreator(objects.TypeNetworkTraffic, func() common.STIXObject {
		return New()
	})
}

func New() *NetworkTraffic {
	var obj NetworkTraffic
	obj.InitSCO(objects.TypeNetworkTraffic)
	return &obj
}

func (o *NetworkTraffic) Valid() []error {
	errors := o.CommonObjectProperties.ValidSDO()

	return errors
}
