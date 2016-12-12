package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import (
	"strconv"
	"time"
	"chapter02/algorithms/qsort"
	"chapter02/algorithms/bubblesort"
	"chapter02/algorithms/mergesort"
	"chapter02/algorithms/selectionsort"
)

//快速解析命令行参数
var infile *string = flag.String("i", "unsorted.txt", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.txt", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input files", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error  {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return  err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return  nil

}

func main() {
	flag.Parse()

	fmt.Println("This is a sort app")

	if infile != nil {
		fmt.Println("infile =", *infile, "\noutfile =", *outfile, "\nalgorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		case "mergesort":
			mergesort.MergeSort(values)
		case "selectionsort":
			selectionsort.SelectionSort(values)
		default:
			fmt.Println("Sourting algorithm", *algorithm, "is unknow")
		}

		t2 := time.Now()
		fmt.Println("The sorting process consts", t2.Sub(t1), "to complete")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
