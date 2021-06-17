package main

import "fmt"

type dog struct{}

func (d *dog) String() string {
	return "bow"
}

type frog struct{}

func (d *frog) String() string {
	return "ribbit"
}

type goat struct{}

func (d *goat) String() string {
	return "baa"
}

func PrintOut(s fmt.Stringer) {
	println(s.String())
}

func PrintOutInterface(i interface{}) {
	switch v := i.(type) {
	case dog:
		println("Dog: " + v.String())
	case frog:
		println("Frog: " + v.String())
	case goat:
		println("Goat: " + v.String())
	default:
		println("Unknown")
	}
}

func main() {
	d := dog{}
	f := frog{}
	g := goat{}

	PrintOut(&d)
	PrintOut(&f)
	PrintOut(&g)

	PrintOutInterface(d)
	PrintOutInterface(f)
	PrintOutInterface(g)
}
