package main

import "fmt"

type Course struct {
	name string
	desc string
	time int
}

func main() {
	myHobbies := [3]string{"Football", "Cooking", "Travel"}
	fmt.Println(myHobbies)

	fmt.Println(myHobbies[0])
	fmt.Println(myHobbies[1:3])

	mainHobbies := myHobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(len(mainHobbies), cap(mainHobbies))
	// mainHobbies = mainHobbies[1:]
	fmt.Println(mainHobbies)
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	goals := []string{"Go", "Golang"}
	fmt.Println(goals)
	goals[1] = "Golang advanced"
	goals = append(goals, "Golang again")
	fmt.Println(goals)

	courses := []Course{{"Go", "Basics", 30}, {"Js", "Advanced", 67}, {"Go", "Advanced", 170}}
	fmt.Println(courses)

	newCourse := Course{
		"New Lang",
		"test",
		1,
	}

	courses = append(courses, newCourse)
	fmt.Println(courses)
}
