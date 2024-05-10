package utils

func InList(list []string, key string) bool {
	for _, val := range list {
		if val == key {
			return true
		}
	}
	return false
}
