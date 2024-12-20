package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		currentPersonID       string
		playedMoviesCounter   int = 0
		directedMoviesCounter int = 0
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		if len(fields) < 3 {
			fmt.Fprintf(os.Stderr, "Missing data in combiner input: %s\n", line)
			continue
		}

		personID, addPlayedFlag, addDirectedFlag := fields[0], fields[1], fields[2]

		if currentPersonID == "" {
			currentPersonID = personID
		} else if currentPersonID != personID {
			printCombinedData(currentPersonID, playedMoviesCounter, directedMoviesCounter)
			currentPersonID = personID
			playedMoviesCounter = 0
			directedMoviesCounter = 0
		}

		playedMovieInt, playedCastErr := strconv.Atoi(addPlayedFlag)
		if playedCastErr != nil {
			fmt.Fprintf(os.Stderr, "Failed conversion to int: %v\n", playedCastErr)
			continue
		}
		if playedMovieInt == 1 {
			playedMoviesCounter++
		}

		directedMovieInt, directedCastErr := strconv.Atoi(addDirectedFlag)
		if directedCastErr != nil {
			fmt.Fprintf(os.Stderr, "Failed conversion to int: %v\n", directedCastErr)
			continue
		}
		if directedMovieInt == 1 {
			directedMoviesCounter++
		}
	}

	if currentPersonID != "" {
		printCombinedData(currentPersonID, playedMoviesCounter, directedMoviesCounter)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading combiner input:", err)
	}
}

func printCombinedData(currentPersonID string, playedMoviesCounter int, directedMoviesCounter int) {
	fmt.Printf("%s\t%d\t%d\n", currentPersonID, playedMoviesCounter, directedMoviesCounter)
}
