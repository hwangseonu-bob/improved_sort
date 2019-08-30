package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func loadData(filename string) []int {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	s := string(c)
	data := make([]int, 0)

	for _, v := range strings.Split(s, "\r\n") {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		data = append(data, i)
	}
	return data
}

func main() {
	t := time.Now()

	data := loadData(os.Args[1])

	QSort(data, 0, len(data))

	for _, v := range data {
		fmt.Printf("%d, ", v)
	}
	fmt.Println()

	qresult := time.Now().Sub(t).Nanoseconds()
	fmt.Printf("퀵 정렬 실행시간: %d ns\n", qresult)

	t = time.Now()
	data = loadData(os.Args[1])

	ISort(data)

	for _, v := range data {
		fmt.Printf("%d, ", v)
	}
	fmt.Println()

	fmt.Printf("수정된 정렬 실행시간: %d ns\n", time.Now().Sub(t).Nanoseconds())
}
