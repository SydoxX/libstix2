package factory_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/nextpart/libstix2/objects"
	"github.com/nextpart/libstix2/objects/factory"

	_ "github.com/nextpart/libstix2/objects/factory/importall"
)

func TestGetJsonPropertyNames(t *testing.T) {
	expectedNames := map[objects.ObjectType][]factory.JsonPropertyName{
		objects.TypeAttackPattern: {
			{Name: "type"},
			{Name: "spec_version"},
			{Name: "id"},
			{Name: "created_by_ref"},
			{Name: "created"},
			{Name: "modified"},
			{Name: "revoked"},
			{Name: "labels"},
			{Name: "confidence"},
			{Name: "lang"},
			{Name: "external_references"},
			{Name: "object_marking_refs"},
			{Name: "granular_markings"},
			{Name: "custom"},
			{Name: "name"},
			{Name: "description"},
			{Name: "aliases"},
			{Name: "kill_chain_phases"},
			{Name: "extensions", IsIdContributing: true},
		},
		objects.TypeCampaign: {
			{Name: "type"},
			{Name: "spec_version"},
			{Name: "id"},
			{Name: "created_by_ref"},
			{Name: "created"},
			{Name: "modified"},
			{Name: "revoked"},
			{Name: "labels"},
			{Name: "confidence"},
			{Name: "lang"},
			{Name: "external_references"},
			{Name: "object_marking_refs"},
			{Name: "granular_markings"},
			{Name: "custom"},
			{Name: "name"},
			{Name: "description"},
			{Name: "aliases"},
			{Name: "first_seen"},
			{Name: "last_seen"},
			{Name: "objective"},
			{Name: "extensions", IsIdContributing: true},
		},
		objects.TypeDomainName: {
			{Name: "type"},
			{Name: "spec_version"},
			{Name: "id"},
			{Name: "created_by_ref"},
			{Name: "created"},
			{Name: "modified"},
			{Name: "revoked"},
			{Name: "labels"},
			{Name: "confidence"},
			{Name: "lang"},
			{Name: "external_references"},
			{Name: "object_marking_refs"},
			{Name: "granular_markings"},
			{Name: "custom"},
			{Name: "value", IsIdContributing: true},
			{Name: "resolves_to_refs"},
			{Name: "extensions", IsIdContributing: true},
		},
	}

	for typ := range objects.ObjectTypes {
		names := factory.GetJsonPropertyNames(typ)

		t.Log(typ, names)

		expected := expectedNames[typ]
		if expected == nil {
			continue
		}

		valid := len(expected) == len(names)
		if valid {
			sort.Slice(expected, func(i, j int) bool {
				return strings.Compare(expected[i].Name, expected[j].Name) < 0
			})
			sort.Slice(names, func(i, j int) bool {
				return strings.Compare(names[i].Name, names[j].Name) < 0
			})

			for i := range expected {
				if expected[i] != names[i] {
					valid = false
					break
				}
			}
		}

		if !valid {
			t.Errorf("Wrong json property names for %s:\nExpected: %v\nGot:      %v",
				typ, expected, names)
		}
	}

}
