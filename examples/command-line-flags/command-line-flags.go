// [_Прапорці командного рядку_](https://uk.wikipedia.org/wiki/%D0%86%D0%BD%D1%82%D0%B5%D1%80%D1%84%D0%B5%D0%B9%D1%81_%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4%D0%BD%D0%BE%D0%B3%D0%BE_%D1%80%D1%8F%D0%B4%D0%BA%D0%B0#.D0.A4.D0.BE.D1.80.D0.BC.D0.B0.D1.82_.D0.BA.D0.BE.D0.BC.D0.B0.D0.BD.D0.B4.D0.B8) це загально прийнятий
// спосіб вказувати налаштування для программ
// призначений для командного рядку. Наприклад,
// в `wc -l` прапорецем є `-l`

package main

// Пакет Go `flag` підтримує основні операції
// розбору прапорців командного рядку.
// Ми скористаємось цим пакетом, щоб створити
// власну программу для командного рядка.
import "flag"
import "fmt"

func main() {

    // Основні декларування прапорців доступні для рядків,
    // цілих чисел та логічного типу данних. В цьому прикладі,
    // ми декларуємо прапорець `word` з стандартним значенням
    // `"foo"` та короткою довідкою. Функція `flag.String`
    // повертає вказівник рядка (а не рядкову змінну), ми ще
    // побачимо як користуватись таким вказівником трошки нижче.
    wordPtr := flag.String("word", "foo", "a string")

    // Тут ми декларуємо прапорці `numb` та `fork`
    // використовуючи вже знаомий по прапорцю `word` підхід.
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // Також можливо декларувати налаштування таким чином
    // щоб вже існуюча змінна прийняла на себе значення
    // передане прапорцю. Зауважте, що потрібно передати
    // вказівник цієї змінної функції що декларує прапорець.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Як тільки усі прапорці декларовані, викличемо `flag.Parse()`
    // щоб виконати аналіз конмадного рядку.
    flag.Parse()

    // Ми просто виведемо на екран усі розпізнані налаштуванняч та
    // позиційні аргументи. Зауважте, що нам потрібно розіменувати
    // вказівник, наприклад `*wordPtr`, щоб добути справжнє значення
    // налаштування.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
