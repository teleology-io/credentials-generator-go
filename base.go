package generator

import "strings"

func numbers() []string {
	return strings.Split("0123456789", "")
}

func lowercase() []string {
	return strings.Split("abcdefghijklmnopqrstuvwxyz", "")
}

func uppercase() []string {
	return strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
}

func special() []string {
	return []string{
		"@",
		"%",
		"+",
		"\\",
		"/",
		"\"",
		"!",
		"#",
		"$",
		"^",
		"?",
		":",
		".",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"~",
		"`",
		"-",
		"_",
		".",
	}
}
