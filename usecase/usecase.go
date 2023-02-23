package usecase

import (
	"log"

	en "github.com/wanton-idol/cli-address-book/entity"
	repo "github.com/wanton-idol/cli-address-book/repository"
)

func InsertContactUC(contact en.Contact) error {
	err := repo.InsertContactRepo(contact)
	if err != nil {
		log.Fatalf("[InsertContactUC] Error insterting contact details: %v", err)
		return err
	}

	return nil
}

func SearchContactBasedOnNameUC(name string) ([]en.Contact, error) {
	contact, err := repo.SearchContactBasedOnNameRepo(name)
	if err != nil {
		log.Fatalf("[SearchContactBasedOnNameUC] Error searching name in the address book: %v", err)
		return contact, err
	}

	return contact, nil
}

func SearchContactBasedOnPhoneUC(phoneNumber int) (en.Contact, error) {
	contact, err := repo.SearchContactBasedOnPhoneRepo(phoneNumber)
	if err != nil {
		log.Fatalf("[SearchContactBasedOnPhoneUC] Error searching phone number in the address book: %v", err)
		return contact, err
	}

	return contact, nil
}
