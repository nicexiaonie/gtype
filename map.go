package gtype

func StringMapToStringMapString(v map[string]interface{}) map[string]string {
	ret := make(map[string]string)
	for k, v := range v {
		ret[k] = ToString(v)
	}
	return ret
}
func StringMapToStringMapInt64(v map[string]interface{}) map[string]int64 {
	ret := make(map[string]int64)
	for k, v := range v {
		ret[k] = ToInt64(v)
	}
	return ret
}

