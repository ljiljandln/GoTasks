package main

import "fmt"

type Singer struct {
	name string
}

func (singer *Singer) Singing() {
	fmt.Printf("Great singer %s is singing now!\n", singer.name)
}

type Actor struct {
	name string
}

func (actor *Actor) Acting() {
	fmt.Printf("Great actor %s is acting now!\n", actor.name)
}

// PersonAdapter целевой интерфейс
type PersonAdapter interface {
	Action()
}

// SingerAdapter адаптер для певца
type SingerAdapter struct {
	*Singer
}

// Action певца
func (adapter *SingerAdapter) Action() {
	adapter.Singing()
}

// NewSingerAdapter конструктор адаптера для певца
func NewSingerAdapter(singer *Singer) PersonAdapter {
	return &SingerAdapter{singer}
}

// ActorAdapter адаптер для актера
type ActorAdapter struct {
	*Actor
}

// Action актера
func (adapter *ActorAdapter) Action() {
	adapter.Acting()
}

// NewActorAdapter конструктор адаптера актера
func NewActorAdapter(actor *Actor) PersonAdapter {
	return &ActorAdapter{actor}
}

func main() {
	actor := Actor{name: "Sergey Bezrukov"}
	singer := Singer{name: "Nikolay Rastorguev"}

	aAdapter := NewActorAdapter(&actor)
	sAdapter := NewSingerAdapter(&singer)

	personAdapter1 := PersonAdapter(aAdapter)
	personAdapter1.Action()

	personAdapter2 := PersonAdapter(sAdapter)
	personAdapter2.Action()
}
