package main

// this program inputs byte array from go program from stdin or argument and outputs the string representation of the byte array

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	var err error
	
	if len(os.Args) > 1 {
		result, err = convertByteArrayFromArg(os.Args[1:])
	} else {
		result, err = convertByteArrayFromStdin()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func convertByteArrayFromArg(args []string) (string, error) {
	s := strings.Join(args, " ")
	s = strings.Trim(s, "[")
	s = strings.Trim(s, "]")
	s = strings.Trim(s, "{")
	s = strings.Trim(s, "}")
	s = strings.Trim(s, " ")
	s = strings.ReplaceAll(s, " ", "")
	args = strings.Split(s, ",")

	var byteArray []byte
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return "", fmt.Errorf("invalid argument %q: %w", arg, err)
		}
		byteArray = append(byteArray, byte(num))
	}
	return string(byteArray), nil
}

func convertByteArrayFromStdin() (string, error) {
	var byteArray []byte
	var num int
	for {
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return "", fmt.Errorf("reading from stdin: %w", err)
			}

			break
		}
		byteArray = append(byteArray, byte(num))
	}
	return string(byteArray), nil
}
