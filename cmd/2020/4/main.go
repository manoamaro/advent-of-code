package main

import (
	"regexp"
	"strconv"
	"strings"

	"manoamaro.github.com/advent-of-code/pkg/aoc"
)

var challenge = aoc.New(2020, 4, parseInput, part1, part2)

func main() {
	challenge.Run()
}

func parseInput(input string) []passport {
	passports := []passport{}

	tempPassport := ""
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			rawFields := strings.Fields(tempPassport)
			newPassport := passport{
				fields: make(map[string]string),
			}

			for _, rawField := range rawFields {
				fieldValue := strings.Split(rawField, ":")
				newPassport.fields[fieldValue[0]] = fieldValue[1]
			}

			passports = append(passports, newPassport)

			tempPassport = ""
		} else {
			tempPassport = tempPassport + " " + line
		}
	}
	return passports
}

func part1(passports []passport) int {
	validPassports := 0
	return validPassports
}

func part2(passports []passport) int {
	validPassports := 0
	for _, passport := range passports {
		if passport.isValid() {
			validPassports++
		}
	}
	return validPassports
}

type passport struct {
	fields map[string]string
}

var passportValidationFields = map[string]func(string) bool{
	"byr": func(s string) bool {
		v, err := strconv.Atoi(s)
		return err == nil && v >= 1920 && v <= 2002
	},
	"iyr": func(s string) bool {
		v, err := strconv.Atoi(s)
		return err == nil && v >= 2010 && v <= 2020
	},
	"eyr": func(s string) bool {
		v, err := strconv.Atoi(s)
		return err == nil && v >= 2020 && v <= 2030
	},
	"hgt": func(s string) bool {
		r, _ := regexp.Compile("^([0-9]+)(cm|in)$")
		submatch := r.FindStringSubmatch(s)
		if len(submatch) < 3 {
			return false
		}
		value, err := strconv.Atoi(submatch[1])
		if err != nil {
			return false
		}
		unit := submatch[2]

		return (unit == "cm" && value >= 150 && value <= 193) || (unit == "in" && value >= 59 && value <= 76)
	},
	"hcl": func(s string) bool {
		matched, err := regexp.MatchString("^#(?:[0-9a-f]{6})$", s)
		return err == nil && matched
	},
	"ecl": func(s string) bool {
		values := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, value := range values {
			if value == s {
				return true
			}
		}
		return false
	},
	"pid": func(s string) bool {
		matched, err := regexp.MatchString("^[0-9]{9}$", s)
		return err == nil && matched
	},
}

func (receiver passport) isValid() bool {
	for requiredField, f := range passportValidationFields {
		v, ok := receiver.fields[requiredField]
		if !ok || !f(v) {
			return false
		}
	}
	return true
}
