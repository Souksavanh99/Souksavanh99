package mapendstruct

import (
	"fmt"
)

func Learn() {

	//Maps

	m := map[string]int{"souk": 200, "pao": 100}
	fmt.Println(m)
	fmt.Println(m["souk"])
	m["souk"] = 1000
	fmt.Println(m["souk"])

	//loop

	for key, v := range m {
		fmt.Printf("%v => %v \n", key, v)
	}

	//delete map
	delete(m, "pao")
	for key, v := range m {
		fmt.Printf("%v => %v \n", key, v)
	}

	// add maps key

	m["golang"] = 400
	for key, v := range m {
		fmt.Printf("%v => %v \n", key, v)

	}

	// struct

	type User struct {
		id       int
		username string
		Password string
	}

	souk := User{
		id:       1,
		username: "souk pmc",
		Password: "123456",
	}
	fmt.Println(souk.username)
	souk.Password = "123456789"
	fmt.Println(souk.Password)

}
