package main

import "fmt"

type Dog struct {
    Name     string
    Breed    string
    Age      int
}

// Конструктор для создания новой собаки
func NewDog(name string, breed string, age int) *Dog {
    return &Dog{
        Name:  name,
        Breed: breed,
        Age:   age,
    }
}

// Метод GetAge получает возраст собаки
func (dog *Dog) GetAge() int {
    return dog.Age
}

// Метод SetAge устанавливает новый возраст собаке
func (dog *Dog) SetAge(age int) {
    dog.Age = age
}

// Метод String позволяет удобно выводить данные о собаке
func (dog *Dog) String() string {
    return fmt.Sprintf("%s (%s) is %d years old.", dog.Name, dog.Breed, dog.Age)
}

func main() {
    // Создаём новую собаку
    myDog := NewDog("Рекс", "Лабрадор", 3)

    // Получаем возраст собаки
    fmt.Println(myDog.GetAge()) // Выведет: 3

    // Устанавливаем новый возраст
    myDog.SetAge(4)

    // Показываем обновлённую информацию о собаке
    fmt.Println(myDog.String()) // Выведет: Рекс (Лабрадор) is 4 years old.
}