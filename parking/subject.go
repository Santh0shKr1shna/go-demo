package parking

type Subject interface {
	register(obs Observer)
	notify()
}