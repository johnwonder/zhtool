package util

func FindFirstMatchStr(s1 []string,s2 []string) string {
	if  s1 == nil || s2 == nil || len(s1) < 1 || len(s2) < 1 {
		return ""
	}

	//https://www.cnblogs.com/liuxingxing/p/14437084.html
	//https://github.com/wxnacy/wgo/blob/a17b9d43fb9464b43937bd93db10239fe87fcd9b/arrays/contains.go#L25
	set := make(map[string]struct{})
	for _, value := range s2{
		set[value] = struct{}{}
	}

	for _,v := range s1 {
		if _, ok := set[v];ok {
			return  v
		}
	}
	return ""
}

func ContainsStr(arr []string, val string) (index int) {
	index = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			index = i
			return
		}
	}
	return
}
