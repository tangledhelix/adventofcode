/*
 * --- Day 6: Custom Customs ---
 * As your flight approaches the regional airport where you'll switch to a much
 * larger plane, customs declaration forms are distributed to the passengers.
 *
 * The form asks a series of 26 yes-or-no questions marked a through z. All you
 * need to do is identify the questions for which anyone in your group answers
 * "yes". Since your group is just you, this doesn't take very long.
 *
 * However, the person sitting next to you seems to be experiencing a language
 * barrier and asks if you can help. For each of the people in their group, you
 * write down the questions for which they answer "yes", one per line. For
 * example:
 *
 * abcx
 * abcy
 * abcz
 *
 * In this group, there are 6 questions to which anyone answered "yes": a, b, c,
 * x, y, and z. (Duplicate answers to the same question don't count extra; each
 * question counts at most once.)
 *
 * Another group asks for your help, then another, and eventually you've
 * collected answers from every group on the plane (your puzzle input). Each
 * group's answers are separated by a blank line, and within each group, each
 * person's answers are on a single line. For example:
 *
 * abc
 *
 * a
 * b
 * c
 *
 * ab
 * ac
 *
 * a
 * a
 * a
 * a
 *
 * b
 *
 * This list represents answers from five groups:
 *
 * - The first group contains one person who answered "yes" to 3 questions: a, b,
 *   and c.
 * - The second group contains three people; combined, they answered "yes" to 3
 *   questions: a, b, and c.
 * - The third group contains two people; combined, they answered "yes" to 3
 *   questions: a, b, and c.
 * - The fourth group contains four people; combined, they answered "yes" to only
 *   1 question, a.
 * - The last group contains one person who answered "yes" to only 1 question, b.
 *
 *In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.
 *
 * For each group, count the number of questions to which anyone answered "yes".
 * What is the sum of those counts?
 */

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Read input file into an array of strings (one per line)
	dat, err := os.ReadFile("input.txt")
	check(err)
	rawData := strings.Split(string(dat), "\n")

	// Track the running total (this is a sum of the counts of the questions
	// asked in each group)
	totalCount := 0

	// On each turn we'll mark which questions have been answered, using this
	// map. Between turns, reset the map so it's all false again.
	letters := map[string]bool{
		"a": false, "b": false, "c": false, "d": false, "e": false,
		"f": false, "g": false, "h": false, "i": false, "j": false,
		"k": false, "l": false, "m": false, "n": false, "o": false,
		"p": false, "q": false, "r": false, "s": false, "t": false,
		"u": false, "v": false, "w": false, "x": false, "y": false,
		"z": false}

	for i := 0; i < len(rawData); i++ {
		line := rawData[i]

		// Blank line means a new record has started. Count up this group's
		// answers and add to the running total, then reset the data for next
		// round.
		if line == "" {
			groupCount := countGroupResponses(letters)
			totalCount += groupCount
			lettersReset(letters)
			continue
		}

		for _, c := range line {
			letters[string(c)] = true
		}

	}

	fmt.Printf("Total: %d\n", totalCount)

}

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

// Reset all letters to false (do this between groups)
func lettersReset(letters map[string]bool) {
	for key, _ := range letters {
		letters[key] = false
	}
}

// When we get to the end of a group, we want to tally their responses so we can
// add them to the total count.
func countGroupResponses(letters map[string]bool) int {
	count := 0
	for _, val := range letters {
		if val {
			count++
		}
	}

	return count
}
