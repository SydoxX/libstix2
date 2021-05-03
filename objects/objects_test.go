package objects

import (
	"testing"
)

func TestObjectTypeIds(t *testing.T) {
	expected := map[ObjectType]int{
		TypeAttackPattern:       1,
		TypeBundle:              2,
		TypeCampaign:            3,
		TypeCourseOfAction:      4,
		TypeGrouping:            5,
		TypeIdentity:            6,
		TypeIndicator:           7,
		TypeInfrastructure:      8,
		TypeIntrusionSet:        9,
		TypeLanguageContent:     10,
		TypeLocation:            11,
		TypeMalware:             12,
		TypeMalwareAnalysis:     13,
		TypeNote:                14,
		TypeMarkingDefinition:   15,
		TypeObservedData:        16,
		TypeOpinion:             17,
		TypeReport:              18,
		TypeRelationship:        19,
		TypeThreatActor:         20,
		TypeTool:                21,
		TypeSighting:            22,
		TypeVulnerability:       23,
		TypeArtifact:            24,
		TypeAutonomousSystem:    25,
		TypeDirectory:           26,
		TypeDomainName:          27,
		TypeEmailAddress:        28,
		TypeEmailMessage:        29,
		TypeFile:                30,
		TypeIPv4Address:         31,
		TypeIPv6Address:         32,
		TypeMACAddress:          33,
		TypeMutex:               34,
		TypeNetworkTraffic:      35,
		TypeProcess:             36,
		TypeSoftware:            37,
		TypeURL:                 38,
		TypeUserAccount:         39,
		TypeWindowsRegistryKey:  40,
		TypeX509Certificate:     41,
		TypeExtensionDefinition: 42,
	}

	for typ, expectedId := range expected {
		id, ok := ObjectTypeToId(typ)
		if !ok {
			t.Errorf("Got invalid id response for %s", typ)
		} else if id != ObjectTypeId(expectedId) {
			t.Errorf("expected id does not match %s %d != %d", typ, id, expectedId)
		}
	}

	if _, ok := ObjectTypeIdToType(ObjectTypeIdInvalid); ok {
		t.Errorf("Type id 0 is okay but is supposed to be bad.")
	}
}
