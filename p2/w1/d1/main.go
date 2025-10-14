package main

import (
	"fmt"
	"reflect"
	"regexp"
)

func main() {

	type User struct {
		Name   string `require:"true"`
		Alamat string `require:"true" minLen:"5"`
		Email  string `require:"true"`
	}

	var NewUser = User{
		Name:   "sad",
		Alamat: "jl. Lumbu Timur 5",
		Email:  "raflymail.com",
	}

	typeUser := reflect.TypeOf(NewUser)
	userField := typeUser.NumField()

	fmt.Println("type user: ", typeUser)
	fmt.Println(typeUser.Field(1).Tag.Get("minLen"))
	fmt.Println("there is ", userField, "field in User struct")

	err := ValidateStruct(NewUser)
	fmt.Println(err)
}

func ValidateStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		if field.Tag.Get("require") == "true" {
			if strVal, ok := value.(string); ok && strVal == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
		if field.Name == "Email" {
			if strVal, ok := value.(string); ok {
				match, _ := regexp.MatchString(`[a-z]+@[a-z]+\.[a-z]`, strVal)
				if !match {
					return fmt.Errorf("%s value is not email standard", field.Name)
				}
			}
		}
	}
	return nil
}
