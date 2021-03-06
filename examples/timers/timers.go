// Інколи нам потрібно виконати певний код на якийсь момент
// в майбутньому, або запускати його через певні інтервали
// часу - це стає можливими завдяки сосбливим можливостям Go
// відомим як: _таймер_ та _маятник_. Спершу ми познайомимось
// з таймерами, а в насутпному прикладі побачимо що таке
// [маятник](tickers).

package main

import "time"
import "fmt"

func main() {

    // Таймер представляє поодиноку подію що станеться в
    // майбуутньому. Ви вказуєте таймеру на скільки часу
    // ви хочете зачекати, і таймер забезпечує вам канал
    // по якому ви можете отримати повідомлення про те що
    // час очікування скінчився. Наш таймер, наприклад,
    // розрахований на 2 секунди.
    timer1 := time.NewTimer(2 * time.Second)

    // `<-timer1.C` блокується каналом `C` допоки не буде
    // надіслано повідомлення що час витік.
    <-timer1.C
    fmt.Println("Таймер 1 просрочено")

    // Якщо вам просто хочеться зачекати, ми можете використати
    // `time.Sleep`. Одна з причин костності таймер, те що він
    // може бути зупинений в будьякий момент.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Таймер 2 просрочено")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Тaймер 2 зупинено")
    }
}
