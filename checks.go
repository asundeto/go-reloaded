package reloaded

import "fmt"

func checkSign(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
		} else {
			if i+1 < len(s) && s[i+1] != ',' && s[i+1] != '.' && s[i+1] != '!' && s[i+1] != '?' && s[i+1] != ' ' && s[i+1] != '\n' && s[i+1] != ':' {
				res += string(s[i])
			}
		}
	}
	return res
}

func checkSlashSpaces(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		if s[i] == '\n' || s[i] == 13 {
			if i+1 < len(s) {
				if s[i+1] != '\n' && s[i+1] != 13 {
					res += string(s[i])
				}
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}

func checkQuotes(s string) string {
	res, num := "", 0
	for i := 0; i < len(s); i++ {
		if i+1 != len(s) {
			if num%2 != 0 && s[i] == ' ' && (s[i+1] == '\'' || s[i+1] == '"') {
				continue
			}
			if s[i] == '\'' || s[i] == '"' || s[i] == '`' {
				if i != 0 && i+1 < len(s) && s[i-1] == 'n' && s[i+1] == 't' {
					res += string(s[i])
					continue
				}
				if s[i+1] == '\'' || s[i+1] == '"' || s[i+1] == '`' {
					res += string(s[i])
					continue
				}
				num++
				if num != 0 {
					if num%2 == 0 { // CLOSE QUOTES
						res += string(s[i])
						if i+1 != len(s) && s[i+1] != ' ' {
							res += " "
						}
					} else {
						if i != 0 && s[i-1] != ' ' && s[i-1] != '\n' && s[i-1] != 13 { // OPEN QUOTES
							res += " "
						}
						res += string(s[i])
						for true {
							if s[i+1] == ' ' || s[i+1] == 13 || s[i+1] == '\n' {
								i++
							} else {
								break
							}
						}
					}
				}
			} else {
				res += string(s[i])
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}

func checkSpace(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ',' || s[i] == '.' || s[i] == '!' || s[i] == '?' {
			if i+1 != len(s) {
				if i+2 < len(s) && s[i] == '.' && s[i+1] == '.' && s[i+2] == '.' {
					res += "... "
					i += 2
				} else if (s[i] == '!' && s[i+1] == '?') || (s[i] == '!' && s[i+1] == '!') {
					res += string(s[i]) + string(s[i+1])
					i++
				} else if s[i+1] != ' ' && s[i+1] != '.' && s[i+1] != ',' && s[i+1] != '?' && s[i+1] != '!'  {
					res += string(s[i])
					res += " "
				} else {
					res += string(s[i])
				}
			} else {
				res += string(s[i])
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}

func checkArticle(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i != 0 && i+2 <= len(s) {
			if (s[i-1] == ' ' || s[i-1] == 13 || s[i-1] == '\n') && (s[i] == 'a' || s[i] == 'A') && (s[i+1] == ' ' || s[i+1] == 13 || s[i+1] == '\n') {
				if s[i+2] == 'a' || s[i+2] == 'y' || s[i+2] == 'e' || s[i+2] == 'o' || s[i+2] == 'i' || s[i+2] == 'h' || s[i+2] == 'A' || s[i+2] == 'Y' || s[i+2] == 'E' || s[i+2] == 'O' || s[i+2] == 'I' || s[i+2] == 'H' {
					res += string(s[i])
					res += "n"
				} else {
					res += string(s[i])
				}
			} else {
				res += string(s[i])
			}
		} else {
			if i+2 < len(s) && (s[i] == 'a' || s[i] == 'A') && s[i+1] == ' ' {
				if s[i+2] == 'a' || s[i+2] == 'y' || s[i+2] == 'e' || s[i+2] == 'o' || s[i+2] == 'i' || s[i+2] == 'h' || s[i+2] == 'A' || s[i+2] == 'Y' || s[i+2] == 'E' || s[i+2] == 'O' || s[i+2] == 'I' || s[i+2] == 'H' {
					res += string(s[i])
					res += "n"
				} else {
					res += string(s[i])
				}
			} else {
				res += string(s[i])
			}
		}
	}
	return res
}

func checkSlash(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] == ' ' {
				for true {
					if i+1 < len(s) {
						if s[i+1] != ' ' {
							break
						}
						i++
					}
				}
			} else {
				res += string(s[i])
			}
		} else if s[i] == '\n' || s[i] == 13 {
			res += string(s[i])
			for true {
				if i+1 < len(s) {
					if s[i+1] != ' ' {
						break
					}
					i++
				}
			}
		} else {
			res += string(s[i])
		}
	}
	return res
}
