package pagerank

func isExist(s []interface{}, target interface{}) bool {
	for _, v := range s {
		if target == v {
			return true
		}
	}
	return false
}
