package main

const (
	statusOK       = 200
	statusNotFound = 404

	statusInternalServerError uint16 = 500
)

func main() {
	println("Status OK:", statusOK)
	println("Status Not Found:", statusNotFound)
	println("Status Internal Server Error:", statusInternalServerError)
}
