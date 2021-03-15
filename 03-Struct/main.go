package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactinfo //no need for specifier
}

type contactinfo struct {
	email string
	zipCode  int
}

func (p person)print()  {
	fmt.Println("Name: ",p.firstName)
	fmt.Println("SurName: ",p.lastName)
}

func (p *person) updateName(newFirstName string){
	(*p).firstName=newFirstName
}


func main() {
	serdar:=person{"Serdar", "Karaman",contactinfo{"serdar@karaman.dev",123123}}
	fmt.Println(serdar)
	serdar.print()
	fmt.Printf("%+v\n",serdar)


	var alex person
	fmt.Printf("%+v\n",alex)
	alex.firstName="Alex"
	alex.lastName="Anderson"
	fmt.Printf("%+v\n",alex)

	alex.updateName("WOLOLO")
	fmt.Printf("%+v\n",alex)
}
