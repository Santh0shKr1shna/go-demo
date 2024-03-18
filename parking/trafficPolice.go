// concrete observer
package parking

type TrafficPolice struct {
	msg string
}

func (tp *TrafficPolice) Update() {
	tp.msg = "Parking lot full"
}