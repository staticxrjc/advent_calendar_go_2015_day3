package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please enter filename")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	santa := NewSantaWithRobot()
	for scanner.Scan() {
		line := scanner.Text()
		santa.Route(line)
	}

	fmt.Printf("Santa delivered presents to %d houses\n", santa.uniqueDeliveries)
}
