// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"encoding/json"
	"fmt"

	"github.com/nextpart/libstix2/objects/sdo/attackpattern"
	"github.com/nextpart/libstix2/objects/sdo/bundle"
	"github.com/nextpart/libstix2/objects/sdo/campaign"
	"github.com/nextpart/libstix2/objects/sdo/courseofaction"
	"github.com/nextpart/libstix2/objects/sdo/grouping"
	"github.com/nextpart/libstix2/objects/sdo/identity"
	"github.com/nextpart/libstix2/objects/sdo/indicator"
	"github.com/nextpart/libstix2/objects/sdo/infrastructure"
	"github.com/nextpart/libstix2/objects/sdo/intrusionset"
	"github.com/nextpart/libstix2/objects/sdo/location"
	"github.com/nextpart/libstix2/objects/sdo/malware"
	"github.com/nextpart/libstix2/objects/sdo/malwareanalysis"
	"github.com/nextpart/libstix2/objects/sdo/note"
	"github.com/nextpart/libstix2/objects/sdo/observeddata"
	"github.com/nextpart/libstix2/objects/sdo/opinion"
	"github.com/nextpart/libstix2/objects/sdo/report"
	"github.com/nextpart/libstix2/objects/sdo/threatactor"
	"github.com/nextpart/libstix2/objects/sdo/tool"
	"github.com/nextpart/libstix2/objects/sdo/vulnerability"
	"github.com/nextpart/libstix2/objects/sro/relationship"
	"github.com/nextpart/libstix2/objects/sro/sighting"
)

func main() {
	sm := bundle.New()

	sm.AddObject(attackpattern.New())
	sm.AddObject(campaign.New())
	sm.AddObject(courseofaction.New())
	sm.AddObject(grouping.New())
	sm.AddObject(identity.New())
	sm.AddObject(indicator.New())
	sm.AddObject(infrastructure.New())
	sm.AddObject(intrusionset.New())
	sm.AddObject(location.New())
	sm.AddObject(malware.New())
	sm.AddObject(malwareanalysis.New())
	sm.AddObject(note.New())
	sm.AddObject(observeddata.New())
	sm.AddObject(opinion.New())
	sm.AddObject(relationship.New())
	sm.AddObject(report.New())
	sm.AddObject(sighting.New())
	sm.AddObject(threatactor.New())
	sm.AddObject(tool.New())
	sm.AddObject(vulnerability.New())

	var data []byte
	data, _ = json.MarshalIndent(sm, "", "    ")

	fmt.Println(string(data))
}
