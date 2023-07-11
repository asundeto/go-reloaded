package reloaded

import "fmt"

func mainCheck(s string) string {
	res, arr := []rune{}, []rune{}
	for i := 0; i < len(s); i++ {
		arr = append(arr, rune(s[i]))
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == '(' {
			if i == 0 {
				res = append(res, arr[i])
				continue
			}
			iStart, numStr, num, checkWord, timeI, iPlus, slashN, checkMin := i, "", 1, "", i+1, 0, false, false
			f := true
			for f {
				if arr[timeI] >= 'a' && arr[timeI] <= 'z' {
					iPlus++
					checkWord += string(arr[timeI])
				} else if arr[timeI] == ',' {
					iPlus++
					f = false
					if timeI+2 < len(arr) {
						if arr[timeI+1] != ' ' {
							i = iStart
							res = append(res, arr[i])
							checkMin = true
							break
						}
						timeI += 2
					}
					for true {
						if arr[timeI] >= '0' && arr[timeI] <= '9' {
							numStr += string(arr[timeI])
						} else if arr[timeI] == '-' {
							i = iStart
							res = append(res, arr[i])
							checkMin = true
							break
						} else if arr[timeI] == ')' {
							break
						} else {
							i = iStart
							res = append(res, arr[i])
							checkMin = true
							break
						}
						timeI++
						iPlus++
					}
					iPlus++
					break
				} else if arr[timeI] == ')' {
					if timeI+1 < len(arr) {
						if arr[timeI+1] == 13 || arr[timeI+1] == '\n' {
							slashN = true
						}
					}
					iPlus++
					f = false
				} else {
					i = iStart
					break
				}
				timeI++
			}
			if checkMin {
				fmt.Println("Error! Incorrect format!")
				continue
			}
			if numStr != "" {
				num = Atoi(numStr)
			}
			i++
			if checkWord == "cap" {
				res = cap(res, len(res)-1, num, slashN)
				i += iPlus
			} else if checkWord == "up" {
				res = LowUp(res, len(res)-1, num, 1, slashN)
				i += iPlus
			} else if checkWord == "low" {
				res = LowUp(res, len(res)-1, num, 0, slashN)
				i += iPlus
			} else if checkWord == "bin" {
				res = HexBin(res, len(res)-1, 2, slashN)
				i += iPlus
			} else if checkWord == "hex" {
				res = HexBin(res, len(res)-1, 1, slashN)
				i += iPlus
			} else {
				res = append(res, arr[i-1])
				i = iStart
				continue
			}
		} else {
			res = append(res, arr[i])
		}
	}
	return RuneToStr(res)
}

func HexBin(s []rune, i, num int, slashN bool) []rune {
	numStr := ""
	for true {
		if s[i] != ' ' && s[i] != '\n' && s[i] != 13 {
			break
		} else {
			i--
		}
	}
	for true {
		if i > 0 {
			if s[i] != ' ' && s[i] != '\n' && s[i] != 13 {
				numStr += string(s[i])
				i--
			} else {
				i++
				break
			}
		} else if i == 0 {
			numStr += string(s[i])
			break
		}
	}
	numStr = myReverse(numStr)
	startLen := len(numStr)
	numStr = hexAndBinToDec(numStr, num)
	check := 0
	difference := startLen - len(numStr)
	if difference < 0 {
		check = 2
		difference *= -1
	} else {
		check = 1
	}
	if check == 2 {
		for j := 0; j < len(numStr); j++ {
			if i <= len(s) {
				s[i] = rune(numStr[j])
				if i == len(s) {
					s = append(s, rune(numStr[j]))
				}
				if j+1 == len(numStr) {
					s = append(s, ' ')
				}
				if i+1 <= len(s) {
					i++
				}
			}
		}
	} else if check == 1 {
		for j := 0; j < len(numStr); j++ {
			s[i] = rune(numStr[j])
			if i+1 <= len(s) {
				i++
			}
		}
		for j := 0; j < difference; j++ {
			s[i] = ' '
			if i+1 <= len(s) {
				i++
			}
		}
	} else {
		for j := 0; j < len(numStr); j++ {
			s[i] = rune(numStr[j])
			if i+1 <= len(s) {
				i++
			}
		}
	}
	if slashN {
		s = append(s, '\n')
	}
	return s
}

func cap(s []rune, i, num int, slashN bool) []rune {
	x := 0
	for true {
		if s[i] != ' ' && s[i] != '\n' && s[i] != 13 && (s[i] < '0' || s[i] > '9') {
			break
		} else {
			i--
		}
	}
	for x < num {
		if i > 0 {
			if s[i] == ' ' || s[i] == '\n' || s[i] == 13 {
				if i+1 < len(s) {
					if (s[i+1] >= 'a' && s[i+1] <= 'z') || (s[i+1] >= 'A' && s[i+1] <= 'Z') {
						s[i+1] = toUpperCase(s[i+1])
						x++
					}
				}
				for true {
					if s[i] != ' ' && s[i] != '\n' && s[i] != 13 {
						break
					} else {
						if i-1 >= 0 {
							i--
						} else {
							break
						}
					}
				}
			} else {
				s[i] = toLowerCase(s[i])
				i--
			}
		} else if i == 0 {
			if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
				s[i] = toUpperCase(s[i])
			} else {
				s[i+1] = toUpperCase(s[i+1])
			}
			break
		}
	}
	if slashN {
		s = append(s, '\n')
	}
	return s
}

func LowUp(s []rune, i, num, loc int, slashN bool) []rune {
	x := 0
	for true {
		if s[i] != ' ' && s[i] != '\n' && s[i] != 13 && (s[i] < '0' || s[i] > '9') {
			break
		} else {
			if i-1 >= 0 {
				i--
			} else {
				break
			}
		}
	}
	for x < num {
		if i > 0 {
			if s[i] != ' ' && s[i] != '\n' && s[i] != 13 {
				if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
					if loc == 1 {
						s[i] = toUpperCase(s[i])
					} else {
						s[i] = toLowerCase(s[i])
					}
				}
				i--
			} else {
				x++
				for true {
					if s[i] != ' ' && s[i] != '\n' && s[i] != 13 {
						break
					} else {
						i--
					}
				}
			}
		} else if i == 0 {
			if loc == 1 {
				s[i] = toUpperCase(s[i])
			} else {
				s[i] = toLowerCase(s[i])
			}
			break
		}
	}
	if slashN {
		s = append(s, '\n')
	}
	return s
}

func RuneToStr(r []rune) string {
	res := ""
	for i := 0; i < len(r); i++ {
		res += string(r[i])
	}
	return res
}
