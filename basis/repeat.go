package main

const repeatCount = 5

func Repeat(word string) string {
	var repeat string
	for i := 0; i < repeatCount; i++ {
		repeat += word
	}
	return repeat
}
