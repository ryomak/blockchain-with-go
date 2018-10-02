package util
func Contains(s []interface{}, e interface{}) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}