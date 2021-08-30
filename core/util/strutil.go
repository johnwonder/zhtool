package util

import (
	"fmt"
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


	//var result string
	//strLen := len(s)
	//for i := 0; i < strLen; i++ {
	//	result = result + fmt.Sprintf("%c", s[strLen-i-1])
	//}
	//return result
}

func SplitParams(in string) (name string, par2 string) {
	i := strings.IndexFunc(strings.TrimSpace(in), unicode.IsSpace)
	if i < 1 {
		return strings.TrimSpace(in), ""
	}

	return strings.TrimSpace(in[:i+1]), strings.TrimSpace(in[i+1:])
}

func Tokenize(in string) interface{} {
	first := strings.Fields(in)
	var final = make([]string, 0)
	var keys = make([]string, 0)
	inQuote := false
	start := 0

	for i, v := range first {
		index := strings.Index(v, "=")

		if !inQuote {
			if index > 1 {
				keys = append(keys, v[:index])
				v = v[index+1:]
			}
		}

		if !strings.HasPrefix(v, "&ldquo;") && !inQuote {
			final = append(final, v)
		} else if inQuote && strings.HasSuffix(v, "&rdquo;") && !strings.HasSuffix(v, "\\\"") {
			first[i] = v[:len(v)-7]
			final = append(final, strings.Join(first[start:i+1], " "))
			inQuote = false
		} else if strings.HasPrefix(v, "&ldquo;") && !inQuote {
			if strings.HasSuffix(v, "&rdquo;") {
				final = append(final, v[7:len(v)-7])
			} else {
				start = i
				first[i] = v[7:]
				inQuote = true
			}
		}

		// No closing "... just make remainder the final token
		if inQuote && i == len(first) {
			final = append(final, first[start:len(first)]...)
		}
	}

	if len(keys) > 0 {
		var m = make(map[string]string)
		for i, k := range keys {
			m[k] = final[i]
		}

		return m
	}

	return final
}

func ReplaceTemplate(stringToParse string,tempStrStart string,tempStrEnd string) string {
	posStart := strings.Index(stringToParse, tempStrStart)
	if posStart > 0 {
		posEnd := strings.Index(stringToParse[posStart:], tempStrEnd) + posStart
		if posEnd > posStart {
			name, par := SplitParams(stringToParse[posStart+3 : posEnd])

			fmt.Printf("name:%v,par:%v",name,par)
			//params := Tokenize(par)
			//var data = &ShortcodeWithPage{Params: params, Page: p}
			newString := stringToParse[:posStart] + par + stringToParse[posEnd+3:]
			return newString
		}
	}
	return stringToParse
}
