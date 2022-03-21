package models

import "fmt"


type Data struct{
	FirstName   string `json:"FirstName"`
	Lastname    string `json:"LastName"`
	Phonenumber string `json:"Phonenumber"`
	Age         string `json:"Age"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	DateOfBirth string `json:"DateOfBirth"`
}
 
func AddOne(object Data){

	fmt.Println(object.FirstName)
	fmt.Println(object.Lastname)
	fmt.Println(object.Phonenumber)
	fmt.Println(object.Age)
	fmt.Println(object.Email)
	fmt.Println(object.DateOfBirth)
	fmt.Println(object.Password)
}