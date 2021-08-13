package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFirstMatchStr(t *testing.T) {

	//https://blog.csdn.net/972301/article/details/89318703
	a := []string{"1","2","3","4"}
	b := []string{"5","6","4","3"}
	assert.Equal(t, "3", FindFirstMatchStr(a,b))
}

func TestContainsStr(t *testing.T) {
	b := []string{"5","6","4","3"}
	assert.Equal(t, 0, ContainsStr(b,"5"))
	assert.Equal(t, -1, ContainsStr(b,"10"))


}
