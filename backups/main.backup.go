package backups

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name_1"`
	Age  int    `json:"age"`
}

func runs() {
	p1 := Person{
		Name: "P'owen Eiei",
		Age:  25,
	}

	fmt.Println("P1", p1.Name)

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println("JSON Output:", string(jsonData))

	jsonArray := []Person{}

	jsonArray = append(jsonArray, Person{Name: "test", Age: 20})
	jsonArray = append(jsonArray, Person{Name: "test2", Age: 20})

	// i = index
	// d = value
	for _, d := range jsonArray {
		fmt.Println("d", d)
	}
	fmt.Println("jsonArray", jsonArray)

	for i := 0; i < len(jsonArray); i++ {
		fmt.Println("json Array stand", jsonArray[i])
	}
}
