package main

import "fmt"

type Store[P product] interface {
	sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type Book struct {
	title  string
	author string
	price  float64
}

type Toy struct {
	price float64
	name  string
}

func (b Book) Price() float64 {
	return b.price
}
func (t Toy) Price() float64 {
	return t.price
}

func (b Book)Name() string  {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}
func (t Toy)Name() string  {
	return t.name
}

type ToyShop struct {
	soldToys []Toy
}

type BookShop struct {
	soldBooks []Book
}
func (toyShop *ToyShop)sell(t Toy) {
	toyShop.soldToys = append(toyShop.soldToys, t)
}
func (bookShop *BookShop)sell(b Book) {
	bookShop.soldBooks = append(bookShop.soldBooks, b)
}

func sellProducts[P product](store Store[P], items []P){
	for _, item := range items {
		store.sell(item)
	}
}

func main()  {
	bs := BookShop{
		soldBooks: []Book{},
	}

	// By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[Book](&bs, []Book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.soldBooks)

	ts := ToyShop{
		soldToys: []Toy{},
	}
	sellProducts[Toy](&ts, []Toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.soldToys)
}