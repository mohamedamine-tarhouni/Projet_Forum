// package render

// import (
// 	"net/mail"
// 	"os/user"
// )

// type Address struct {
// 	Name    string // Proper name; may be empty.
// 	Address string // user@domain
// }

// func verif_nom(first_name string, last_name string) string {
// 	if len(first_name) < 2 {
// 		return("You must enter your firstname")

// 	} else if isAlpha(first_name) == false{
// 		return("Your firstname must be alphabethique")
// 	}
// 	if len(last_name) < 2 {
// 		return("You must enter your lastname")

// 	} else if isAlpha(last_name) == false{
// 		return("Your lastname must be alphabethique")
// 	}
// 	return "OK"
// }

// /*func verif_nom(name string) string {
//     if len(name)<4  {
//         return "You must enter your firstname"
//     }else if isAlpha(name)==false{
// return "your name must be alphabethique"
// }
// return ""*/
// func verif_mail(Address string) string{
// 	if Address != user@domain{
// 		println("Your mail must be in this form: user@domain")
// 	}
// 	return "OK"
// }
// func verif_pass(password string) bool {
// 	if len(password) < 8{
// 		return false
// 	}
// 	return true
// }
// /*verif(){
// if verif(last_name)!=""{
// return verif(last_name)
// }else if verif(first_name)!=""{
// return verif(first_name)
// }
// */
// func verif(first_name string, last_name string, Address string, password string) string{
// 	if verif_nom(first_name) == false {
// 		return verif_nom(first_name)
// 	}
// 	if verif_nom(last_name) == false {
// 		return verif_nom(last_name)
// 	}
// 	if verif_mail(Address) == false {
// 		return "You choose the right form"
// 	}
// 	if verif_pass(password) == false {
// 		return "Password invalid"
// 	}
// 	return "All the informations are correct"
// }
package render
