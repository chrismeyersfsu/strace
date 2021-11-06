package process

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CombineByTimestamp(fileName1 string, fileName2 string) {
	f1, err := os.Open(fileName1)
	check(err)

	defer f1.Close()
	f2, err := os.Open(fileName2)
	check(err)

	defer f2.Close()
}
