package main

type store[P product] interface {
	sell(P)
}

type product interface {
	price() float64
	name() string
}

type book struct {
	title  string
	author string
	price  float64
}

type toy struct {
	price float64
	name  string
}

func price[P product](product P) float64 {
	return product.price()
}

func name[P product](product P) string {
	return product.name()
}
