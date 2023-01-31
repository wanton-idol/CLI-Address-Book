package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	ctrl "github.com/wanton-idol/cli-address-book/controller"
	en "github.com/wanton-idol/cli-address-book/entity"
)

// addcontactCmd represents the addcontact command
var addcontactCmd = &cobra.Command{
	Use:   "addcontact",
	Short: "Add a contact by specifying the firstName, lastName, address(should be single word) and phoneNumber",
	Long: `
	Add a contact by specifying the firstName, lastName, address(should be single word) and phoneNumber
	Should be in this format: addcontact firstName lastName address(should be single word) phoneNumber 
	For example: addcontact Nishchal Gupta Jaipur 9649070075`,
	Run: func(cmd *cobra.Command, args []string) {
		AddContact(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addcontactCmd)
}

func AddContact(cmd *cobra.Command, args []string) {
	contact := en.Contact{}

	for i, arg := range args {
		switch i {
		case 0:
			contact.FirstName = arg
		case 1:
			contact.LastName = arg
		case 2:
			contact.Address = arg
		case 3:
			phone, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatal(err)
			}
			contact.PhoneNumber = phone
		}
	}

	err := ctrl.InsertContact(contact)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contact Added: %v", contact)
}
