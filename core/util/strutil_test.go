package util

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

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

func TestReverse(t *testing.T) {
	assert.Equal(t, "dcba", Reverse("abcd"))
	assert.Equal(t, "谁是我", Reverse("我是谁"))

	assert.Equal(t, "dcba谁是我", Reverse("我是谁abcd"))

	assert.Equal(t, "", Reverse(""))

}

func TestReplaceTemplate(t *testing.T) {

	in := "111 {{% tem john %}}"
	t.Log(strings.TrimSpace(in))
	t.Log(strings.IndexFunc(strings.TrimSpace(" tem john "), unicode.IsSpace))
	posStart := strings.Index(in, "{{%")
	t.Logf("posStart:%v",posStart)
	posEnd := strings.Index(in[posStart:], "%}}") + posStart
	t.Logf("posEnd:%v", posEnd)

	//https://www.cnblogs.com/liubiaos/p/9367504.html
	//https://blog.csdn.net/KimBing/article/details/113584176
	t.Logf("去除{{%% %%}}后的字符串为：%v",in[posStart+3 : posEnd])
	assert.Equal(t, "我是johnwonder", ReplaceTemplate("我是{{% tem john %}}wonder","{{%","%}}"))

}

func TestSplitParams(t *testing.T) {
	name, par := SplitParams("img src=\"/media/spf13.jpg\" title=\"Steve Francia\"")

	t.Log(name)
	t.Log(par)
}

func TestTokenize(t *testing.T) {

	name, par := SplitParams("img src=&ldquo;/media/spf13.jpg&rdquo; title=&ldquo;Steve Francia&rdquo;")

	s := "s="
	t.Log(strings.Index(s,"="))
	ldquo := "\"ss"
	t.Log(strings.HasPrefix(ldquo, "\""))

	//img
	t.Log(name)
	//src="/media/spf13.jpg" title="Steve Francia"
	t.Log(par)


	first := strings.Fields(par)
	//输出3
	t.Logf("长度为:%v",len(first))
	for i ,v := range first {
		if i == len(first) {
			t.Log(v)
			t.Log(v)
		}else {
			t.Logf("索引为:%v,值为%v",i,v)
		}
	}

	t.Log(Tokenize(par))
}

func TestTrimRight(t *testing.T) {

	domain:= "/sss/"
	//从后面开始截取到 发现么有为止
	t.Log(strings.TrimRight(domain, "/s"))
	//从前面开始截取到 发现么有为止
	t.Log(strings.TrimLeft(domain, "/"))
}
