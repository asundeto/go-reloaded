package reloaded

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadWholeFile(fileName string) string {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return "Error"
	}
	res := checkArticle(checkQuotes(checkSpace(checkSign(mainCheck(checkSlash(string(contents)))))))
	return res
}

func WriteToFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		fmt.Println(err.Error())
		return
	}
}

func CheckTxt(s string) bool {
	if len(s) < 4 {
		return true
	}
	rev := myReverse(s)
	if rev[:4] != "txt." {
		return true
	}
	return false
}
