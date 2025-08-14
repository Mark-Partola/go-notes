package main

import "fmt"

func main() {
	person := &Person{Name: "Alice"}
	fmt.Println(person.Name)
	ChangeName(person)
	fmt.Println(person.Name)
	ChangeName2(&person)
	fmt.Println(person.Name)
}

type Person struct {
	Name string
}

func ChangeName(p *Person) {
	*p = Person{Name: "John"}
}

func ChangeName2(p **Person) {
	*p = &Person{Name: "Bob"}
}
