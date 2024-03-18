// concrete observer
package parking

type Owner struct {
	msg string
}

func (own *Owner) Update() {
	own.msg = "Parking lot full"
}