//------------------------------------------------------------------------------
// Example program demonstrating usage of the actionstats library
//------------------------------------------------------------------------------
package main

import (
    "./lib"
    "fmt"
)

func main() {
    s := actionstats.NewActionStats()

    s.AddAction(`{"action":"jump","time":100}`)
    s.AddAction(`{"action":"run","time":75}`)
    s.AddAction(`{"action":"jump","time":200}`)

    fmt.Println(s.GetStats())
}
