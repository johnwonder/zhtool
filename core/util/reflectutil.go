package util

import "time"

// InterfaceToString 接口转换为字符串
func InterfaceToString(i interface{}) string  {
	switch s := i.(type) {
	case string:
		return  s
	default:
		Error("Only Strings are supported")
	}

	return  ""
}

// InterfaceArrayToStringArray 接口转换为字符串数组
func InterfaceArrayToStringArray(i interface{}) []string {

	var a []string
	switch vv := i.(type) {
	case []interface{}:
		for _,u  := range vv{
			a = append(a,InterfaceToString(u))
		}
	}

	return  a
}

// InterfaceToStringDate 接口转换为时间格式
func InterfaceToStringDate(i interface{}) time.Time{
	s := InterfaceToString(i)
	d ,e := time.Parse("02 Jan 06 15:04 MST",s)

	if e != nil {
		d ,e = time.Parse("2006-01-02",s)
	}

	if e != nil {
		d ,e = time.Parse("02 Jan 06",s)
	}

	return  d
}

// InterfaceToBool 接口转换为 bool类型
func InterfaceToBool(i interface{}) bool {
	switch b := i.(type) {
	case bool:
		return  b
	default:
		Error("Only Boolean values are supported ")
	}
	return  false
}