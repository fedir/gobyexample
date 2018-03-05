// _Структури_ в Go це типізовані колекції полів,
// корисні для групування данних разом задля формування
// записів.

package main

import "fmt"

// Ця структура `person` має поля для імені (`name`) та віку (`age`).
type person struct {
    name string
    age  int
}

func main() {

    // Наступний синтаксис створить нову структуру.
    fmt.Println(person{"Bob", 20})

    // Дозволяється іменувати поля при ініціалізації структури.
    fmt.Println(person{name: "Alice", age: 30})

    // Пропущені поля будуть створені відповідно нульового
    // значення типу пропущеного поля.
    fmt.Println(person{name: "Fred"})

    // Префікс `&` поверне вказівник на структуру.
    fmt.Println(&person{name: "Ann", age: 40})

    // Доступ до полів надається через синстксис крапки `.`
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // Дозволяється використовувати крапки з вказівниками структур,
    // вказівники, в такому разі, автоматично розіменовані.
    sp := &s
    fmt.Println(sp.age)

    // Данні в структурі можна змінювати ( тобто вони `mutable`,
    // такі що дозволяється змінбвати).
    sp.age = 51
    fmt.Println(sp.age)
}
