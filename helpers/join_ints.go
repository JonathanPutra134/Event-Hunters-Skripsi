package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func JoinInts(ints []int) string {
	fmt.Println("CATEGORY IDS in Join Ints")
	fmt.Println(ints)
	var strInts []string
	for _, i := range ints {
		strInts = append(strInts, strconv.Itoa(i))
	}
	return strings.Join(strInts, ", ")
}
