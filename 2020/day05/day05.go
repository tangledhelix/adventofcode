/*
 * --- Day 5: Binary Boarding ---
 * You board your plane only to discover a new problem: you dropped your
 * boarding pass! You aren't sure which seat is yours, and all of the flight
 * attendants are busy with the flood of people that suddenly made it through
 * passport control.
 *
 * You write a quick program to use your phone's camera to scan all of the
 * nearby boarding passes (your puzzle input); perhaps you can find your seat
 * through process of elimination.
 *
 * Instead of zones or groups, this airline uses binary space partitioning to
 * seat people. A seat might be specified like FBFBBFFRLR, where F means
 * "front", B means "back", L means "left", and R means "right".
 *
 * The first 7 characters will either be F or B; these specify exactly one of
 * the 128 rows on the plane (numbered 0 through 127). Each letter tells you
 * which half of a region the given seat is in. Start with the whole list of
 * rows; the first letter indicates whether the seat is in the front (0 through
 * 63) or the back (64 through 127). The next letter indicates which half of
 * that region the seat is in, and so on until you're left with exactly one row.
 *
 * For example, consider just the first seven characters of FBFBBFFRLR:
 *
 * Start by considering the whole range, rows 0 through 127.
 * F means to take the lower half, keeping rows 0 through 63.
 * B means to take the upper half, keeping rows 32 through 63.
 * F means to take the lower half, keeping rows 32 through 47.
 * B means to take the upper half, keeping rows 40 through 47.
 * B keeps rows 44 through 47.
 * F keeps rows 44 through 45.
 * The final F keeps the lower of the two, row 44.
 *
 * The last three characters will be either L or R; these specify exactly one of
 * the 8 columns of seats on the plane (numbered 0 through 7). The same process
 * as above proceeds again, this time with only three steps. L means to keep the
 * lower half, while R means to keep the upper half.
 *
 * For example, consider just the last 3 characters of FBFBBFFRLR:
 *
 * Start by considering the whole range, columns 0 through 7.
 * R means to take the upper half, keeping columns 4 through 7.
 * L means to take the lower half, keeping columns 4 through 5.
 * The final R keeps the upper of the two, column 5.
 * So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.
 *
 * Every seat also has a unique seat ID: multiply the row by 8, then add the
 * column. In this example, the seat has ID 44 * 8 + 5 = 357.
 *
 * Here are some other boarding passes:
 *
 * BFFFBBFRRR: row 70, column 7, seat ID 567.
 * FFFBBBFRRR: row 14, column 7, seat ID 119.
 * BBFFBBFRLL: row 102, column 4, seat ID 820.
 *
 * As a sanity check, look through your list of boarding passes. What is the
 * highest seat ID on a boarding pass?
 *
 * --- Part Two ---
 * Ding! The "fasten seat belt" signs have turned on. Time to find your seat.
 *
 * It's a completely full flight, so your seat should be the only missing
 * boarding pass in your list. However, there's a catch: some of the seats at
 * the very front and back of the plane don't exist on this aircraft, so they'll
 * be missing from your list as well.
 *
 * Your seat wasn't at the very front or back, though; the seats with IDs +1 and
 * -1 from yours will be in your list.
 *
 * What is the ID of your seat?
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

// A type defining a seat's metadata
type seat struct {
	locationCode string /* The [BFLR]+ code describing this seat location */
	row          int    /* Row number */
	col          int    /* Column number */
	id           int    /* Unique seat ID */
}

// A type defining a grid size
type gridSpec struct {
	rows int /* total rows in the grid */
	cols int /* total cols in the grid */
}

func findRowOrCol(seatLocator string, highBound int, lowLetter string, highLetter string) int {
	lowBound := 0
	highBound--

	for _, c := range seatLocator {
		// Find the difference between the current bounds and divide the space in half.
		diff := (highBound - lowBound + 1) / 2

		// Adjust the bounds as needed, depending on whether we're aiming low or high.
		// If the difference is 1, we're down to two adjacent values. In that case,
		// just pick either lower or higher and return the final answer.
		if string(c) == lowLetter {
			if diff == 1 {
				return lowBound
			}
			highBound = lowBound + diff - 1
		} else if string(c) == highLetter {
			if diff == 1 {
				return highBound
			}
			lowBound = highBound - diff + 1
		}

	}

	// We should never get here
	return -1
}

func findRow(seatLocator string, grid gridSpec) int {
	return findRowOrCol(seatLocator, grid.rows, "F", "B")
}

func findCol(seatLocator string, grid gridSpec) int {
	return findRowOrCol(seatLocator, grid.cols, "L", "R")
}

func findSeatLocation(seatLocator string, grid gridSpec) (int, int) {
	row := findRow(seatLocator, grid)
	col := findCol(seatLocator, grid)
	return row, col
}

func calculateSeatId(row, col int) int {
	return (row * 8) + col
}

// Process a single seat, populating its metadata in seatList
func processSeat(entryNumber int, seatLocator string, s *seat, grid gridSpec, seatMap *[128][8]bool) {
	if seatLocator != "" {
		s.locationCode = seatLocator
		s.row, s.col = findSeatLocation(seatLocator, grid)
		seatMap[s.row][s.col] = true
		s.id = calculateSeatId(s.row, s.col)
	}
}

// Locate the highest seat ID on our list
func findHighestSeatId(seatList *[1000]seat) int {
	var highest int

	for i := 0; i < len(seatList); i++ {
		if seatList[i].id > highest {
			highest = seatList[i].id
		}
	}

	return highest
}

func printSeatMap(seatMap *[128][8]bool) {
	for row := 0; row < len(seatMap); row++ {
		for col := 0; col < len(seatMap[row]); col++ {
			if seatMap[row][col] {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("  %03d\n", row)
	}
}

func findMySeatId(seatMap *[128][8]bool) int {
	for row := 0; row < len(seatMap); row++ {
		for col := 0; col < len(seatMap[row]); col++ {
			if col != len(seatMap[row])-1 {
				if !seatMap[row][col] && seatMap[row][col+1] && seatMap[row][col-1] {
					return calculateSeatId(row, col)
				}
			}
		}
	}
	return -1
}

func main() {
	var seatList [1000]seat
	grid := gridSpec{rows: 128, cols: 8}
	var seatMap [128][8]bool

	// Read input file and break into lines
	dat, err := os.ReadFile("input.txt")
	check(err)
	rawData := strings.Split(string(dat), "\n")

	for i := 0; i < len(rawData); i++ {
		processSeat(i, string(rawData[i]), &seatList[i], grid, &seatMap)
	}

	// Show a literal map of the plane so we can spot our seat.
	// printSeatMap(&seatMap)

	// Once seatList has been filled with the appropriate
	// data, iterate over every boarding pass to find the highest
	// seat ID in the list.
	highestSeatId := findHighestSeatId(&seatList)
	fmt.Println("Highest seat ID:", highestSeatId)

	// Show us our own seat
	mySeatId := findMySeatId(&seatMap)
	fmt.Println("My seat ID:", mySeatId)
}
