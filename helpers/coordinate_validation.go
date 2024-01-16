package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateLatitude(latitude string) (float64, error) {
	// Parse latitude to a floating-point number
	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid latitude format")
	}

	// Range validation for latitude (-90 to 90 degrees)
	if lat < -90 || lat > 90 {
		return 0, fmt.Errorf("latitude out of range")
	}

	// Decimal precision check (limit to a reasonable precision, e.g., 6-8 decimal places)
	if !isValidDecimalPrecision(latitude, 8) {
		return 0, fmt.Errorf("invalid latitude decimal precision")
	}

	return lat, nil
}

func ValidateLongitude(longitude string) (float64, error) {
	// Parse longitude to a floating-point number
	lon, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid longitude format")
	}

	// Range validation for longitude (-180 to 180 degrees)
	if lon < -180 || lon > 180 {
		return 0, fmt.Errorf("longitude out of range")
	}

	// Decimal precision check (limit to a reasonable precision, e.g., 6-8 decimal places)
	if !isValidDecimalPrecision(longitude, 8) {
		return 0, fmt.Errorf("invalid longitude decimal precision")
	}

	return lon, nil
}

func isValidDecimalPrecision(coord string, maxDecimals int) bool {
	parts := splitCoordinate(coord)

	// Check if the decimal part has at most maxDecimals digits
	if len(parts) > 1 && len(parts[1]) > maxDecimals {
		return false
	}

	return true
}

func splitCoordinate(coord string) []string {
	return strings.Split(coord, ".")
}
