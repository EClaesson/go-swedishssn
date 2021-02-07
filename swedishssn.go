package swedishssn

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	luhn "github.com/EClaesson/go-luhn"
	age "github.com/bearbin/go-age"
)

// SwedishSsn represents a swedish social security number (personnummer).
type SwedishSsn struct {
	date   time.Time
	number string
	sign   string
}

// Sex represents biological birth-sex according to SSN.
type Sex int

// Sex definitions
const (
	Male Sex = iota
	Female
)

// FromString parses a SSN as string and creates a SwedishSsn.
func FromString(ssn string) (SwedishSsn, error) {
	cleanRe := regexp.MustCompile(`[^\d\+\-]`)
	clean := cleanRe.ReplaceAllString(ssn, "")

	parseRe := regexp.MustCompile(
		`(\d{2,4})(\d{2})(\d{2})([\+\-]?)(\d{4})`)
	matches := parseRe.FindStringSubmatch(clean)

	yearStr := matches[1]
	yearStrLen := len(yearStr)

	if yearStrLen == 2 {
		yearStr = "19" + yearStr
	}

	year, err := strconv.Atoi(yearStr)
	month, err := strconv.Atoi(matches[2])
	day, err := strconv.Atoi(matches[3])
	sign := matches[4]
	number := matches[5]

	if sign == "" {
		sign = "-"
	}

	if time.Now().Year()-year >= 100 && yearStrLen == 2 && sign == "-" {
		year += 100
	}

	if err != nil {
		return SwedishSsn{}, fmt.Errorf("date parse error")
	}

	return SwedishSsn{
		date:   time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location()),
		number: number,
		sign:   sign,
	}, nil
}

// String converts the SSN to a string with 4-digit year with separator sign.
func (ssn SwedishSsn) String() string {
	year, month, day := ssn.date.Date()

	return fmt.Sprintf("%d%02d%02d%s%s", year, int(month), day, ssn.sign, ssn.number)
}

// StringShort converts the SSN to a string with 2-digit year with separator sign.
func (ssn SwedishSsn) StringShort() string {
	return ssn.String()[2:]
}

// StringNoSign converts the SSN to a string with 4-digit year without separator sign.
func (ssn SwedishSsn) StringNoSign() string {
	year, month, day := ssn.date.Date()

	return fmt.Sprintf("%d%02d%02d%s", year, int(month), day, ssn.number)
}

// StringShortNoSign converts the SSN to a string with 2-digit year without separator sign.
func (ssn SwedishSsn) StringShortNoSign() string {
	return ssn.StringNoSign()[2:]
}

// AgeAt calculates the persons age at a specific time.
func (ssn SwedishSsn) AgeAt(at time.Time) int {
	return age.AgeAt(ssn.date, at)
}

// Age calculates the persons age at the current time.
func (ssn SwedishSsn) Age() int {
	return ssn.AgeAt(time.Now())
}

// Sex calculates the persons birth-sex.
func (ssn SwedishSsn) Sex() Sex {
	sexDigit, _ := strconv.Atoi(string(ssn.number[2]))

	if sexDigit%2 == 0 {
		return Female
	}

	return Male
}

// IsValid checks that the ssn and control digit is valid.
func (ssn SwedishSsn) IsValid() bool {
	valid, _ := luhn.IsValid(ssn.StringShortNoSign())

	return valid
}
