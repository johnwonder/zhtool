package util

import (
	"strings"
	"unicode"
)

// StrIsBlank Verify that the string is blank
func StrIsBlank(s string) bool {

	//https://studygolang.com/articles/17290?fr=sidebar
	//https://studygolang.com/articles/14060
	//
	//if len(s) == 0  || len(strings.Trim(s, " \n\t")) == 0 {
	//	return true
	//}
	if len(s) == 0 || len(strings.TrimSpace(s)) == 0 {
		return true
	}

	//https://studygolang.com/articles/14060
	//https://www.runoob.com/note/27977
	//https://blog.csdn.net/qq_34169240/article/details/80581365
	b := true
	for _, c := range s {
		if !unicode.IsSpace(c) {
			b = false
			break
		}
	}
	return b

}

// StrIsEmpty Verify that the string is empty
func StrIsEmpty(s string) bool {
	return len(s) == 0
}

// Reverse can reverse string
func Reverse(s string) string{

	if StrIsBlank(s) {
		return s
	}

	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
	//https://studygolang.com/articles/1147
	//var result []byte
	//tmp := []byte(s)
	//length := len(s)
	//for i := 0; i < length; i++ {
	//	result = append(result, tmp[length-i-1])
	//}
	//return string(result)
}
