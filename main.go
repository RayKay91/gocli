package main

import (
	JSON "encoding/json"
	f "fmt"
)

func main() {
	users := map[string]int{"ray": 31, "bobby": 28}
mainLoop:
	for {
		f.Println("Enter the name of the user:")
		name := ""
		_, err := f.Scanln(&name)
		if err != nil {
			panic(err)
		}
		if a, ok := users[name]; !ok {
			f.Println("user not found! Enter their age:")
			age := 0
			_, err := f.Scanln(&age)
			if err != nil {
				panic(err)
			}
			users[name] = age
			f.Println("User added!")
		} else {
			f.Printf("Yes! we found user %v with an age of %d\n", name, a)
		}
		f.Println("Would you like to check again? [y/n]")
	tryAgain:
		for {

			answer := ""
			_, e := f.Scanln(&answer)
			if e != nil {
				panic(e)
			} else {
				switch answer {
				case "y", "yes", "Y":
					f.Println("Let's go again!")
					continue mainLoop
				case "n":
					exit(&users)
					break mainLoop
				default:
					f.Println("you must enter yes or no!")
					continue tryAgain
				}
			}
		}
	}

}

func exit(users *map[string]int) {

	f.Println("finished!")
	u, err := JSON.Marshal(users)
	if err != nil {
		panic(err)
	}
	f.Println(string(u))
}
