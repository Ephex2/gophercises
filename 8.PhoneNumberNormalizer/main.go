package main

import (
	"flag"
	"fmt"

	"github.com/ephex2/gophercises/8.PhoneNumberNormalizer/numberdb"
)

var phoneNumbers = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
}

var err error
var normalize bool
var setup bool

func init() {
	flag.BoolVar(&normalize, "normalize", false, "Specify this flag to normalize the numbers in the PhoneNumbers database.")
	flag.BoolVar(&setup, "setup", false, "Specify this flag to insert the values to be normalized into the PhoneNumber table in the PhoneNumberNormalizer database, if it exists (see the readme). NOTE: This will delete all rows currently in the PhoneNumber table.")
	flag.Parse()
}

func main() {
	if !normalize && !setup {
		// skip connecting to the database at all if no flag specified.
		fmt.Println("Please specify either the -setup or -normalize flag. Function cannot run without a mode selected.")
		return
	}

	err = numberdb.Store.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer numberdb.Store.Close()

	if setup {
		setupDB()
	}

	if normalize {
		normalizeDB()
	}
}

func setupDB() {
	fmt.Println("Starting setup...")
	fmt.Println("Deleting all rows from PhoneNumber table...") // consider truncate?
	err = numberdb.Store.RemoveAllNumbers()
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err.Error())
		return
	}

	fmt.Println("Inserting base phone numbers into the PhoneNumbers table (defined in main.go)...")
	err = numberdb.Store.NewNumbers(phoneNumbers)
	if err != nil {
		fmt.Printf("Error during setup: %v\n", err.Error())
		return
	}
	fmt.Println("Done.")
}

func normalizeDB() {
	fmt.Println("Normalizing PhoneNumber table...")
	err := numberdb.NormalizeDb()
	if err != nil {
		fmt.Printf("Error during normalization: %v\n", err.Error())
		return
	}
	fmt.Println("Done.")
}
