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
			fmt.Fprintf(os.Stderr, "Missing data in reducer input: %s\n", line)
			continue
		}

		personID, addPlayedMovies, addDirectedMovies := fields[0], fields[1], fields[2]

		if currentPersonID == "" {
			currentPersonID = personID
		} else if currentPersonID != personID {
			printReducedData(currentPersonID, playedMoviesCounter, directedMoviesCounter)
			currentPersonID = personID
			playedMoviesCounter = 0
			directedMoviesCounter = 0
		}

		playedMovieInt, playedCastErr := strconv.Atoi(addPlayedMovies)
		if playedCastErr != nil {
			fmt.Fprintf(os.Stderr, "Failed conversion to int: %v\n", playedCastErr)
			continue
		}
		if playedMovieInt > 0 {
			playedMoviesCounter += playedMovieInt
		}

		directedMovieInt, directedCastErr := strconv.Atoi(addDirectedMovies)
		if directedCastErr != nil {
			fmt.Fprintf(os.Stderr, "Failed conversion to int: %v\n", directedCastErr)
			continue
		}
		if directedMovieInt > 0 {
			directedMoviesCounter += directedMovieInt
		}
	}

	if currentPersonID != "" {
		printReducedData(currentPersonID, playedMoviesCounter, directedMoviesCounter)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading reducer input:", err)
	}
}

func printReducedData(currentPersonID string, playedMoviesCounter int, directedMoviesCounter int) {
	fmt.Printf("%s\t%d\t%d\n", currentPersonID, playedMoviesCounter, directedMoviesCounter)
}
