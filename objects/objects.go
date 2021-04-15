// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

/*
ValidObjectType - This function will take in a STIX object type and return
true if the string represents an actual STIX object type. This is used for
determining if input from an outside source is actually a defined STIX object or
not.
*/

type ObjectType string

const (
	TypeAttackPattern     = "attack-pattern"
	TypeBundle            = "bundle"
	TypeCampaign          = "campaign"
	TypeCourseOfAction    = "course-of-action"
	TypeGrouping          = "grouping"
	TypeIdentity          = "identity"
	TypeIndicator         = "indicator"
	TypeInfrastructure    = "infrastructure"
	TypeIntrusionSet      = "intrusion-set"
	TypeLanguageContent   = "language-content"
	TypeLocation          = "location"
	TypeMalware           = "malware"
	TypeMalwareAnalysis   = "malware-analysis"
	TypeNote              = "note"
	TypeMarkingDefinition = "marking-definition"
	TypeObservedData      = "observed-data"
	TypeOpinion           = "opinion"
	TypeReport            = "report"
	TypeRelationship      = "relationship"
	TypeThreatActor       = "threat-actor"
	TypeTool              = "tool"
	TypeSighting          = "sighting"
	TypeVulnerability     = "vulnerability"

	TypeArtifact           = "artifact"
	TypeAutonomousSystem   = "autonomous-system"
	TypeDirectory          = "directory"
	TypeDomainName         = "domain-name"
	TypeEmailAddress       = "email-addr"
	TypeEmailMessage       = "email-message"
	TypeFile               = "file"
	TypeIPv4Address        = "ipv4-addr"
	TypeIPv6Address        = "ipv6-addr"
	TypeMACAddress         = "mac-addr"
	TypeMutex              = "mutex"
	TypeNetworkTraffic     = "network-traffic"
	TypeProcess            = "process"
	TypeSoftware           = "software"
	TypeURL                = "url"
	TypeUserAccount        = "user-account"
	TypeWindowsRegistryKey = "windows-registry-key"
	TypeX509Certificate    = "x509-certificate"

	TypeExtensionDefinition = "extension-definition"
)

var ObjectTypes = map[ObjectType]bool{
	TypeAttackPattern:       true,
	TypeBundle:              true,
	TypeCampaign:            true,
	TypeCourseOfAction:      true,
	TypeGrouping:            true,
	TypeIdentity:            true,
	TypeIndicator:           true,
	TypeInfrastructure:      true,
	TypeIntrusionSet:        true,
	TypeLanguageContent:     true,
	TypeLocation:            true,
	TypeMalware:             true,
	TypeMalwareAnalysis:     true,
	TypeNote:                true,
	TypeMarkingDefinition:   true,
	TypeObservedData:        true,
	TypeOpinion:             true,
	TypeReport:              true,
	TypeRelationship:        true,
	TypeThreatActor:         true,
	TypeTool:                true,
	TypeSighting:            true,
	TypeVulnerability:       true,
	TypeArtifact:            true,
	TypeAutonomousSystem:    true,
	TypeDirectory:           true,
	TypeDomainName:          true,
	TypeEmailAddress:        true,
	TypeEmailMessage:        true,
	TypeFile:                true,
	TypeIPv4Address:         true,
	TypeIPv6Address:         true,
	TypeMACAddress:          true,
	TypeMutex:               true,
	TypeNetworkTraffic:      true,
	TypeProcess:             true,
	TypeSoftware:            true,
	TypeURL:                 true,
	TypeUserAccount:         true,
	TypeWindowsRegistryKey:  true,
	TypeX509Certificate:     true,
	TypeExtensionDefinition: true,
}

func IsValidObjectType(t string) bool {
	return ObjectTypes[ObjectType(t)]
}
