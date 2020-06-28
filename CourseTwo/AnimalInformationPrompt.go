package CourseTwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
	name  string
	eat   string
	move  string
	speak string
}
type Cow struct {
	name  string
	eat   string
	move  string
	speak string
}
type Snake struct {
	name  string
	eat   string
	move  string
	speak string
}

//Bird
func (bird Bird) Eat() {
	fmt.Println(bird.eat)
}
func (bird Bird) Speak() {
	fmt.Println(bird.speak)
}
func (bird Bird) Move() {
	fmt.Println(bird.move)
}

//Snake
func (snake Snake) Eat() {
	fmt.Println(snake.eat)
}
func (snake Snake) Speak() {
	fmt.Println(snake.speak)
}
func (snake Snake) Move() {
	fmt.Println(snake.move)
}

//Cow
func (cow Cow) Eat() {
	fmt.Println(cow.eat)
}
func (cow Cow) Speak() {
	fmt.Println(cow.speak)
}
func (cow Cow) Move() {
	fmt.Println(cow.move)
}

func PrintRequestedProperty(property string, animal Animal) {
	if strings.TrimRight(property, "\n") == "eat" {
		animal.Eat()
	} else if strings.TrimRight(property, "\n") == "move" {
		animal.Move()
	} else if strings.TrimRight(property, "\n") == "speak" {
		animal.Speak()
	} else {
		fmt.Println("Wrong animal property")
	}
}

func FindAnimal(name string, animalMap map[string]Animal) (Animal, bool) {
	return animalMap[name], animalMap[name] != nil
}

func CreateAnimal(animalName string, animalType string) Animal {
	var animal Animal
	switch animalType {
	case "cow":
		cow := Cow{animalName, "grass", "walk", "moo"}
		animal = cow
	case "bird":
		bird := Bird{animalName, "worms", "fly", "peep"}
		animal = bird
	case "snake":
		snake := Cow{animalName, "mice", "slither", "hsss"}
		animal = snake
	}
	return animal
}

func main() {
	var line string
	animalMap := make(map[string]Animal)
	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		line, _ = reader.ReadString('\n')
		sliceOfLine := strings.Split(line, " ")
		if sliceOfLine[0] == "query" {
			animal, isRight := FindAnimal(sliceOfLine[1], animalMap)
			if isRight {
				PrintRequestedProperty(sliceOfLine[2], animal)
			} else {
				fmt.Println("Wrong animal name")
			}
		} else if sliceOfLine[0] == "newanimal" {
			name := sliceOfLine[1]
			animalType := sliceOfLine[2]
			CreateNewAnimal(name, animalType, &animalMap)
			fmt.Println("Created it!")
		}
	}
}

func CreateNewAnimal(name string, animalType string, m *map[string]Animal) {
	(*m)[name] = CreateAnimal(name, strings.TrimRight(animalType, "\n"))
	for i, v := range *m {
		fmt.Print(i + " ")
		fmt.Println(v)
	}
}
