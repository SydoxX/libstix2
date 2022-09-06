package factory

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/avast/libstix2/objects"
)

const IdContribTagName = "idcontrib"

var (
	propertyNamesCache   = map[objects.ObjectType][]JsonPropertyName{}
	propertyNamesCacheMu = sync.Mutex{}
)

type JsonPropertyName struct {
	Name             string
	IsIdContributing bool
}

func GetJsonPropertyNames(typ objects.ObjectType) []JsonPropertyName {
	propertyNamesCacheMu.Lock()
	defer propertyNamesCacheMu.Unlock()

	if props, prs := propertyNamesCache[typ]; prs {
		return props
	}

	obj := Create(typ)
	if obj == nil {
		return nil
	}

	objTyp := reflect.ValueOf(obj).Elem().Type()
	if objTyp.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Got invalid type %T %s for %s", obj, objTyp, typ))
	}

	props := collectJsonFieldNames(objTyp, nil, false)
	propertyNamesCache[typ] = props
	return props
}

func isIdContributing(field reflect.StructField) bool {
	switch strings.ToLower(field.Tag.Get(IdContribTagName)) {
	case "1", "true", "yes":
		return true
	default:
		return false
	}
}

func collectJsonFieldNames(typ reflect.Type, names []JsonPropertyName, parentIdContributing bool) []JsonPropertyName {
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		idContributing := isIdContributing(f)
		if f.Anonymous {
			fieldType := f.Type
			for fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}

			if fieldType.Kind() == reflect.Struct {
				names = collectJsonFieldNames(fieldType, names, idContributing)
				continue
			}
		}

		if name := parseJsonTag(&f); name != "" {
			names = append(names, JsonPropertyName{
				Name:             name,
				IsIdContributing: idContributing || parentIdContributing,
			})
		}

	}
	return names
}

func parseJsonTag(f *reflect.StructField) string {
	jsonTag := f.Tag.Get("json")
	if jsonTag == "" {
		if f.PkgPath == "" {
			return f.Name
		}
		return ""
	}

	if jsonTag == "-" {
		return ""
	}

	tokens := strings.Split(jsonTag, ",")
	if len(tokens) < 2 || tokens[0] != "-" {
		return tokens[0]
	}
	return ""
}
