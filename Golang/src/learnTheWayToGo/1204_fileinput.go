/*
接收控制台输入的文件路径，然后尝试打开文件，逐行打印。
*/

package learnTheWayToGo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func Main1204() {
	var path string
	var e error
	for e := errors.New(""); e != nil; {
		fmt.Println("Please enter file-path: ")
		_, e = fmt.Scanln(&path)
	}

	inputFile, e := os.Open(path)
	if e != nil {
		fmt.Printf("Failed when open file {%s}", path)
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, e := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if e == io.EOF {
			return
		}
	}
}

func Main1206() {
	var path string // C:\Users\lewin\mydata\测试.txt
	var e error
	for e := errors.New(""); e != nil; {
		fmt.Println("Please enter file-path: ")
		_, e = fmt.Scanln(&path)
	}

	inputFile, e := os.Open(path)
	if e != nil {
		fmt.Printf("Failed when open file {%s}", path)
		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(inputReader, &v1, &v2, &v3)
		// scans until newline
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
		if err == io.EOF {
			break
		}
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func Main1215() {
	// unbuffered
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered: os.Stdout implements io.Writer
	buf := bufio.NewWriter(os.Stdout)
	// and now so does buf.
	for i := 1000; i < 2000; i++ {
		fmt.Fprintf(buf, "%d\n", i)

	}
	time.Sleep(time.Second * 2)
	//buf.Flush()
}

func Exer1207() {
	inputFile, _ := os.Open("C:/Users/lewin/mydata/测试.txt")
	outputFile, _ := os.OpenFile("C:/Users/lewin/mydata/测试2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	defer outputWriter.Flush()
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("Conversion done")
}
