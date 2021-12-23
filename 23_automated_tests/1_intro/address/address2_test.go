// unit test
package address_test

// the package can be renamed to <package_name>_test and be on same folder as the
// original package but it needs to import the functions from the original package

import (
	// the dot here indicates that its the default import, so all the functions
	// used from this package dont need to be on format address.function(), just function()
	. "intro/address"
	"testing"
)

type testScenario struct {
	inputAddress string
	expectedType string
}

func TestGetType2(t *testing.T) {
	// if this is called it means this test can run in parallel, not necessarily a sequence
	t.Parallel()

	testScenarios := []testScenario{
		{"Avenida Paulista", "Avenida"},
		{"Rua ABC", "Rua"},
		{"Praça X", "Invalid type"},
		{"Estrada ABC", "Estrada"},
		{"rua ABC", "Rua"},
		{"RUA ABC", "Rua"},
		{"rodovia ABC", "Rodovia"},
		{"", "Invalid type"},
	}

	for _, scenario := range testScenarios {
		receivedType := GetType(scenario.inputAddress)
		if receivedType != scenario.expectedType {
			t.Errorf("Expected %s but got %s", scenario.expectedType, receivedType)
		}
	}
}

func TestAnything(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Error")
	}
}
