package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var b string
	var znak string
	var lishnie string
	fmt.Scanln(&a, &znak, &b, &lishnie)
	if lishnie != "" {
		panic("Формат математической операции не удовлетворяет калькулятор")
	}
	calculate(a, b, znak)
}
func calculate(a, b, znak string) {
	rimskie := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	arabskie := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}
	_, key := rimskie[a]
	_, key2 := rimskie[b]
	if key == true && key2 == true {
		switch znak {
		case "+":
			znachenie1 := rimskie[a]
			znachenie2 := rimskie[b]
			result := znachenie1 + znachenie2
			if result <= 10 {
				newkey := arabskie[result]
				fmt.Println(newkey)
			} else {
				panic("Invalid result")
			}
		case "-":
			znachenie1 := rimskie[a]
			znachenie2 := rimskie[b]
			result := znachenie1 - znachenie2
			if result > 0 {
				newkey := arabskie[result]
				fmt.Println(newkey)
			} else {
				panic("Invalid result")
			}
		case "/":
			znachenie1 := rimskie[a]
			znachenie2 := rimskie[b]
			result := znachenie1 / znachenie2
			if result > 0 {
				newkey := arabskie[result]
				fmt.Println(newkey)
			} else {
				panic("Invalid result")
			}
		case "*":
			znachenie1 := rimskie[a]
			znachenie2 := rimskie[b]
			result := znachenie1 * znachenie2
			if result > 0 && result <= 10 {
				newkey := arabskie[result]
				fmt.Println(newkey)
			} else {
				panic("Invalid result")
			}
		default:
			panic("Invalid znak")
		}
	} else if key == true && key2 == false {
		panic("Введите только римские либо только арабские числа")
	} else if key == false && key2 == true {
		panic("Введите только римские либо только арабские числа")
	} else if key == false && key2 == false {
		if b == "" {
			panic("Введите второй операнд")
		}
		number, err := strconv.Atoi(a)
		number2, err2 := strconv.Atoi(b)
		if err != nil || err2 != nil {
			panic("Вводить можно только числа")
		} else if number >= 1 && number <= 10 && number2 >= 1 && number2 <= 10 {
			switch znak {
			case "+":
				fmt.Println(number + number2)
			case "-":
				fmt.Println(number - number2)
			case "/":
				fmt.Println(number / number2)
			case "*":
				fmt.Println(number * number2)
			default:
				panic("Invalid znak")
			}
		} else {
			panic("Неверно введенные числа")
		}
	}
}
