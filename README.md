# go-swedishssn
Parsing of swedish social security numbers (*personnummer*).

## Usage
```go
package main

import (
    swedishssn "github.com/EClaesson/go-swedishssn"
)

// type SwedishSsn struct {
//     date   time.Time
//	   number string
// 	   sign   string
// }

func main() {
    // Creates a SwedishSsn from string (handles + sign for age >= 100)
    ssn, err := swedishssn.FromString("010203-1234")

    // String
    ssn.String() // => "20010203-1234"
    ssn.StringShort() // => "010203-1234"
    ssn.StringNoSign() // => "200102031234"
    ssn.StringShortNoSign() // => "0102031234"

    // Age now
    ssn.Age() // => 20

    // Age at specific time.Time
    ssn.AgeAt(time.Now()) // => 20

    // swedishssn.Male or swedishssn.Female
    ssn.Sex() // => swedishssn.Male

    // Luhn control digit check
    ssn.IsValid() // => true
}
```