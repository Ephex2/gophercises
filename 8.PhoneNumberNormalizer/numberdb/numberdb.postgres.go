package numberdb

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dataStore struct {
	db *sqlx.DB
}

var Store dataStore

func (s *dataStore) Connect() error {
	db, err := sqlx.Connect("postgres", os.Getenv("PHONENUMBERNORMALIZER_URL"))
	s.db = db
	return err
}

func (s *dataStore) Close() {
	s.db.Close()
}

func (s *dataStore) NewNumber(number string) error {
	tx := Store.db.MustBegin()
	tx.MustExec("INSERT INTO public.\"PhoneNumbers\"\" (\"number\") VALUES ($1)", number)
	err := tx.Commit()
	return err
}

func (s *dataStore) NewNumbers(numbers []string) error {
	tx := Store.db.MustBegin()
	for _, number := range numbers {
		tx.MustExec("INSERT INTO public.\"PhoneNumbers\" (\"number\") VALUES ($1)", number)
	}
	err := tx.Commit()
	return err
}

func (s *dataStore) RemoveNumber(id int) error {
	tx := Store.db.MustBegin()
	tx.MustExec("DELETE FROM public.\"PhoneNumbers\" WHERE id=$1", id)
	err := tx.Commit()
	return err
}

func (s *dataStore) RemoveAllNumbers() error {
	tx := Store.db.MustBegin()
	tx.MustExec("DELETE FROM public.\"PhoneNumbers\"")
	err := tx.Commit()
	return err
}

func (s *dataStore) GetNumber(id int) (number phoneNumber, err error) {
	err = Store.db.Select(&number, "SELECT * FROM public.\"PhoneNumbers\" WHERE id=$1", id)
	return
}

func (s *dataStore) GetAllNumbers() (numbers []phoneNumber, err error) {
	err = Store.db.Select(&numbers, "SELECT * FROM public.\"PhoneNumbers\"")
	return
}

func NormalizeDb() error {
	numbers, err := Store.GetAllNumbers()
	if err != nil {
		return err
	}

	var normalizedNumbers []string
	var numberStrings []string
	for _, number := range numbers {
		numberStrings = append(numberStrings, number.Number)
	}

	for _, number := range numberStrings {
		number = normalizeFormat(number)
		normalizedNumbers = append(normalizedNumbers, number)
	}

	normalizedNumbers = getUnique(normalizedNumbers)

	err = Store.RemoveAllNumbers()
	if err != nil {
		return err
	}

	err = Store.NewNumbers(normalizedNumbers)
	return err
}

func normalizeFormat(old string) string {
	// Remove all non-integer characters from phone number string
	var newCharSlice string

	for _, char := range old {
		c := string(char)
		intChar, err := strconv.Atoi(c)

		if err == nil {
			newCharSlice = newCharSlice + fmt.Sprint(intChar)
		}
	}

	return string(newCharSlice)
}

func getUnique(numbers []string) (uniqueNumbers []string) {
	// get unique string from slice of strings, in this case the strings represent phone numbers.
	existenceMap := make(map[string]bool)

	for _, number := range numbers {
		_, ok := existenceMap[number]
		if !ok {
			uniqueNumbers = append(uniqueNumbers, number)
			existenceMap[number] = true
		}
	}

	return
}
