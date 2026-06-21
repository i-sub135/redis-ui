package globutils

import "strconv"

const (
	DefaultPage  = 1
	DefaultLimit = 10
)

// ParsePagination convert page dan limit string to int dengan default values
func ParsePagination(page, limit string) (int, int) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if pageInt == 0 {
		pageInt = DefaultPage
	}
	if limitInt == 0 {
		limitInt = DefaultLimit
	}

	return pageInt, limitInt
}

// CalculateOffset hitung offset dari page dan limit
func CalculateOffset(page, limit int) int {
	return (page - 1) * limit
}
