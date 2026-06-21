package enums

import (
	"fmt"
)

type ExampleEnums string

const (
	ExampleIn  ExampleEnums = "Example_in"
	ExampleOut ExampleEnums = "Example_out"
)

var ExampleTypes = []string{
	string(ExampleIn),
	string(ExampleOut),
}

func (s ExampleEnums) IsValid() bool {
	switch s {
	case ExampleIn, ExampleOut:
		return true
	default:
		return false
	}
}

func ParseExampleType(v string) (ExampleEnums, error) {
	s := ExampleEnums(v)

	if !s.IsValid() {
		return "", fmt.Errorf("invalid sync type: %s", v)
	}
	return s, nil
}
