package main

import "fmt"

func main() {
	statusCode := map[string]string{
		"200": "OK",
		"404": "Not Found",
	}
	fmt.Println("Status Code 200:", statusCode["200"])
	fmt.Println("Status Code 404:", statusCode["404"])

	student := []map[string]any{
		{
			"name": "KM",
			"age":  20,
		},
	}
	fmt.Println("Student 1:", student[0]["name"])
	fmt.Println("Student 1 Age:", student[0]["age"])

	//adding student
	student = append(student, map[string]any{
		"name": "KM2",
		"age":  21,
	})

	fmt.Println("Student 2:", student[1]["name"])

	delete(student[1], "age")
	fmt.Println("after delete age:", student[1])

	age, exist := student[1]["age"]
	if exist {
		fmt.Println("Student 2 Age:", age)
	} else {
		fmt.Println("Student 2 Age not found")
	}

	//slice of map
	users := []map[string]any{
		{
			"name": "KM",
			"age":  20,
		},
		{
			"name": "KM2",
			"age":  21,
		},
	}

	fmt.Println(users)
	for _, user := range users {
		fmt.Println("User Name:", user["name"])
		fmt.Println("User Age:", user["age"])
	}
}
