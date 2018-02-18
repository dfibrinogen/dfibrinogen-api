package util

import "strconv"

func ConvertStringToUint(origin string) uint {

	uid64, _ := strconv.ParseUint(origin, 10, 32)
	change := uint(uid64)

	return change
}

func GetLimitAndOffset(perPage string, page string) (int, int) {

	limit, _ := strconv.Atoi(perPage)
	offset, _ := strconv.Atoi(page)

	if offset == 1 {
		offset = 0
	} else {
		offset = (offset * limit) - limit
	}

	return limit, offset
}
