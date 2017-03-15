package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type Gamer struct {
	Person
	First        string
	Plays        int
	FavoriteGame string
}

type BoardGamer struct {
	Person
	Gamer
	First string
}

func main() {
	b := BoardGamer{}
	b.First = "BoardGamer Joe"
	b.Gamer.First = "Gamer Joe"
	b.Person.First = "Person Joe"
	b.Last = "smith"
	b.Age = 42
	b.Plays = 100
	b.FavoriteGame = "LotR"
	fmt.Println(b)
	fmt.Println(&b)
}
