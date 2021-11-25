/*
 * --- Day 3: Toboggan Trajectory ---
 * With the toboggan login problems resolved, you set off toward the airport.
 * While travel by toboggan might be easy, it's certainly not safe: there's very
 * minimal steering and the area is covered in trees. You'll need to see which
 * angles will take you near the fewest trees.
 *
 * Due to the local geology, trees in this area only grow on exact integer
 * coordinates in a grid. You make a map (your puzzle input) of the open squares
 * (.) and trees (#) you can see. For example:
 *
 * ..##.......
 * #...#...#..
 * .#....#..#.
 * ..#.#...#.#
 * .#...##..#.
 * ..#.##.....
 * .#.#.#....#
 * .#........#
 * #.##...#...
 * #...##....#
 * .#..#...#.#
 *
 * These aren't the only trees, though; due to something you read about once
 * involving arboreal genetics and biome stability, the same pattern repeats to
 * the right many times:
 *
 * ..##.........##.........##.........##.........##.........##.......  --->
 * #...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
 * .#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
 * ..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
 * .#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
 * ..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
 * .#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
 * .#........#.#........#.#........#.#........#.#........#.#........#
 * #.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
 * #...##....##...##....##...##....##...##....##...##....##...##....#
 * .#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
 *
 * You start on the open square (.) in the top-left corner and need to reach the
 * bottom (below the bottom-most row on your map).
 *
 * The toboggan can only follow a few specific slopes (you opted for a cheaper
 * model that prefers rational numbers); start by counting all the trees you
 * would encounter for the slope right 3, down 1:
 *
 * From your starting position at the top-left, check the position that is right
 * 3 and down 1. Then, check the position that is right 3 and down 1 from there,
 * and so on until you go past the bottom of the map.
 *
 * The locations you'd check in the above example are marked here with O where
 * there was an open square and X where there was a tree:
 *
 * ..##.........##.........##.........##.........##.........##.......  --->
 * #..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
 * .#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
 * ..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
 * .#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
 * ..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
 * .#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
 * .#........#.#........X.#........#.#........#.#........#.#........#
 * #.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
 * #...##....##...##....##...#X....##...##....##...##....##...##....#
 * .#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
 *
 * In this example, traversing the map using this slope would cause you to
 * encounter 7 trees.
 *
 * Starting at the top-left corner of your map and following a slope of right 3
 * and down 1, how many trees would you encounter?
 *
 * --- Part Two ---
 * Time to check the rest of the slopes - you need to minimize the probability
 * of a sudden arboreal stop, after all.
 *
 * Determine the number of trees you would encounter if, for each of the
 * following slopes, you start at the top-left corner and traverse the map all
 * the way to the bottom:
 *
 * Right 1, down 1.
 * Right 3, down 1. (This is the slope you already checked.)
 * Right 5, down 1.
 * Right 7, down 1.
 * Right 1, down 2.
 *
 * In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s)
 * respectively; multiplied together, these produce the answer 336.
 *
 * What do you get if you multiply together the number of trees encountered on
 * each of the listed slopes?
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// This is a given because we know the input size
	const rowsOfTrees = 323
	const colsOfTrees = 31

	var pathsToCheck = [5][2]int{
		{1, 1},
		{3, 1}, /* checked in round 1 already */
		{5, 1},
		{7, 1},
		{1, 2}}

	// Create a tree map using a matrix of bools.
	// true: a tree, false: an empty square
	var treeMap [rowsOfTrees][colsOfTrees]bool

	// Read input data
	dat, _ := os.ReadFile("input.txt")
	raw_data := strings.Split(string(dat), "\n")

	// Loop over input, parse file
	for row := 0; row < len(raw_data); row++ {
		for col, c := range raw_data[row] {
			// Store this square's value
			if string(c) == "." {
				treeMap[row][col] = false
			} else if string(c) == "#" {
				treeMap[row][col] = true
			} else if string(c) != "\n" {
				// We can ignore newline, but if we see something else,
				// that's very unexpected.
				panic("Unexpected character" + string(c))
			}
		}
	}

	// Store our final answer
	answer := 1

	for path := 0; path < len(pathsToCheck); path++ {
		// How many trees we have encountered so far
		encounteredTrees := 0

		// Our position in the grid right now, starting from upper left.
		// This is standard grid coordinate notation, X is the COLUMN, Y is the ROW.
		posX := 0
		posY := 0

		addToX := pathsToCheck[path][0]
		addToY := pathsToCheck[path][1]

		stillInTheWoods := true
		// fmt.Println("Now checking", addToX, addToY)

		// Go use "for" instead of "while"... it's weird, but let's go with it
		for stillInTheWoods {
			// For each move, we need to move X +3 and Y +1. Apparently our toboggan is
			// a chess knight. Note that starting like this means we assume there is no
			// tree at (0,0) - perhaps we should check that, but we aren't here.

			posX += addToX
			posY += addToY
			// fmt.Printf("Now at (%d,%d)\n", posX, posY)

			// The first thing to note is that posX, posY, which are meant to track the
			// matrix treeMap, are zero-indexed. So we should always add 1 to their
			// value when doing comparisions to the rows, cols numbers, so we are
			// comparing apples to apples. But that's only when we do math to see if
			// we've exceeded the boundary of the map - never do that if looking into
			// the matrix data itself; the posX and posY are already using the proper
			// values to access matrix data.
			//
			// We need to look at the column we are in, and find out if we've
			// wrapped past the right edge and must back to the left (because it
			// repeats infinitely to the right).
			//
			// Example 1:
			// x=29 (adding 1 gives 30). Adding 3 gives 33. That's 2 more than colsOfTrees.
			// Result: must wrap back (we want to get to col 2, or index 1)
			// At this point the actual value of X should be 29+3 or 32.
			// (We got the 33 figure by adding 1 during comparision)
			// So to get to the desired index, deduct colsOfTrees from X: 32 - 31 = 1
			//
			// Example 2:
			// x=20 (adding 1 gives 21). Adding 3 gives 24, which is less than colsOfTrees.
			// Result: no change needed
			//
			// Example 3:
			// x=27 (adding 1 gives 28). Adding 3 gives 31, which is equal to colsOfTrees.
			// We've hit the right edge but not gone past it yet.
			// Result: no change needed
			//
			// Expressed as code, we can therefore say:
			//
			// x += 3
			// if x+1 > colsOfTrees {
			// 		/* we have to do something now */
			// }
			//
			// Next question is what to do about it? Given the above examples, the thing
			// to do is subtract colsOfTrees from X to get the new index.

			if posX+1 > colsOfTrees {
				posX -= colsOfTrees
				// fmt.Printf("Reset X to %d, now (%d,%d)\n", posX, posX, posY)
			}

			// Check is whether we are now past the bottom of the map, because then
			// we are finished. If Y+1 > rowsOfTrees, then we're out of the woods
			// already and we're done.

			if posY+1 > rowsOfTrees {
				// We are no longer in the woods!
				// break
				stillInTheWoods = false
			} else {
				// Is there a tree at this position?
				// fmt.Printf("RAW - %t\n", treeMap[posY][posX])
				if treeMap[posY][posX] {
					encounteredTrees++
					// fmt.Printf("Encountered a tree at (%d,%d)\n", posX, posY)
				}
			}
		}
		fmt.Printf("(%d,%d) encountered %d trees.\n", addToX, addToY, encounteredTrees)
		answer *= encounteredTrees
	}

	fmt.Println("Final answer:", answer)

}
