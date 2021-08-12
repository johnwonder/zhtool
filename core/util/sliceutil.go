package util

func FindFirstMatchStr(s1 []string,s2 []string) string {
	if  s1 == nil || s2 == nil || len(s1) < 1 || len(s2) < 1 {
		return ""
	}

	//https://www.cnblogs.com/liuxingxing/p/14437084.html
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
