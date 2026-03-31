package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func upper(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToUpper(s)
}

func lower(s string) string {
    if len(s) == 0 {
        return s
    }
    return strings.ToLower(s)
}

func cap(s string) string {
    word := strings.Fields(s)

    for i, item := range word {
        word[i] = strings.ToUpper(string(item[0])) + strings.ToLower(string(item[1:]))
    }
    return strings.Join(word, " ")
}

func title(s string) string {

    words := strings.Fields(s)
    smallWords := map[string]bool{

        "a": true, "an": true, "the": true, "and": true, "but": true, "or": true,
        "for": true, "nor": true, "on": true, "at": true, "to": true, "by": true, "in": true,
        "of": true, "up": true, "as": true, "is": true, "it": true,
    }

    for i, item := range words {
        lower := strings.ToLower(item)

        if i == 0 || !smallWords[lower] {
            words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
        } else {
            words[i] = lower
        }
    }
    return strings.Join(words, " ")
}
func snake(s string) string {
    lower := strings.ToLower(s)
    words := strings.Fields(lower)

    for i, item := range words {

        updated := ""

        for _, char := range item {
            if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
                updated = string(char)
            }
        }
        words[i] = updated
    }
    return strings.Join(words, "_")
}

func reverse(s string) string {
    words := strings.Fields(s)

    word := ""
    char := ""

    for i := 0; i < len(words); i++ {
        char = words[i] + " "
        for j := len(char) - 1; j >= 0; j-- {
            word += string(char[j])
        }
    }
    return word
}

func main() {

start:
    fmt.Print("C&C> ")
    readInput := bufio.NewReader(os.Stdin)

    input, err := readInput.ReadString('\n')
    input = strings.TrimSpace(input)
    if err != nil {
        fmt.Println("Invalid input")
        return
    }
    if len(input) <= 1 {
        fmt.Println(" No text provided. Usage: upper <text>")
    }
    if len(input) == 0 {
        goto start
    }
    if strings.ToLower(input) == "quit" {
        fmt.Print("shutting down strings transformer\n GoodBye")
    }

    parts := strings.Fields(input)
    command := parts[0]
    textslice := parts[1:]
    words := strings.Join(textslice, " ")

    switch strings.ToLower(command) {
    case "upper":
        fmt.Println(upper(words))
    case "lower":
        fmt.Println(lower(words))
    case "cap":
        fmt.Println(cap(words))
    case "title":
        fmt.Println(title(words))
    case "snake":
        fmt.Println(snake(words))
    case "reverse":
        fmt.Println(reverse(words))

    default:
        fmt.Printf("Invalid command")
    }

}

