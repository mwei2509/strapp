package app

func wrapLog(args ...any) []any {
	colorReset := "\033[0m"

	// colorRed := "\033[31m"
	// colorGreen := "\033[32m"
	// colorYellow := "\033[33m"
	// colorBlue := "\033[34m"
	// colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	// colorWhite := "\033[37m"
	logs := []any{string(colorCyan), "[[ STRAPP ]]", string(colorReset)}
	logs = append(logs, args...)
	return logs
}
