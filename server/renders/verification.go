package render

import (
	"net/mail"
	"strings"
)

type Address struct {
	Name    string // Proper name; may be empty.
	Address string // user@domain
}

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

func verif_pass(password string, length int) string {
	if len(password) < length {
		return "0"
	}
	return "1"
}
func verif_confirm(password, Cpassword string) string {
	if password != Cpassword {
		return "0"
	}
	return "1"
}
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
