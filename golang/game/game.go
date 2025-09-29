package main

import (
	"fmt"
	"slices"
)

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // must specify all fields
	fmt.Printf("i: %#v\n", i)

	// Can be in any order, can be partial
	i = Item{X: 11, Y: 22}
	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItemByPtr(10, 20))
	fmt.Println(NewItemByPtr(10, 2000 ))

	// Aside: Use %#v for debugging/logging
	/*a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n", a, b)
	fmt.Printf("a=%#v, b=%#v\n", a, b)
	*/

	i.MoveVal(10, 20)
	fmt.Printf("i (move val): %#v\n", i)

	i.MovePtr(10, 20)
	fmt.Printf("i (move ptr): %#v\n", i)

	p1 := Player{
		Name: "Parzival",
	}
	fmt.Printf("p1: %+v\n", p1)
	
	// Same way to access
	fmt.Println("p1.X:", p1.X)
	fmt.Println("p1.X:", p1.Item.X)

	p1.MovePtr(100, 200)
	fmt.Printf("p1 (move ptr): %+v\n", p1) 

	fmt.Println(p1.Found(Copper)) // <nil>
	fmt.Println(p1.Found(Copper)) // <nil>
	fmt.Println(p1.Found(Key(7))) // unknown key: "gold"
	fmt.Println("keys:", p1.Keys) // keys: [copper]

	ms := []Mover{
		&i,
		&p1,
	}

	moveAll(ms, 50, 70)
	for _, m := range ms {
		fmt.Println(m)
	}
}

/*
type Sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}
*/

/* Interfaces 
- Set of methods (and types)
- We define interfaces as "what you need", not "what you provide"
	- Interfaces are small (stdlib average ~2)
	- If you have interfaces with more than 4 methods, think again
*/
type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

// Implements Mover interface
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}

// String implements the fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

type Key byte
const (
	Copper Key = iota + 1
	Jade
	Crystal
)

func (p *Player) Found(key Key) error {
	switch key {
	case Copper, Jade, Crystal:
			// ok
	default:
		return fmt.Errorf("unknown key: %q", key)
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}

type Player struct {
	Name string
	Keys []Key
	Item // Player embeds Item
}

// Move moves i by delta x & delta y.
// "i" is called "the receiver"
func (i Item) MoveVal(dx, dy int) {
	i.X += dx
	i.Y += dy
}

// Move moves i by delta x & delta y.
// "i" is called "the receiver"
// i is a pointer receiver
/* Value vs pointer reveiver
- In general use value semantics
- Try to keep same semantics on all methods
- When you must use pointer reveiver
	- If you have lock fields
	- If you need to mutate the struct
	- Decoding/unmarshaling
*/
func (i *Item) MovePtr(dx, dy int) {
	i.X += dx
	i.Y += dy
}

/* Types of "new" or factory functions.
func NewItem(x, y int) Item
func NewItem(x, y int) *Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error)
*/

// Value semantics
func NewItemByVal(x, y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return Item{}, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	return i, nil
}

func NewItemByPtr(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	// The go compiler does escape analysis and will allocate i on the heap
	return &i, nil
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int
	Y int 
}