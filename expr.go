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

func calc_part(args []string) ([]string, os.Error) {
    var ret, i int
    newArgs := make([]string, 0, len(args) - 2)
    for i = 2; i < len(args); i++ {
        a, _ := strconv.Atoi(args[i-2])
        b, _ := strconv.Atoi(args[i-1])
        switch (args[i]) {
            case "+": ret = add(a, b); goto Rebuild; break
            case "-": ret = sub(a, b); goto Rebuild; break
            case "*": ret = mul(a, b); goto Rebuild; break
            case "/": ret = div(a, b); goto Rebuild; break
            case "%": ret = mod(a, b); goto Rebuild; break
            case "<": ret = lt (a, b); goto Rebuild; break
            case "<=":ret = lte(a, b); goto Rebuild; break
            case ">": ret = gt (a, b); goto Rebuild; break
            case ">=":ret = gte(a, b); goto Rebuild; break
            case "=": ret = eq (a, b); goto Rebuild; break
            case "!=":ret = neq(a, b); goto Rebuild; break
        }
    }

Rebuild:
    if len(args) != 3 {
        _ = copy(newArgs, args[0:i-2])
        newArgs = append(newArgs, strconv.Itoa(ret))
        if i+1 != len(args) { newArgs = append(newArgs, args[i+1:]...) }
        return calc_part(newArgs)
    }

    return []string{strconv.Itoa(ret)}, nil
}

func calculate(args []string) (int, os.Error) {
    str_res, err := calc_part(args)
    if err != nil { return 0, err }
    res, err := strconv.Atoi(str_res[0])
    if err == nil { return res, nil }
    return 0, err
}

func main() {
    var args []string

    flag.Parse();
    if len(flag.Args()) == 0 {
        fmt.Fprintln(os.Stderr, "Must specify arguments to expr")
        os.Exit(1)
    }

    args = flag.Args()
    ret, err := calculate(translate(args))
    if err != nil { fmt.Println(err.String()); os.Exit(1) }
    fmt.Printf("%d\n", ret)
}

// vim:ft=go
