package main

import (
	"bufio"
	"fmt"
	"hw3_bench/structs"
	"io"
	"os"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		userJsonBytes := scanner.Bytes()
		user := structs.User{}

		err := user.UnmarshalJSON(userJsonBytes)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(out, user)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
}

func main() {
	FastSearch(os.Stdout)
}
