package main

import "fmt"

func main() {
	users := make(map[string]User)

	firstOrder := Order{
		ID:       "asd879",
		Products: []int{2, 4, 6},
	}

	secondOrder := Order{
		ID:       "939393",
		Products: []int{5, 2, 4, 7},
	}

	juan := User{
		ID:     1,
		Name:   "Juan",
		Orders: []Order{firstOrder},
	}

	jhon := User{
		Name:   "Jhon",
		Orders: []Order{secondOrder},
	}

	users[juan.Name] = juan
	users[jhon.Name] = jhon

	for _, u := range users {
		fmt.Println("User: ", u.Name)
		fmt.Println("Orders: ", u.CountOrders())
	}
}

type Order struct {
	ID       string
	Products []int
}

type User struct {
	ID     int
	Name   string
	Orders []Order
}

func (u User) CountProducts() int {
	t := 0

	for _, order := range u.Orders {
		t += len(order.Products)
	}

	return t
}

func (u User) CountOrders() int {
	return len(u.Orders)
}
