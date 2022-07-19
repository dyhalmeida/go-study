package main

import "fmt"

func main() {

	/*
	Obs: Switch em GO, não precisa da cláusula break, pois quando o case é avaliado
	e verdadeiro, o switch é encerrado.
	*/

	day := dayOfWeekExample1(2) // Monday
	fmt.Println(day)

	day = dayOfWeekExample2(4)
	fmt.Println(day) // Wednesday

	day = dayOfWeekExample3(6)
	fmt.Println(day) // Friday
}

func dayOfWeekExample1(value int) string {
	switch value {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3: 
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday" 
	case 7:
		return "Saturday"
	default:
		return "Invalid day number"
	}
}

func dayOfWeekExample2(value int) string {
	switch {
	case value == 1:
		return "Sunday"
	case value == 2:
		return "Monday"
	case value == 3: 
		return "Tuesday"
	case value == 4:
		return "Wednesday"
	case value == 5:
		return "Thursday"
	case value == 6:
		return "Friday" 
	case value == 7:
		return "Saturday"
	default:
		return "Invalid day number"
	}
}

func dayOfWeekExample3(value int) string {
	var day string

	switch {
	case value == 1:
		day = "Sunday"
	case value == 2:
		day = "Monday"
	case value == 3: 
		day = "Tuesday"
	case value == 4:
		day = "Wednesday"
	case value == 5:
		day = "Thursday"
	case value == 6:
		day = "Friday" 
	case value == 7:
		day = "Saturday"
	default:
		day = "Invalid day number"
	}

	return day
}