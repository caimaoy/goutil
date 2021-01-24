package slice

//StringInSlice returns a string in a string scile or not
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//DeleteStringSlice remove a item from the slice
func DeleteStringSlice(a []string, index int) []string {
	a = append(a[:index], a[index+1:]...)
	return a
}

//StringSliceEqule return if two string slice equle or not
func StringSliceEqule(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
