/* --- Day 1: Report Repair ---
 *
 * After saving Christmas five years in a row, you've decided to take a vacation
 * at a nice resort on a tropical island. Surely, Christmas will go on without
 * you.
 *
 * The tropical island has its own currency and is entirely cash-only. The gold
 * coins used there have a little picture of a starfish; the locals just call
 * them stars. None of the currency exchanges seem to have heard of them, but
 * somehow, you'll need to find fifty of these coins by the time you arrive so
 * you can pay the deposit on your room.
 *
 * To save your vacation, you need to get all fifty stars by December 25th.
 *
 * Collect stars by solving puzzles. Two puzzles will be made available on each
 * day in the Advent calendar; the second puzzle is unlocked when you complete
 * the first. Each puzzle grants one star. Good luck!
 *
 * Before you leave, the Elves in accounting just need you to fix your expense
 * report (your puzzle input); apparently, something isn't quite adding up.
 *
 * Specifically, they need you to find the two entries that sum to 2020 and then
 * multiply those two numbers together.
 *
 * For example, suppose your expense report contained the following:
 *
 * 1721
 * 979
 * 366
 * 299
 * 675
 * 1456
 *
 * In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying
 * them together produces 1721 * 299 = 514579, so the correct answer is 514579.
 *
 * Of course, your expense report is much larger. Find the two entries that sum
 * to 2020; what do you get if you multiply them together?
 *
 * ------------- PART TWO ----------
 * The Elves in accounting are thankful for your help; one of them even offers
 * you a starfish coin they had left over from a past vacation. They offer you a
 * second one if you can find three numbers in your expense report that meet the
 * same criteria.
 *
 * Using the above example again, the three entries that sum to 2020 are 979,
 * 366, and 675. Multiplying them together produces the answer, 241861950.
 *
 * In your expense report, what is the product of the three entries that sum to
 * 2020?
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

func main() {

	// This is given (for this problem) because we know the input length
	const expenseCount int = 200

	// Read input file
	dat, err := os.ReadFile("input.txt")
	check(err)

	// Convert string contents to integers
	raw_expenses := strings.Split(string(dat), "\n")
	var expenses [expenseCount]int

	for i := 0; i < len(raw_expenses); i++ {
		var num int
		n, err := fmt.Sscanf(raw_expenses[i], "%d", &num)
		check(err)

		// Store a number, but only if one was returned.
		// If we are at EOF, then n will be 0.
		if n == 1 {
			expenses[i] = num
		}
	}

	// Part 1 - look for a pair of numbers which sum to 2020, and return their
	// product.

	var found bool = false

	for i := 0; i < len(expenses) && !found; i++ {
		for j := 0; j < len(expenses) && !found; j++ {
			if expenses[i]+expenses[j] == 2020 {
				fmt.Println("Answer 1 =", expenses[i]*expenses[j])
				found = true
			}
		}
	}

	// Repeat for part 2 - now looking for 3 numbers that sum to 2020. Again,
	// return their product.

	found = false

	for i := 0; i < len(expenses) && !found; i++ {
		for j := 0; j < len(expenses) && !found; j++ {
			for k := 0; k < len(expenses) && !found; k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					fmt.Println("Answer 2 =", expenses[i]*expenses[j]*expenses[k])
					found = true
				}
			}
		}
	}
}
