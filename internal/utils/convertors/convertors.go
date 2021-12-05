package convertors

import "strconv"

func GetStringListAsInt(input []string) []int {
	values := []int{}
	for _, entry := range input {
		value, _ := strconv.Atoi(entry)
		values = append(values, value)
	}
	return values
}