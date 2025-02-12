```go
// реалізація базової обробки даних в Go

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// створюємо зразок даних
	data := []Person{
		{"John", 28},
		{"Jane", 32},
		{"Mark", 30},
		{"Alice", 25},
		{"Bob", 27},
	}

	fmt.Println("Original Data:")
	printData(data)

	// сортуємо дані за віком
	sort.Slice(data, func(i, j int) bool {
		return data[i].Age < data[j].Age
	})

	fmt.Println("\nData sorted by age:")
	printData(data)

	// фільтруємо дані, відбираємо тільки тих, кому від 28 років
	filteredData := filterData(data, func(p Person) bool {
		return p.Age >= 28
	})

	fmt.Println("\nFiltered data (Age >= 28):")
	printData(filteredData)

	// перетворюємо дані, використовуючи функцію map
	mappedData := mapData(data, func(p Person) Person {
		p.Name = strings.ToUpper(p.Name)
		return p
	})

	fmt.Println("\nMapped data (Names in upper case):")
	printData(mappedData)
}

// функція для друку даних
func printData(data []Person) {
	for _, person := range data {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}

// функція для фільтрації даних
func filterData(data []Person, filterFunc func(Person) bool) []Person {
	var filteredData []Person
	for _, person := range data {
		if filterFunc(person) {
			filteredData = append(filteredData, person)
		}
	}
	return filteredData
}

// функція для перетворення даних
func mapData(data []Person, mapFunc func(Person) Person) []Person {
	var mappedData []Person
	for _, person := range data {
		mappedData = append(mappedData, mapFunc(person))
	}
	return mappedData
}
```

Цей код демонструє базові операції обробки даних, такі як сортування, фільтрація та перетворення. Всі ці операції реалізовані за допомогою вбудованих функцій Go та функцій вищого порядку.