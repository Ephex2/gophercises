# Phone Number Normalizer - Gophercise 8

The purpose of this mini-project is to get us used to using the SQL libraries for Go.

The problem to solve is to normalize a specific set of numbers in a database, removing special characters and duplicates.


## Database setup
The database I used was a local postgres SQL server, logging in with the default user (postgres).
The connection string was stored in an environment variable called **PHONENUMBERNORMALIZER_URL**.
This is not a production setup -- nor is the database a production database. It contains one table named PhoneNumbers which has two columns, id and number. The script to create the:
- Database is [here](./CreateDatabase.postgres.sql).
- Table is [here](./CreateTable.postgres.sql).

I ran the scripts using the psql shell. A new connection to the PhoneNumberNormalizer database needs to be made after it is created, which is why the scripts are separated.

## Usage of the command-line tool
Two flags are defined when calling main.go:
- **setup**: Specify this flag to insert the values to be normalized into the PhoneNumber table in the PhoneNumberNormalizer database, if it exists. NOTE: This will delete all rows currently in the PhoneNumber table.
- **normalize**: Specify this flag to normalize the numbers in the PhoneNumbers database.

**NOTE** - if no flag is specified, the executable will simply not do anything.

<br></br>
The project was relatively simple so I didn't abstract the calls to the database more than putting them in a separate package.