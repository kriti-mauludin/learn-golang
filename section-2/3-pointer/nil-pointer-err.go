package main

import "fmt"

type User struct {
	first string
	last  string
}

func searchUser(id int) (*User, error) {
	if id != 1 {
		return nil, fmt.Errorf("USER NOT FOUND")
	}
	return &User{first: "Jane", last: "Doe"}, nil
}
func main() {
	n, err := searchUser(2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Full name:", n.first, n.last)
}
