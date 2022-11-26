package main

import (
	"fmt"
	"regexp"
	"time"
)

func switchMultiExpression() {
    switch time.Now().Weekday().String() {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("It's time to learn some Go.")
    default:
        fmt.Println("It's weekend, time to rest!")
    }

    fmt.Println(time.Now().Weekday().String())
}


func switchWithRegex() {
    var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
    var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

    contact := "foo@bar.com"

    switch {
    case email.MatchString(contact):
        fmt.Println(contact, "is an email")
    case phone.MatchString(contact):
        fmt.Println(contact, "is a phone number")
    default:
        fmt.Println(contact, "is not recognized")
    }
}

func switchFallthrough() {
    switch num := 15; {
    case num < 50:
        fmt.Printf("%d is less than 50\n", num)
        fallthrough
    case num > 100:
        fmt.Printf("%d is greater than 100\n", num)
        fallthrough
    case num < 200:
        fmt.Printf("%d is less than 200", num)
    }
}

