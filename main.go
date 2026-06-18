package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var currencies = map[string]float64{
	"USD": 1.00,
	"BRL": 5.10,
	"EUR": 0.87,
}

func convert(value float64, from, to string) (float64, error) {
	taxFrom, ok := currencies[from]
	if !ok {
		return 0, fmt.Errorf("unknown origin coin: %s", from)
	}
	taxTo, ok := currencies[to]
	if !ok {
		return 0, fmt.Errorf("unknown target coin: %s", to)
	}
	return value * taxTo / taxFrom, nil
}

func prompt(scanner *bufio.Scanner, out io.Writer, msg string) (string, bool) {
	fmt.Fprintln(out, msg)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		}
		return "", false
	}
	return scanner.Text(), true
}

func run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var value float64

	for {
		fmt.Fprintln(out, "--------------------------------------")
		fmt.Fprintln(out, "Go CLI Currency Exchange")
		fmt.Fprintln(out, "by Joaquim Prieto 2026")
		fmt.Fprintln(out, "")

		fmt.Fprintln(out, "(Type close to end the program at any time)")

		input, valid := prompt(scanner, out, "Type the value you want to exchange ->")
		if !valid {
			break
		}

		input = strings.ToUpper(input)
		switch input {
		case "":
			fmt.Fprintln(out, "Expected value")
			continue
		case "CLOSE":
			return
		default:
			var err error
			input = strings.Replace(input, ",", ".", 1)
			value, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Fprintln(out, "value must be a number")
				continue
			}
		}

		originCurr, valid := prompt(scanner, out, "Type the origin currency ->")
		if !valid {
			break
		}
		originCurr = strings.ToUpper(originCurr)
		if originCurr == "" {
			fmt.Fprintln(out, "Expected origin currency")
			continue
		} else if originCurr == "CLOSE" {
			return
		}

		targetCurr, valid := prompt(scanner, out, "Type the target currency ->")
		if !valid {
			break
		}
		targetCurr = strings.ToUpper(targetCurr)
		if targetCurr == "" {
			fmt.Fprintln(out, "Expected target currency")
			continue
		} else if targetCurr == "CLOSE" {
			return
		}

		result, err := convert(value, originCurr, targetCurr)
		if err != nil {
			fmt.Fprintln(out, "Error:", err)
			continue
		}

		fmt.Fprintf(out, "Result: %s %.2f\n", targetCurr, result)

		fmt.Fprintln(out, "--------------------------------------")
		fmt.Fprintln(out, "")
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}
