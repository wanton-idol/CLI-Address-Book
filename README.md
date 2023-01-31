# 📖: Address Book CLI in Golang  
 
This Address Book allows adding and searching for contacts through the CLI.

## :computer: Tech used 

1. Golang
2. Cobra library
3. PostgreSQL (with indexing)

## Application Requirement

1. Golang https://golang.org/dl/
2. Cobra library for CLI `go install github.com/spf13/cobra-cli@latestx`
3. PostgreSQL library to connect with PostgreSQL `go get github.com/lib/pq`
4. github.com/joho/godotenv to access the environment variable.

## 🐼: Walk through the application

### Make Sure your PostgreSQL server is running and the values of connection string in `.env` file are correct.
1. Run the program `go run main.go`
2. After successfully running the program, you will see a CLI which mainly has 2 commands.
3. `addcontact` and `searchcontact`.
4. `addcontact` takes a []string as input and to run this use command `go run main.go addcontact firstName lastName address(Only one word) phoneNumber`
5. `searchcontact` has two flags one is `-n or --name` for searching in address book based on first or last name and other is `-p or --phone` for searching based on phone number.
6. To run `searchcontact` commands use `go run main.go searchcontact -n [firstName or lastname]` and `go run main.go searchcontact -p phonenumber`.

## 🐼: Background working of application

1. Created indexing on first_name, last_name and phone_number colummns in PostgreSQL which increases the speed of searching in the data. Also, For more advancement we can implement elasticsearch along with PostgreSQL for searching.
2. Built the application by following design patterns.   

## Author   

#### :wave: [Nishchal Gupta](https://www.linkedin.com/in/itsnishchal/)


## License

[MIT License](LICENSE).
