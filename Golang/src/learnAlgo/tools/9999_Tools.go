package tools

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
)

func generate_ints_file(path string, num int) {
	fileObj, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer fileObj.Close()
	w := bufio.NewWriter(fileObj)
	for i := 0; i < num; i++ {
		w.WriteString(strconv.Itoa(rand.Int()) + ",")
	}
	w.Flush()

}

func Gen_ints_list(num int) []int {
	data := make([]int, num)
	for i := 0; i < num; i++ {
		data[i] = rand.Int()
	}
	return data
}

func Main9999() {
	generate_ints_file("E://data/Learning/ints1000.txt", 1000)
}
