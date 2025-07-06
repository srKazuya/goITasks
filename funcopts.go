// package main

// import (
// 	"fmt"
// 	"time"
// )

// type User struct {
// 	name    string
// 	surname string
// 	bdate   time.Time
// }

// func NewUser (opts ... option) *User{
// 	i := &User{
// 		name: "user",
// 		surname: "user_surname",
// 		bdate: time.Now(),

// 	}
// 	for _, opt :=range opts{
// 		opt(i)
// 	}
// 	return i
// }

// type option func(*User)

// func SetName(name string) option {
// 	return func(i *User){
// 		i.name = name
// 	}
// }
// func SetSurname(surname string) option {
// 	return func(i *User){
// 		i.surname = surname
// 	}
// }

// func SetBdate(bdate time.Time) option {
// 	return func(i *User){
// 		i.bdate = bdate
// 	}
// }

// func printUser(u *User) {
// 	fmt.Printf("Name: %s, Surname: %s, BDate: %s\n", u.name, u.surname, u.bdate.Format("2006-01-02"))
// }

// func main()  {
// 	user1 := NewUser()
// 	user2 := NewUser(SetName("Anatoliy")) 
// 	user3 := NewUser(SetName("John"), SetSurname("Doe")) 
// 	user4 := NewUser(SetName("Steve"), SetSurname("Doe")) 
// 	user5 := NewUser(SetName("Steve"), SetSurname("Doe"), SetBdate(time.Date(1995, time.December, 10, 0, 0, 0, 0, time.UTC))) 
	
// 	fmt.Println("User1:"); printUser(user1)
// 	fmt.Println("User2:"); printUser(user2)
// 	fmt.Println("User3:"); printUser(user3)
// 	fmt.Println("User4:"); printUser(user4)
// 	fmt.Println("User5:"); printUser(user5)
// }