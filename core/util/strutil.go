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
	if len(s) == 0  || len(strings.TrimSpace(s)) == 0 {
		return true
	}
	//return false
	b := true
	for _, c := range s {
		if !unicode.IsSpace(c){
			b =false
			break
		}
	}
	return  b

}
// StrIsEmpty Verify that the string is empty
func StrIsEmpty(s string)  bool{
	return len(s) == 0
}
