package CourseTwo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat   string
	move  string
	speak string
}

func CreateAnimal(eat string, move string, speak string) Animal {
	animalObj := *(new(Animal))
	animalObj.eat = eat
	animalObj.move = move
	animalObj.speak = speak
	return animalObj
}

func (animalObj *Animal) Eat() string {
	return animalObj.eat
}

func (animalObj *Animal) Move() string {
	return animalObj.move
}

func (animalObj *Animal) Speak() string {
	return animalObj.speak
}

func main() {
	var line string
	animalMap := CreateAnimals()
	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		line, _ = reader.ReadString('\n')
		sliceOfLine := strings.Split(line, " ")
		animal, isRight := FindAnimal(sliceOfLine[0], animalMap)
		if isRight {
			PrintRequestedProperty(sliceOfLine[1], animal)
		} else {
			fmt.Println("Wrong animal name")
		}
	}
}

func PrintRequestedProperty(property string, animal Animal) {
	if strings.TrimRight(property, "\n") == "eat" {
		fmt.Println(animal.Eat())
	} else if strings.TrimRight(property, "\n") == "move" {
		fmt.Println(animal.Move())
	} else if strings.TrimRight(property, "\n") == "speak" {
		fmt.Println(animal.Speak())
	} else {
		fmt.Println("Wrong animal property")
	}
}

func FindAnimal(s string, animalMap map[string]Animal) (Animal, bool) {
	var animal Animal
	isRight := true
	switch s {
	case "cow":
		animal = animalMap["cow"]
	case "bird":
		animal = animalMap["bird"]
	case "snake":
		animal = animalMap["snake"]
	default:
		isRight = false
	}
	return animal, isRight
}

func CreateAnimals() map[string]Animal {
	sliceOfAnimals := make(map[string]Animal)
	cow := CreateAnimal("grass", "walk", "moo")
	bird := CreateAnimal("worms", "fly", "peep")
	snake := CreateAnimal("mice", "slither", "hsss")
	sliceOfAnimals["cow"] = cow
	sliceOfAnimals["snake"] = snake
	sliceOfAnimals["bird"] = bird
	return sliceOfAnimals
}
