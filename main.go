package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1
	hobbies := [3]string{"Sport", "Cooking", "Reading"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5
	courseGoals := []string{"Learn Go", "Learn All The Basics"}
	fmt.Println(courseGoals)

	//6
	courseGoals[1] = "Learn All The Details!"
	courseGoals = append(courseGoals, "Learn All The Basics!")
	fmt.Println(courseGoals)

	//7
	product := []Product{
		Product{
			"First-product",
			"A First Product",
			12.99,
		},
		Product{
			"Second-product",
			"A Second Product",
			12.99,
		},
	}

	fmt.Println(product)

	newProduct := Product{
		"Third-product",
		"A Third Product",
		12.99,
	}

	product = append(product, newProduct)
	fmt.Println(product)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
