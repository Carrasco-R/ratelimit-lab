package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var argMap = map[string]func() string{
	"rate":       testRate,
	"concurrent": testConcurrent,
}

func main() {
	arg, err := getArg(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	for k := range argMap {
		fmt.Println(k)
	}
	argFn := argMap[arg]
	argFn()
}

func getArg(args []string) (string, error) {
	argLen := len(args)
	if argLen == 1 {
		return "", errors.New("no arguement given")
	} else if argLen > 2 {
		return "", errors.New("too many args")
	}
	arg := os.Args[1]
	return arg, nil
}

func testRate() string {
	fmt.Println("testRate")  
	return ""
}

func testConcurrent() string {
	fmt.Println("testRate")
	return ""
}
