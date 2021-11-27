package process

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTimestamp(line string) string {
	tokens := strings.Split(line, " ")
	return tokens[0]
}

func nextLine(reader *bufio.Reader) (res string, err error) {
	l, _, err := reader.ReadLine()
	return string(l), err
}

func CombineByTimestamp(fileName1 string, fileName2 string, outputFileName string) {
	f1, err := os.Open(fileName1)
	check(err)
	defer f1.Close()

	f2, err := os.Open(fileName2)
	check(err)
	defer f2.Close()

	f3, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY, 0666)
	check(err)
	defer f3.Close()

	r1 := bufio.NewReader(f1)
	r2 := bufio.NewReader(f2)

	l1, _ := nextLine(r1)
	l2, _ := nextLine(r2)

	t1 := getTimestamp(l1)
	t2 := getTimestamp(l2)

	for {
		if t2 > t1 {
			f3.WriteString("first " + l1 + "\n")
			l1, err = nextLine(r1)
			if err == io.EOF {
				break
			}
			t1 = getTimestamp(l1)
		} else {
			f3.WriteString("second " + l2 + "\n")
			l2, err = nextLine(r2)
			if err == io.EOF {
				break
			}
			t2 = getTimestamp(l2)
		}
	}
}
