package main

import "fmt"

func main() {

	// Testing Normal Printing Mechanism
	fmt.Println("Hello World")
	fmt.Printf("%d\n", 10)
	fmt.Print(5.6, "\n")
	fmt.Print(2)
	fmt.Println(!true)
	fmt.Println(!false && true)
	fmt.Println(false || !false)

	// Testing Variables

	var a, b, c int = 2, 3, 4
	fmt.Println(a, b, c)

	x := 2
	fmt.Println(x)

	message := "Hi Guyz"
	fmt.Println(message)

	var msg string = "Hello Guyz"
	fmt.Println(msg)

	p := 5
	fmt.Print("Munna has ", p, " apples\n")

	// Testing loop

	for i := 0; i < 3; i++ {
		fmt.Println("Loop")
	}

	// Tetsting If..else

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}

	x = 3
	if x == 1 {
		fmt.Println("Saturday")
	} else if x == 2 {
		fmt.Println("Sunday")
	} else if x == 3 {
		fmt.Println("Monday")
	} else {
		fmt.Println("Other days")
	}

	// Testing Switch

	switch x {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	default:
		fmt.Println("Any other day")
	}

	fn := func(val interface{}) {
		fmt.Println(val)
	}
	fn("hi")

	// Testing Array
	var odd [5]int
	var even [5]int
	var j, k int = 0, 0

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even[j] = i
			j++
		} else {
			odd[k] = i
			k++
		}
	}

	for i := 0; i < 5; i++ {
		fn(odd[i])
	}

	// Testing Slices
	oddSlice := make([]int, 0)
	evenSlice := make([]int, 0)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			evenSlice = append(evenSlice, i)
		} else {
			oddSlice = append(oddSlice, i)
		}
	}
	fn(oddSlice)
	fn(evenSlice)
	fn(oddSlice[1:])

	// Testing Maps

	myMap := make(map[string]string)
	myMap["munna"] = "boby"
	myMap["rubayet"] = "tanzina"
	myMap["sarwar"] = "mita"
	myMap["romeo"] = "juliet"
	myMap["mojnu"] = "laili"

	fn(myMap)

	_, stat := myMap["munnssa"]
	fn(stat)

}
