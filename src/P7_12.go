package main

func test_string_split1() {
	s := "123456789"
	s1, s2 := string_split1(s, 8)
	print(s1, "\n")
	print(s2, "\n")

	s3 := string_split2("StringSplit21")
	print(s3 + "\n")
}

func string_split1(s string, index int) (s1 string, s2 string) {
	s1 = s[:index]
	s2 = s[index:]
	return s1, s2
}

func string_split2(s string) (s1 string) {
	s1 = s[len(s)/2:] + s[:len(s)/2]
	return s1
}
