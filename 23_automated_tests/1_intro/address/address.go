package address

import "strings"

// GetType verifies if address type is valid and returns it
func GetType(address string) string {
	validAddressTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	lowerCaseAddress := strings.ToLower(address)
	firstAddressWord := strings.Split(lowerCaseAddress, " ")[0]

	isValidAddress := false
	for _, addressType := range validAddressTypes {
		if addressType == firstAddressWord {
			isValidAddress = true
		}
	}

	if isValidAddress {
		return strings.Title(firstAddressWord)
	}

	return "Invalid type"
}
