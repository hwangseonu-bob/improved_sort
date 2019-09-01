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

func improved(data []int) int64 {
	t := time.Now()
	ISort(data)

	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println()

	return time.Now().Sub(t).Nanoseconds()
}

func quick(data []int) int64 {
	t := time.Now()
	QSort(data, 0, len(data) - 1)

	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println()

	return time.Now().Sub(t).Nanoseconds()
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(-1)
	}

	i := improved(loadData(os.Args[1]))
	q := quick(loadData(os.Args[1]))

	fmt.Printf("개선된 정렬 실행시간: %d microseconds \n", i / 1000)
	fmt.Printf("퀵정렬 실행시간: %d microseconds\n", q / 1000)

	fmt.Printf("개선된 시간: %d microseconds \n", (q - i) / 1000)
}
