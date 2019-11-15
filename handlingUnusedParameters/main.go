package main

type Person struct {
	name string
}

func (p *Person) Name(capitalize bool) string {
	return p.name
}
