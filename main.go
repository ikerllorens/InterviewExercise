package main

import "fmt"

// To execute Go code, please declare a func main() in a package "main"

func main() {
	crypt := []string{"SEND", "MORE", "MONEY"}
	solution := [][]rune{{'O','0'}, {'M', '1'}, {'Y', '2'}, {'E', '5'}, {'N', '6'}, {'D', '7'}, {'R', '8'}, {'S', '9'}}

	res := isCryptSolution(crypt, solution)
	fmt.Println(res)
}

func isCryptSolution(crypt []string, solution [][]rune) bool {

	number1 := toNumber(crypt[0], solution)
	number2 := toNumber(crypt[1], solution)
	number3 := toNumber(crypt[2], solution)

	if(number1 + number2) == number3 {
		return true
	}
	return false
}


func toNumber(word string, solution [][]rune) int {
	number := 0
	for i, c := range word {
		position := (len(word)-1) -i
		for _, pair := range solution {
			if (pair[0] == c) {
				valueNumber := int(pair[1]) - 48
				for i:=0 ; i<position; i++ {
					valueNumber = valueNumber * 10
				}


				number = number + valueNumber
			}
		}
	}


	fmt.Println(number)

	return number
}

/*
A cryptarithm is a mathematical puzzle for which the goal is to find the correspondence between letters and digits, such that the given arithmetic equation consisting of letters holds true when the letters are converted to digits.

You have an array of strings crypt, the cryptarithm, and an an array containing the mapping of letters and digits, solution. The array crypt will contain three non-empty strings that follow the structure: [word1, word2, word3], which should be interpreted as the word1 + word2 = word3 cryptarithm.

If crypt, when it is decoded by replacing all of the letters in the cryptarithm with digits using the mapping in solution, the answer is true. If it does not become a valid arithmetic solution, the answer is false.

Example

For crypt = ["SEND", "MORE", "MONEY"] and

solution = [['O', '0'],
            ['M', '1'],
            ['Y', '2'],
            ['E', '5'],
            ['N', '6'],
            ['D', '7'],
            ['R', '8'],
            ['S', '9']]
the output should be
isCryptSolution(crypt, solution) = true.

When you decrypt "SEND", "MORE", and "MONEY" using the mapping given in crypt, you get 9567 + 1085 = 10652 which is correct and a valid arithmetic equation.
*/
