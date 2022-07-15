package global

import "strconv"

func StringToUInt(value string) (uint, error) {
	var valueUInt uint

	valueInt, TypeErr := strconv.Atoi(value)

	valueUInt = uint(valueInt)

	return valueUInt, TypeErr
}
