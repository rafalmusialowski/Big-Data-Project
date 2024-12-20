package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	personIDFieldIndex  int    = 2
	roleFieldIndex      int    = 3
	signalPlayedMovie   string = "1\t0"
	signalDirectedMovie string = "0\t1"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		role := fields[roleFieldIndex]
		personID := fields[personIDFieldIndex]

		switch role {
		case "self", "actor", "actress":
			fmt.Println(personID + "\t" + signalPlayedMovie)
		case "director":
			fmt.Println(personID + "\t" + signalDirectedMovie)
		default:
			//ignore other roles
		}
	}
}
