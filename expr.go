package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
)

func boolToInt(a bool) int {
    if a { return 1 }
    return 0
}

func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }
func mod(a int, b int) int { return a % b }
func eq (a int, b int) int { return boolToInt(a == b) }
func neq(a int, b int) int { return boolToInt(a != b) }
func lt (a int, b int) int { return boolToInt(a <  b) }
func lte(a int, b int) int { return boolToInt(a <= b) }
func gt (a int, b int) int { return boolToInt(a >  b) }
func gte(a int, b int) int { return boolToInt(a >= b) }

func translate(args []string) []string {
    return args
}

func calc_part(args []string) []string, os.Error {
    for i := 0; i < len(args); i++ {
        switch (args[i]) {
            case "+":
            case "-":
            case "*":
            case "/":
            case "%":
            case "<":
            case "<=":
            case ">":
            case ">=":
            case "=":
            case "!=":
            default:
                _, err := strconv(args[i])
                if err { return nil, err }
        }
    }
}

func calculate(args []string) int {
    str_res, err := calc_part(args)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Parse error")
        return -1
    }
    res, err := strconv.Atoi(str_res)
    if err == nil { return res }

    fmt.Fprintln(os.Stderr, "Error converting string value to number")
}

func main() {
    var args []string

    flag.Parse();
    if len(flag.Args()) == 0 {
        fmt.Fprintln(os.Stderr, "Must specify arguments to expr")
        os.Exit(1)
    }

    args = flag.Args()
    calculate(translate(args))
}

// vim:ft=go
