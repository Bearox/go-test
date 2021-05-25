package main

import (
    "fmt"
    "github.com/Bearox/go-test/chan-test"
    "github.com/Bearox/go-test/greetings"
    utilTest "github.com/Bearox/go-test/util-test"
    "log"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    chan_test.ChanTest()
}

func jsonTest() {
    values := []string{
        "weiwen",
        "Bearox",
        "Davin",
    }

    messages, err := utilTest.MarshalTest(values)

    if nil != err {
        log.Fatal(err)
    }
    fmt.Println(messages)
}

func greetingsTest()  {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    names := []string{
        "weiwen",
        "Bearox",
        "Davin",
    }

    messages, err := greetings.Hellos(names)

    if nil != err {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
