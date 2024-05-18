package main

type person struct {
	firstName string
	lasttName string
}

func (p person) fullName() string {
	return p.firstName + p.lasttName
}
