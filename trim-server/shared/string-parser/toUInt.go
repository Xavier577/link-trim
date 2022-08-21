package stringparser

import "strconv"

func ToUInt(value string) (uint, error) {
	var valueUInt uint

	valueInt, TypeErr := strconv.Atoi(value)

	valueUInt = uint(valueInt)

	return valueUInt, TypeErr
}
