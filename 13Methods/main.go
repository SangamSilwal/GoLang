// Methods in Golang
package main

import "fmt"

func main() {
	sangam := User{"Sangam Silwal", "SangamSilwal@gmail.com", true, 57}
	sangam.GetStatus()
	sangam.NewMail()
	sangam.ChangeMail("Hello@gmail.com")
	fmt.Println(sangam.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int  as oneAge the o is not capital hence it is not exportable
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "SangamIsNotHey@gamil.com"
	fmt.Println("Email of the user is : ", u.Email)
}

//this is how we pass reference of the User to the function
func (u *User) ChangeMail(NewMail string) {
	u.Email = NewMail
}
