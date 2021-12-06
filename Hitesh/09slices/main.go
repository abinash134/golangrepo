package main

import "fmt"

func main() {
	fmt.Println("welcome to slices")

	var fruitlist = []string{"apple", "banana"}
	fmt.Printf("type is %T\n", fruitlist)

	fruitlist = append(fruitlist, "cucumber", "docker")
	fmt.Println(fruitlist)

	highscores := make([]int, 4)

	highscores[0] = 123
	highscores[1] = 1223
	highscores[2] = 12553
	highscores[3] = 1253

	highscores = append(highscores, 555, 66, 777)

	//removal of values from the slices

	var cources = []string{"react", "java", "flutter", "c"}

	fmt.Println(cources)

	var index int = 2

	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println(cources)

}
