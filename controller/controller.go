package controller

import (
	"log"

	en "github.com/wanton-idol/cli-address-book/entity"
	uc "github.com/wanton-idol/cli-address-book/usecase"
)

func InsertContact(contact en.Contact) error {
	err := uc.InsertContactUC(contact)
	if err != nil {
		log.Printf("[InsertContact] Error inserting contact details: %v", err)
		return err
	}

	return nil
}

func SearchContactBasedOnName(name string) ([]en.Contact, error) {
	contacts, err := uc.SearchContactBasedOnNameUC(name)
	if err != nil {
		log.Printf("[SearchContactBasedOnName] Error searching in address book based on name: %v", err)
		return contacts, err
	}

	return contacts, nil
}

func SearchContactBasedOnPhone(phoneNumber int) (en.Contact, error) {
	contact, err := uc.SearchContactBasedOnPhoneUC(phoneNumber)
	if err != nil {
		log.Printf("[SearchContactBasedOnPhone] Error searching in address book based on phone number: %v", err)
		return contact, err
	}

	return contact, nil
}
