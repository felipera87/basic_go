// super basic unit test
package address

import "testing"

func TestGetType(t *testing.T) {
	addressToTest := "Avenida Paulista"
	expectedAddressType := "Avenida"
	receivedAddressType := GetType(addressToTest)

	if expectedAddressType != receivedAddressType {
		t.Errorf("Expected %s but got %s", expectedAddressType, receivedAddressType)
	}
}
