package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	ctrl "github.com/wanton-idol/cli-address-book/controller"
)

// searchcontactCmd represents the searchcontact command
var searchcontactCmd = &cobra.Command{
	Use:   "searchcontact",
	Short: "Search for any contact in address book based on first name, last name and phone number",
	Long: `
	Search for any contact in address book based on first name, last name and phone number.
	
	Use -n or --name flag to search based on first name and last name.
	use -p or --phone flag to search based on phone number.
	For Example: For search based on first name and last name
	searchcontact -n Name OR searchcontact --name Name

	For search based on Phone number
	searchcontact -p phoneNumber OR searchcontact --phone PhoneNumber

`,
	Run: func(cmd *cobra.Command, args []string) {
		sName, _ := cmd.Flags().GetBool("name")
		sPhone, _ := cmd.Flags().GetBool("phone")

		if sName {
			SearchBasedOnName(cmd, args)
		} else if sPhone {
			SearchBasedOnPhone(cmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchcontactCmd)
	searchcontactCmd.Flags().BoolP("name", "n", false, "Search contact based on name")
	searchcontactCmd.Flags().BoolP("phone", "p", false, "Search contact based on phone number")
}

func SearchBasedOnName(cmd *cobra.Command, args []string) {
	contacts, err := ctrl.SearchContactBasedOnName(args[0])
	if err != nil {
		log.Printf("[SearchContactBasedOnName] Error searching in address book based on name: %v", err)
	}

	fmt.Println(contacts)
}

func SearchBasedOnPhone(cmd *cobra.Command, args []string) {
	phone, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	contact, err := ctrl.SearchContactBasedOnPhone(phone)
	if err != nil {
		log.Printf("[SearchContactBasedOnPhone] Error searching in address book based on phone number: %v", err)
	}

	fmt.Println(contact)
}
