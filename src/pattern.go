package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func test_pattern_mathc() {
	searchIn := "John:2578.34 Willian:4567.23 Steve:5632.18"
	pat := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.##")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
