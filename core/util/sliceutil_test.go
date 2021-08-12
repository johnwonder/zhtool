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
