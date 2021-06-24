package modules

type Person struct {
	name string
	age  int
}

func (p *Person) SeDetails(name string, age int) { // *Person 참조 표시 해야  주소 참조를 해서 변경이 가능핟.
	p.name = name
	p.age = age

}

func (p Person) Getname() string {
	return p.name
}
