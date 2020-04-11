package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}
type Email string

type Customer struct {
	Name     string
	Gender   string
	Email    string
	Address  Address
	Emails   []string
	hasEmail bool
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

func (c Customer) Validate() error {
	var err error
	if c.hasEmail {
		err = validation.ValidateStruct(&c,
			validation.Field(&c.Emails, validation.Each(is.Email)),
		)
	}

	err2 := validation.ValidateStruct(&c,
		validation.Field(&c.Address, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Email, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.Gender, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		validation.Field(&c.Name, validation.Required, validation.By(checkAbc), validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
	// TO DO return combine err and err2
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func checkAbc(value interface{}) error {
	s, _ := value.(string)
	if s != "abc" {
		return errors.New("Custom validate must be abc")
	}
	return nil
}

func main() {

	// a := Address{
	// 	Street: "123456",
	// 	City:   "Unknown",
	// 	State:  "AC",
	// 	Zip:    "12345",
	// }

	// err := a.Validate()
	// fmt.Println(err)

	c1 := Customer{
		Name:     "Qiang Xue",
		hasEmail: true,
		Emails: []string{
			"invalid",
			"valid@example.com",
			"invalid",
		},
	}
	err := c1.Validate()
	fmt.Println(err)
	// Convert err to json
	b, _ := json.Marshal(err)
	fmt.Println(string(b))

	fmt.Println("========= END 1")

	c := Customer{
		Name:  "Qiang Xue",
		Email: "q",
		Address: Address{
			Street: "123 Main Street",
			City:   "Unknown",
			State:  "Virginia",
			Zip:    "12345",
		},
	}
	err = c.Validate()
	fmt.Println(err)
	fmt.Println("========= END 2")

	// Validate array struct
	addresses := []Address{
		Address{State: "MD", Zip: "12345"},
		Address{Street: "123 Main St", City: "Vienna", State: "VA", Zip: "12345"},
		Address{City: "Unknown", State: "NC", Zip: "123"},
	}
	err = validation.Validate(addresses)
	fmt.Println(err)

	data := ""
	err = validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
		is.URL,                    // is a valid URL
	)
	fmt.Println(err)
}
