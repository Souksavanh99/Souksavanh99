package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func Learn() {

	//if

	score := 10

	if score == 10 {

		fmt.Println(" very good!")
	} else {

		fmt.Println("Try again")
	}

	// switch
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MacOS")
	case "Windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s \n", os)
	}

	b, err := ioutil.ReadFile("file.txt")
	if err != nil {

		fmt.Println(err)

	}
	content := string(b)
	println(content)

	//for
	for i := 0; i < 10; i++ {
		println(i)
	}
	//break

	for j := 10; j >= 0; j-- {
		if j == 0 {
			fmt.Println("Boom!")
		}

		fmt.Println(j)
		time.Sleep(2 * time.Second)
	}

}
