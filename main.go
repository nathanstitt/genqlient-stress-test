package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//go:generate go run github.com/Khan/genqlient

func debugStruct(s ...interface{}) {
	fmt.Println(spew.Sdump(s))
}

func main() {

	account := Account_insert_input{
		Expiration_date: time.Now().Add(time.Hour * 24 * 15),
		Start_date:      time.Now(),
		Account_level:   "trial",
		Subdomain:       strings.ToLower("test"),
		Signup_code:     "1234",

		Owner: &Users_obj_rel_insert_input{
			Data: Users_insert_input{
				User_name:   strings.ToLower("test"),
				First_name:  "Test",
				Last_name:   "User",
				Password:    "password",
				Email:       "test@testcom",
				Role:        "owner",
				Last_visit:  time.Now(),
				Login_count: 1,
			},
		},
		Client: &Client_obj_rel_insert_input{
			Data: Client_insert_input{
				Name:           "Testing Inc",
				Notify_options: "none",
			},
		},
	}

	debugStruct(account)

}
