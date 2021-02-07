package swedishssn

import (
	"testing"
	"time"
)

func makeTestSsn() SwedishSsn {
	str := "010203+1234"
	ssn, _ := FromString(str)

	return ssn
}

func TestFromString(t *testing.T) {
	var ssnStrs = []string{
		"20010203-1234",
		"010203-1234",
		"2001 02 03 -   12 3 4",
	}

	for _, str := range ssnStrs {
		ssn, err := FromString(str)

		if err != nil {
			t.Errorf("error in FromString: %s", err)
			continue
		}

		year, month, day := ssn.date.Date()

		if year != 2001 || month != 2 || day != 3 {
			t.Errorf("incorrect date: \"%s\" -> %d, %d, %d", str, year, month, day)
		}

		if ssn.number != "1234" {
			t.Errorf("incorrect number: \"%s\" -> %s", str, ssn.number)
		}

		if ssn.sign != "-" {
			t.Errorf("incorrect sign: \"%s\" -> %s", str, ssn.sign)
		}
	}
}

func TestFromStringPlus(t *testing.T) {
	ssn := makeTestSsn()

	year, month, day := ssn.date.Date()

	if year != 1901 || month != 2 || day != 3 {
		t.Errorf("incorrect date: \"%s\" -> %d, %d, %d", "19010203+1234", year, month, day)
	}
}

func TestString(t *testing.T) {
	ssn := makeTestSsn()

	if ssn.String() != "19010203+1234" {
		t.Errorf("incorrect String()")
	}

	if ssn.StringShort() != "010203+1234" {
		t.Errorf("incorrect StringShort()")
	}

	if ssn.StringNoSign() != "190102031234" {
		t.Errorf("incorrect StringNoSign()")
	}

	if ssn.StringShortNoSign() != "0102031234" {
		t.Errorf("incorrect StringShortNoSign()")
	}
}

func TestAge(t *testing.T) {
	ssn := makeTestSsn()

	age := ssn.AgeAt(time.Date(2001, 02, 03, 0, 0, 0, 0, time.Now().Location()))

	if age != 100 {
		t.Errorf("incorrect age")
	}
}

func TestSex(t *testing.T) {
	ssn := makeTestSsn()

	if ssn.Sex() != Male {
		t.Errorf("incorrect sex")
	}
}

func TestIsValid(t *testing.T) {
	ssn := makeTestSsn()

	if !ssn.IsValid() {
		t.Errorf("incorrect IsValid()")
	}
}
