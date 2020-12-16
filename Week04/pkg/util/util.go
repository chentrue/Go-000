package util


func AddExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}