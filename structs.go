package main

import "fmt"

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell // Anonymous field. Promotes the inner fields of PhoneHomeCell directly into the struct
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

// Replaces defining a method inside a class in c#. Use the struct as an argument prior to naming the func. Can create a method on any existing struct
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Pointer so the underlying value can be updated. Passing a value will not change the underlying value.
func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       42,
		address: Address{
			city:    "Denver",
			country: "USA",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "303-555-5555",
			cell: "303-444-4444",
		},
	}
	p1 := Person{
		firstName: "Jane",
	}
	p1.address.city = "New York"
	p1.address.country = "USA"

	p2 := Person{firstName: "Jane", address: Address{city: "New York", country: "USA"}}

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p.fullName())

	fmt.Println(p.home) // Directly from the anonymous field

	fmt.Println(p.address)
	fmt.Println(p1.address)

	fmt.Println("Are p and p1 equal?", p == p1)
	fmt.Println("Are p1 and p2 equal?", p1 == p2)

	// Anonymous structs
	user := struct {
		userName string
		email    string
	}{
		userName: "user",
		email:    "user@example.com",
	}
	fmt.Println(user)

	fmt.Println("Before Increment:", p.age)
	p.incrementAgeByOne()
	fmt.Println("After Increment:", p.age)

}
