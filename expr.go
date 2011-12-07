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

func or (a int, b int) int { if a != 0 { return a }; return b }
func and(a int, b int) int { if a != 0 && b != 0 { return a }; return 0 }

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
            case "+": ret = a + b; goto Rebuild; break
            case "-": ret = a - b; goto Rebuild; break
            case "*": ret = a * b; goto Rebuild; break
            case "/": ret = a / b; goto Rebuild; break
            case "%": ret = a % b; goto Rebuild; break
            case "<": ret = boolToInt(a <  b); goto Rebuild; break
            case "<=":ret = boolToInt(a <= b); goto Rebuild; break
            case ">": ret = boolToInt(a >  b); goto Rebuild; break
            case ">=":ret = boolToInt(a >= b); goto Rebuild; break
            case "=": ret = boolToInt(a == b); goto Rebuild; break
            case "!=":ret = boolToInt(a != b); goto Rebuild; break
            case "|": ret = or (a, b); goto Rebuild; break
            case "&": ret = and(a, b); goto Rebuild; break
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
    flag.Parse();
    args := flag.Args()
    if len(args) == 0 {
        fmt.Fprintln(os.Stderr, "Must specify arguments to expr")
        os.Exit(1)
    }
    ret, err := calculate(translate(args))
    if err != nil { fmt.Println(err.String()); os.Exit(1) }
    fmt.Printf("%d\n", ret)
}

// vim:ft=go
