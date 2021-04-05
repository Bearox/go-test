package main

import (
    "fmt"
    "github.com/Bearox/go-test/greetings"
    "log"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    message, err := greetings.Hello("")

    if nil != err {
        log.Fatal(err)
    } else {
        fmt.Println(message)
    }
    fmt.Println("process finished.")

}
