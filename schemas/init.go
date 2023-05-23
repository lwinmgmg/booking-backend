package schemas

func AssignEmptyString(input **string) {
	if input == nil {
		defaultStr := ""
		*input = &defaultStr
	}
}
