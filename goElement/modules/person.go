package modules

type Person struct {
	name string
	age  int
}

func (p Person) SeDetails(name string, age int) {
	p.name = name
	p.age = age
}
