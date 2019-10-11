package main

import "fmt"

func testInsertStringSlice() {
	s3 := InsertStringSlice("123456", "789", 3)
	fmt.Printf("%s\n", s3)
}
func InsertStringSlice(s1 string, s2 string, n int) string {
	s1 = s1[:]
	if n > len(s1) {
		s1 += s2
	} else {
		temp := s1[n:]
		s1 = s1[0:n] + s2 + temp
	}
	return s1
}
