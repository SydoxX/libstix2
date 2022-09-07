package objects

import "fmt"

type PropertyError struct {
	PropertyName string
	Err          error
}

func (p *PropertyError) Error() string {
	return p.Err.Error()
}

type PropertyMissingError struct {
	PropertyError
}

type PropertyInvalidValueError struct {
	PropertyError
	Value any
}

func PropertyMissing(propertyName string) error {
	return &PropertyMissingError{
		PropertyError{
			PropertyName: propertyName,
			Err:          fmt.Errorf("required property '%s' is missing", propertyName),
		},
	}
}

func PropertyInvalid(propertyName string, value any, extraInfo string) error {
	return &PropertyInvalidValueError{
		PropertyError: PropertyError{
			PropertyName: propertyName,
			Err: fmt.Errorf("property '%s' has invalid value: %#v (%s)",
				propertyName, value, extraInfo),
		},
		Value: value,
	}
}
