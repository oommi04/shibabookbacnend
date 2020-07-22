package common

func IncludeString(strs []string, s string) bool{
	for _, item := range strs {
		if s == item {
			return true
		}
	}
	return false
}
