package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrIsBlank(t *testing.T) {

	assert.Equal(t, true, StrIsBlank(""))

	assert.Equal(t, true, StrIsBlank(" "))

	//https://segmentfault.com/q/1010000010025001
	var s string
	assert.Equal(t, true, StrIsBlank(s))
	//报错
	//assert.Equal(t, true, StrIsBlank(nil))

	assert.Equal(t, true, StrIsBlank("\n \n"))
	assert.Equal(t, true, StrIsBlank("\t\n\n\n\t \n\t "))

	hex := "\u3000"
	fmt.Print("输出前") //输出“”
	fmt.Print(hex) //输出“”
	fmt.Print("输出") //输出“”
	assert.Equal(t, true, StrIsBlank("\u3000"))

}

func TestStrIsEmpty(t *testing.T) {

	assert.Equal(t, true, StrIsEmpty(""))

	assert.Equal(t, false, StrIsEmpty(" "))

	assert.Equal(t, false, StrIsEmpty("\t"))

}
