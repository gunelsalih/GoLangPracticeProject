package structs

import "fmt"

type customer struct {
	firtsName string
	lastname  string
	age       int
}

func (c customer) update() {
	fmt.Println("Güncellendi", c.firtsName)
}
func Demo2() {

	c := customer{firtsName: "Salih", lastname: "Günel", age: 35}
	c.update()
}
