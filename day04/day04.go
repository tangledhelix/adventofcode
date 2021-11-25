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
*
* --- Part Two ---
* The line is moving more quickly now, but you overhear airport security talking
* about how passports with invalid data are getting through. Better add some
* data validation, quick!
*
* You can continue to ignore the cid field, but each other field has strict
* rules about what values are valid for automatic validation:
*
* byr (Birth Year) - four digits; at least 1920 and at most 2002.
* iyr (Issue Year) - four digits; at least 2010 and at most 2020.
* eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
* hgt (Height) - a number followed by either cm or in:
* If cm, the number must be at least 150 and at most 193.
* If in, the number must be at least 59 and at most 76.
* hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
* ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
* pid (Passport ID) - a nine-digit number, including leading zeroes.
* cid (Country ID) - ignored, missing or not.
*
* Your job is to count the passports where all required fields are both present
* and valid according to the above rules. Here are some example values:
*
* byr valid:   2002
* byr invalid: 2003
*
* hgt valid:   60in
* hgt valid:   190cm
* hgt invalid: 190in
* hgt invalid: 190
*
* hcl valid:   #123abc
* hcl invalid: #123abz
* hcl invalid: 123abc
*
* ecl valid:   brn
* ecl invalid: wat
*
* pid valid:   000000001
* pid invalid: 0123456789
* Here are some invalid passports:
*
* eyr:1972 cid:100
* hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926
*
* iyr:2019
* hcl:#602927 eyr:1967 hgt:170cm
* ecl:grn pid:012533040 byr:1946
*
* hcl:dab227 iyr:2012
* ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277
*
* hgt:59cm ecl:zzz
* eyr:2038 hcl:74454a iyr:2023
* pid:3556412378 byr:2007
* Here are some valid passports:
*
* pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
* hcl:#623a2f
*
* eyr:2029 ecl:blu cid:129 byr:1989
* iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm
*
* hcl:#888785
* hgt:164cm byr:2001 iyr:2015 cid:88
* pid:545766238 ecl:hzl
* eyr:2022
*
* iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
*
* Count the number of valid passports - those that have all required fields and
* valid values. Continue to treat cid as optional. In your batch file, how many
* passports are valid?
 */

package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Data structures to represent and store passports. We don't need to store the
// data, only the presence of the field, so we can use bool.
type passport struct {
	byr string /* birth year */
	iyr string /* issue year */
	eyr string /* expiration year */
	hgt string /* height */
	hcl string /* hair color */
	ecl string /* eye color */
	pid string /* passport id */
	cid bool   /* country id - we ignore this, so don't store more than a bool */
}
type passportDatabase [500]passport

func check(e error) {
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
}

// Only check that required fields have a value; not what
// the value is. There's another function for that.
func checkRequiredFields(record passport) bool {
	if len(record.byr) == 0 {
		return false
	}
	if len(record.iyr) == 0 {
		return false
	}
	if len(record.eyr) == 0 {
		return false
	}
	if len(record.hgt) == 0 {
		return false
	}
	if len(record.hcl) == 0 {
		return false
	}
	if len(record.ecl) == 0 {
		return false
	}
	if len(record.pid) == 0 {
		return false
	}

	return true
}

func convertStringToNumber(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func validatePassport(record passport) bool {
	// Regex to check for 4-digit number
	reYear := regexp.MustCompile(`^\d{4}$`)

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	match := reYear.MatchString(record.byr)
	if match == false {
		return false
	}
	byr := convertStringToNumber(record.byr)
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	match = reYear.MatchString(record.iyr)
	if match == false {
		return false
	}
	iyr := convertStringToNumber(record.iyr)
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	match = reYear.MatchString(record.eyr)
	if match == false {
		return false
	}
	eyr := convertStringToNumber(record.eyr)
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	reHeight := regexp.MustCompile(`^(?P<Value>\d+)(?P<Unit>cm|in)$`)
	matches := reHeight.FindStringSubmatch(record.hgt)
	if len(matches) == 3 {
		measurement := convertStringToNumber(matches[1])
		if matches[2] == "cm" {
			if measurement < 150 || measurement > 193 {
				return false
			}
		} else if matches[2] == "in" {
			if measurement < 59 || measurement > 76 {
				return false
			}
		}
	} else {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	reHairColor := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	match = reHairColor.MatchString(record.hcl)
	if match == false {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	reEyeColor := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	match = reEyeColor.MatchString(record.ecl)
	if match == false {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	rePassportId := regexp.MustCompile(`^\d{9}$`)
	match = rePassportId.MatchString(record.pid)
	if match == false {
		return false
	}

	// cid (Country ID) - ignored, missing or not.

	// Didn't return false yet? Good news, that means it's valid!
	return true
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
				records[recordNum].byr = words[1]
			case "iyr":
				records[recordNum].iyr = words[1]
			case "eyr":
				records[recordNum].eyr = words[1]
			case "hgt":
				records[recordNum].hgt = words[1]
			case "hcl":
				records[recordNum].hcl = words[1]
			case "ecl":
				records[recordNum].ecl = words[1]
			case "pid":
				records[recordNum].pid = words[1]
			case "cid":
				records[recordNum].cid = true
			}
		}
	}

	// Now that the file has been parsed, let's evaluate each record for validity.
	// "valid" here means only that it has all of the required fields
	// "validated" means that the data is also good. Not the same thing!
	validPassports := 0
	validatedPassports := 0

	for recordNum := 0; recordNum < len(records); recordNum++ {
		if checkRequiredFields(records[recordNum]) {
			validPassports++
		}
		if validatePassport(records[recordNum]) {
			validatedPassports++
		}
	}

	fmt.Println("all fields present:", validPassports)
	fmt.Println("data validated:", validatedPassports)
}
