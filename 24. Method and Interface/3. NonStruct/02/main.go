package main

import "fmt"

type book struct {
	title string
	price money
}

type game struct {
	title string
	price money
}

type money float64

type list []*game

func main() {
	var (
		// mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	var items []*game
	items = append(items, &minecraft, &tetris)

	// you can attach methods to a compatible type on the fly:
	// items -> []*game
	// list  -> []*game
	my := list(items)
	// my = nil

	// you can call methods even on a nil value
	my.print()
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

func (m money) string() string {
	// $xx.yy
	return fmt.Sprintf("$%.2f", m)
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}

func (l list) print() {
	// `list` acts like a `[]game`
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}
