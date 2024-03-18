// concrete observer
package parking

import "fmt"

type Owner struct {
	name string
}

func (own *Owner) Update() {
	fmt.Printf("Hi %s, Parking lot full!!", own.name)
}