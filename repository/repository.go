package repository

import (
	"context"
	"log"

	en "github.com/wanton-idol/cli-address-book/entity"
	pkg "github.com/wanton-idol/cli-address-book/pkg"
)

var db = pkg.DBConnection()

func InsertContactRepo(contact en.Contact) error {
	_, err := db.Exec(insertContactDetails, contact.FirstName, contact.LastName, contact.Address, contact.PhoneNumber)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func SearchContactBasedOnNameRepo(name string) ([]en.Contact, error) {
	contacts := make([]en.Contact, 0)
	rows, err := db.QueryContext(context.Background(), selectRowsBasedOnName, name)
	if err != nil {
		return contacts, err
	}

	for rows.Next() {
		contact := en.Contact{}
		err := rows.Scan(&contact.FirstName, &contact.LastName, &contact.Address, &contact.PhoneNumber)
		if err != nil {
			return contacts, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func SearchContactBasedOnPhoneRepo(phoneNumber int) (en.Contact, error) {
	contacts := en.Contact{}
	err := db.QueryRow(selectRowBasedOnPhone, phoneNumber).Scan(&contacts.FirstName, &contacts.LastName, &contacts.Address, &contacts.PhoneNumber)
	if err != nil {
		return contacts, err
	}

	return contacts, nil
}
