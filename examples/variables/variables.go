// В Go, _змінні_ декларуються явно і використовуються
// компілятором, наприклад для перевірки коректності типу
// у викликах функцій.

package main

import "fmt"

func main() {

    // Ключове слово `var` декларує 1 або більше змінних.
    var a = "initiaal"
    fmt.Println(a)

    // Дозволяється декларувати більш ніж одну змінну водночас.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go може припускати тип змінної відхштовхуючись від її значення.
    var d = true
    fmt.Println(d)

    // Змінні декларовані без супровідної ініціалізації
    // мають _нульове значення_. Наприклад, нульовим
    // значенням для цілих чисел (`int`) є 0 (нуль).
    var e int
    fmt.Println(e)

    // Синтаксис `:=` це скорочення для декларації та ініціалізації
    // змінної, наприклад для <nobr>`var f string = "short"`</nobr> в нашому випадку.
    f := "short"
    fmt.Println(f)
}
