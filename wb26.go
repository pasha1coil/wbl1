package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Println(Helper(s1))
	fmt.Println(Helper(s2))
	fmt.Println(Helper(s3))
}

func Helper(s string) bool {
	s = strings.ToLower(s)
	m := make(map[int32]bool)

	for _, i := range s {
		if ok := m[i]; ok {
			return false
		}
		m[i] = true
	}
	return true
}
