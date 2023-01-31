package repository

var (
	insertContactDetails = `INSERT INTO contact(first_name, last_name, address, phone_number)
		VALUES($1, $2, $3, $4)`

	selectRowsBasedOnName = `SELECT first_name, last_name, address, phone_number
		FROM contact
		WHERE first_name = $1 OR last_name = $1`

	selectRowBasedOnPhone = `SELECT first_name, last_name, address, phone_number
		FROM contact
		WHERE phone_number = $1`
)
