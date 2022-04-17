package calculator

var logMessage = "[LOG]"

//Version of calculator
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

func Sum(num1, num2 int) int {
	return num1 + num2
}
