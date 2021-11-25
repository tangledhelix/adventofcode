/*
* --- Day 4: Passport Processing ---
* You arrive at the airport only to realize that you grabbed your North Pole
* Credentials instead of your passport. While these documents are extremely
* similar, North Pole Credentials aren't issued by a country and therefore
* aren't actually valid documentation for travel in most of the world.
*
* It seems like you're not the only one having problems, though; a very long
* line has formed for the automatic passport scanners, and the delay could upset
* your travel itinerary.
*
* Due to some questionable network security, you realize you might be able to
* solve both of these problems at the same time.
*
* The automatic passport scanners are slow because they're having trouble
* detecting which passports have all required fields. The expected fields are as
* follows:
*
* byr (Birth Year)
* iyr (Issue Year)
* eyr (Expiration Year)
* hgt (Height)
* hcl (Hair Color)
* ecl (Eye Color)
* pid (Passport ID)
* cid (Country ID)
*
* Passport data is validated in batch files (your puzzle input). Each passport
* is represented as a sequence of key:value pairs separated by spaces or
* newlines. Passports are separated by blank lines.
*
* Here is an example batch file containing four passports:
*
* ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
* byr:1937 iyr:2017 cid:147 hgt:183cm
*
* iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
* hcl:#cfa07d byr:1929
*
* hcl:#ae17e1 iyr:2013
* eyr:2024
* ecl:brn pid:760753108 byr:1931
* hgt:179cm
*
* hcl:#cfa07d eyr:2025 pid:166559648
* iyr:2011 ecl:brn hgt:59in
*
* The first passport is valid - all eight fields are present. The second
* passport is invalid - it is missing hgt (the Height field).
*
* The third passport is interesting; the only missing field is cid, so it looks
* like data from North Pole Credentials, not a passport at all! Surely, nobody
* would mind if you made the system temporarily ignore missing cid fields. Treat
* this "passport" as valid.
*
* The fourth passport is missing two fields, cid and byr. Missing cid is fine,
* but missing any other field is not, so this passport is invalid.
*
* According to the above rules, your improved system would report 2 valid
* passports.
*
* Count the number of valid passports - those that have all required fields.
* Treat cid as optional. In your batch file, how many passports are valid?
 */

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Data structures to represent and store passports. We don't need to store the
// data, only the presence of the field, so we can use bool.
type passport struct {
	byr bool /* birth year */
	iyr bool /* issue year */
	eyr bool /* expiration year */
	hgt bool /* height */
	hcl bool /* hair color */
	ecl bool /* eye color */
	pid bool /* passport id */
	cid bool /* country id */
}
type passportDatabase [500]passport

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

func main() {

	// A place to store the records we find in the data
	var records passportDatabase

	// Read input file and break into lines
	dat, err := os.ReadFile("input.txt")
	check(err)
	rawData := strings.Split(string(dat), "\n")

	// Keep track of the record number we're on. We have to do this manually
	// because the file is not one-record-per-line and we have to know when
	// to allocate a new passport entry to store our data to.
	recordNum := 0

	// Parse the file contents
	for lineNum := 0; lineNum < len(rawData); lineNum++ {
		line := string(rawData[lineNum])

		// If we started a new record after a blank line, skip to the next cycle
		// and increment the record number
		if line == "" {
			recordNum++
			continue
		}

		// Split the key:val pairs into an array, look at each in turn
		pairs := strings.Split(line, " ")
		for pairNum := 0; pairNum < len(pairs); pairNum++ {
			// Split the key:val into an array to isolate the key
			words := strings.Split(pairs[pairNum], ":")

			// Store true for any field that we found in this record.
			switch words[0] {
			case "byr":
				records[recordNum].byr = true
			case "iyr":
				records[recordNum].iyr = true
			case "eyr":
				records[recordNum].eyr = true
			case "hgt":
				records[recordNum].hgt = true
			case "hcl":
				records[recordNum].hcl = true
			case "ecl":
				records[recordNum].ecl = true
			case "pid":
				records[recordNum].pid = true
			case "cid":
				records[recordNum].cid = true
			}
		}
	}

	// Now that the file has been parsed, let's evaluate each record for validity.
	validPassports := 0
	for recordNum := 0; recordNum < len(records); recordNum++ {
		if records[recordNum].byr && records[recordNum].iyr && records[recordNum].eyr && records[recordNum].hgt && records[recordNum].hcl && records[recordNum].ecl && records[recordNum].pid {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}
