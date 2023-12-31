package pack4

func Division(number1 int, number2 int) (int, int) {
	var Quotient int
	var Remainder int
	Quotient = number2 / number1
	Remainder = number2 % number1
	return Quotient, Remainder
}
