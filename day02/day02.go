/*
 * --- Day 2: Password Philosophy ---
 * Your flight departs in a few days from the coastal airport; the easiest way
 * down to the coast from here is via toboggan.
 *
 * The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day.
 * "Something's wrong with our computers; we can't log in!" You ask if you can
 * take a look.
 *
 * Their password database seems to be a little corrupted: some of the passwords
 * wouldn't have been allowed by the Official Toboggan Corporate Policy that was
 * in effect when they were chosen.
 *
 * To try to debug the problem, they have created a list (your puzzle input) of
 * passwords (according to the corrupted database) and the corporate policy when
 * that password was set.
 *
 * For example, suppose you have the following list:
 *
 * 1-3 a: abcde
 * 1-3 b: cdefg
 * 2-9 c: ccccccccc
 *
 * Each line gives the password policy and then the password. The password
 * policy indicates the lowest and highest number of times a given letter must
 * appear for the password to be valid. For example, 1-3 a means that the
 * password must contain a at least 1 time and at most 3 times.
 *
 * In the above example, 2 passwords are valid. The middle password, cdefg, is
 * not; it contains no instances of b, but needs at least 1. The first and third
 * passwords are valid: they contain one a or nine c, both within the limits of
 * their respective policies.
 *
 * How many passwords are valid according to their policies?
 *
 * --- Part Two ---
 * While it appears you validated the passwords correctly, they don't seem to be
 * what the Official Toboggan Corporate Authentication System is expecting.
 *
 * The shopkeeper suddenly realizes that he just accidentally explained the
 * password policy rules from his old job at the sled rental place down the
 * street! The Official Toboggan Corporate Policy actually works a little
 * differently.
 *
 * Each policy actually describes two positions in the password, where 1 means
 * the first character, 2 means the second character, and so on. (Be careful;
 * Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of
 * these positions must contain the given letter. Other occurrences of the
 * letter are irrelevant for the purposes of policy enforcement.
 *
 * Given the same example list from above:
 *
 * 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
 * 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
 * 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
 *
 * How many passwords are valid according to the new interpretation of the policies?
 */

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

// use an empty interface, so we can pass a random series of
// values as arguments, just like fmt.Println accepts
func debug(a ...interface{}) {
	// set true or false to control debug output
	var debug bool = false

	if debug {
		fmt.Println("DEBUG:", a)
	}
}

func main() {

	// Read input file and break into lines
	dat, err := os.ReadFile("input.txt")
	check(err)

	raw_data := strings.Split(string(dat), "\n")

	// Keep track how many passwords are valid
	validPasswordCount := 0

	debug("looping over each line of file...")

	for i := 0; i < len(raw_data); i++ {
		// sample of input data:
		// 2-3 b: bkkb
		// <low_bound>-<high_bound> <letter>: <password>

		var low, high int
		var letter, password string

		// We need to split left and right side because the %s consumes the :
		// and therefore the fmt string doesn't match.
		// This could also be solved using range or similar to iterate over
		// the line character by character, but that's tedious.
		chunks := strings.Split(string(raw_data[i]), ":")
		debug("chunks:", chunks)

		if len(chunks) == 2 {
			debug("chunks[0]:", chunks[0])
			debug("chunks[1]:", chunks[1])
			_, err := fmt.Sscanf(chunks[0], "%d-%d %s", &low, &high, &letter)
			check(err)
			_, err = fmt.Sscanf(chunks[1], "%s", &password)
			check(err)
		}

		// Determine how many times letter occurs in password
		numberOccurrences := 0
		for _, c := range password {
			if string(c) == letter {
				numberOccurrences++
			}
		}
		debug("numberOccurrences=", numberOccurrences)

		// If this number of occurrences is valid, increase the counter
		// Ignore if the number of occurrences was 0
		if numberOccurrences > 0 && numberOccurrences >= low && numberOccurrences <= high {
			validPasswordCount++
			debug(fmt.Sprintf("counting this as VALID (now %d)", validPasswordCount))
		}
	}

	// Print the final count
	fmt.Println("Answer 1:", validPasswordCount)

	// Now let's do this again with new rules for Part Two.
	validPasswordCount = 0

	for i := 0; i < len(raw_data); i++ {
		var pos1, pos2 int
		var letter, password string

		// We need to split left and right side because the %s consumes the :
		// and therefore the fmt string doesn't match.
		// This could also be solved using range or similar to iterate over
		// the line character by character, but that's tedious.
		chunks := strings.Split(string(raw_data[i]), ":")
		debug("chunks:", chunks)

		if len(chunks) == 2 {
			debug("chunks[0]:", chunks[0])
			debug("chunks[1]:", chunks[1])
			_, err := fmt.Sscanf(chunks[0], "%d-%d %s", &pos1, &pos2, &letter)
			check(err)
			_, err = fmt.Sscanf(chunks[1], "%s", &password)
			check(err)
		}

		// check the two positions and make sure only ONE contains
		// the letter. If neither or both do, then it fails.
		numberOccurrences := 0
		for n, c := range password {
			if n+1 == pos1 && string(c) == letter {
				numberOccurrences++
			}
			if n+1 == pos2 && string(c) == letter {
				numberOccurrences++
			}
		}

		// If only ONE column contained the letter, good. Else bad.
		if numberOccurrences == 1 {
			validPasswordCount++
		}
	}

	// Print the final (final) count
	fmt.Println("Answer 2:", validPasswordCount)

}
