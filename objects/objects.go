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
type ObjectTypeId uint16

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

// Do NOT change order of these, they create the ObjectTypeId value
var objectTypeArray = [...]ObjectType{
	TypeAttackPattern,
	TypeBundle,
	TypeCampaign,
	TypeCourseOfAction,
	TypeGrouping,
	TypeIdentity,
	TypeIndicator,
	TypeInfrastructure,
	TypeIntrusionSet,
	TypeLanguageContent,
	TypeLocation,
	TypeMalware,
	TypeMalwareAnalysis,
	TypeNote,
	TypeMarkingDefinition,
	TypeObservedData,
	TypeOpinion,
	TypeReport,
	TypeRelationship,
	TypeThreatActor,
	TypeTool,
	TypeSighting,
	TypeVulnerability,
	TypeArtifact,
	TypeAutonomousSystem,
	TypeDirectory,
	TypeDomainName,
	TypeEmailAddress,
	TypeEmailMessage,
	TypeFile,
	TypeIPv4Address,
	TypeIPv6Address,
	TypeMACAddress,
	TypeMutex,
	TypeNetworkTraffic,
	TypeProcess,
	TypeSoftware,
	TypeURL,
	TypeUserAccount,
	TypeWindowsRegistryKey,
	TypeX509Certificate,
	TypeExtensionDefinition,
}

const ObjectTypeIdInvalid ObjectTypeId = 0

var ObjectTypes map[ObjectType]ObjectTypeId

func init() {
	ObjectTypes = make(map[ObjectType]ObjectTypeId, len(objectTypeArray))
	for i, typ := range objectTypeArray {
		ObjectTypes[typ] = ObjectTypeId(i + 1)
	}
}

func IsValidObjectType(t string) bool {
	return ObjectTypes[ObjectType(t)] != 0
}

func ObjectTypeToId(typ ObjectType) (id ObjectTypeId, ok bool) {
	id, ok = ObjectTypes[typ]
	return
}

func ObjectTypeIdToType(id ObjectTypeId) (typ ObjectType, ok bool) {
	idx := int(id) - 1
	if idx < 0 || idx >= len(objectTypeArray) {
		return
	}
	ok = true
	typ = objectTypeArray[idx]
	return
}

func (id ObjectTypeId) Type() ObjectType {
	typ, _ := ObjectTypeIdToType(id)
	return typ
}

func (typ ObjectType) TypeId() ObjectTypeId {
	id, _ := ObjectTypeToId(typ)
	return id
}
