package main

import (
	"bufio"
	"fmt"
	"hw3_bench/structs"
	"io"
	"os"
	"strings"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	seenBrowsers := make(map[string]interface{}, 32)

	fmt.Fprintln(out, "found users:")

	for i := 0; scanner.Scan(); i++ {
		userJsonBytes := scanner.Bytes()
		user := structs.User{}

		err := user.UnmarshalJSON(userJsonBytes)
		if err != nil {
			panic(err)
		}

		var hasMSIE, hasAndroid bool

		for _, browser := range user.Browsers {
			isMSIE := strings.Contains(browser, "MSIE")
			isAndroid := strings.Contains(browser, "Android")
			_, browserHasBeenAlreadySeen := seenBrowsers[browser]

			hasMSIE = hasMSIE || isMSIE // equivalent to "if isMSIE { hasMSIE = true }"
			hasAndroid = hasAndroid || isAndroid

			if !browserHasBeenAlreadySeen && (isMSIE || isAndroid) {
				seenBrowsers[browser] = struct{}{}
			}
		}

		if !(hasMSIE && hasAndroid) {
			continue
		}

		email := strings.ReplaceAll(user.Email, "@", " [at] ")
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}

func main() {
	//SlowSearch(os.Stdout)
	FastSearch(os.Stdout)
}
