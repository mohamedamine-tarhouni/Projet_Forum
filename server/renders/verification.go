package render

import (
	"net/mail"
	"strings"
)

type Address struct {
	Name    string // Proper name; may be empty.
	Address string // user@domain
}

//this function verifies if a string has only alphabetic letters
func isAlpha(str string) bool {
	i := 0
	for range str {
		if (strings.ToUpper(string(str[i])) < "A") || (strings.ToUpper(string(str[i])) > "Z") {
			return false
		}
		i++
	}
	return true
}

//this function will be the verifier for the first and last names(alphabetic and contains its length is atleast 2)
func verif_nom(name string) string {
	if strings.Index(name, " ") != -1 {
		name = strings.ReplaceAll(name, " ", "")
	}
	if len(name) < 2 {
		return "0"

	} else if isAlpha(name) == false {
		return "2"
	}
	return "1"
}

//this function verifies if the given string is in the right form(user@domain) and it should be with a length of atleast 3
func verif_mail(Address string) string {
	if len(Address) < 3 {
		return "0"
	} else {
		_, err := mail.ParseAddress(Address)

		if err != nil {
			// log.Fatalf(err)
			return "2"
		}
	}

	return "1"
}

//this function verifies if the given string is atleast with a (length the variable) length
func verif_pass(password string, length int) string {
	if len(password) < length {
		return "0"
	}
	return "1"
}

//this function verifies the password confirmation
func verif_confirm(password, Cpassword string) string {
	if password != Cpassword {
		return "0"
	}
	return "1"
}

//this function takes a signup data(firstname,lastname,mail,password,confirm_password)
//and returns a character to each of these elements according to how the verification went
//("0" not good(mostly checks for lengths))
//("1" all good)
//("2" not a good form or data(alphabetic or user@Domain))
func Verif(first_name string, last_name string, Address string, password string, Cpassword string, User_name string) Errors {
	var Err Errors
	Err.Err_name = verif_nom(first_name)
	Err.Err_surname = verif_nom(last_name)
	Err.Err_User_name = verif_pass(User_name, 2)
	Err.Err_Email = verif_mail(Address)
	Err.Err_password = verif_pass(password, 8)
	Err.Err_Cpassword = verif_confirm(password, Cpassword)
	return Err
}
