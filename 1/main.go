package main

import "fmt"

// Human struct
type Human struct {
	Sex  string
	Name string
}

func (h Human) GetPronoun() string {
	if h.Sex == "male" {
		return "he"
	}
	return "she"
}

func (h Human) Talk() {
	fmt.Printf("%s is talking, because %s is a human\n", h.Name, h.GetPronoun())
}

// Action struct
type Action struct {
	Human
	Action string
}

// New Action constructor
func NewAction(sex string, name string, age int8, action string) *Action {
	return &Action{
		Human:  Human{Sex: sex, Name: name},
		Action: action,
	}
}

func (a Action) Do() {
	fmt.Printf("%s is %s, because %s can\n", a.Name, a.Action, a.GetPronoun())
}

func main() {
	newAction := NewAction("female", "Katya", 20, "programming")

	newAction.Talk()
	newAction.Do()
}
