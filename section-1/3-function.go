package main

func calculate(numberOne int, numTwo int) (sum int, diff int) {
	sum = numberOne + numTwo
	diff = numberOne - numTwo
	return sum, diff
}

func gabungName(nameOne, nameTwo string) string {
	return nameOne + " " + nameTwo
}

func main() {
	sum, diff := calculate(10, 5)
	println("Sum:", sum)
	println("Difference:", diff)

	namaGabungan := gabungName("K", "M")
	println("Nama Gabungan:", namaGabungan)
}
