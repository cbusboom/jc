# JumpCloud Interview Assignment

## Action Statistics Library

This library class provides action statistics.
Action times are recorded in the statistics via the **addAction** method.
Action averages can be retrieved via the **getStats** method.
Each of these methods are described below:

### Add Action

    addAction(string) returning error

This function accepts a JSON serialized string of the form below and maintains an average time
for each action. Below are three sample inputs:

    1) {"action":"jump", "time":100}
    2) {"action":"run", "time":75}
    3) {"action":"jump", "time":200}

### Get Statistics

    getStats() returning string

This function returns a JSON array of the average time for each action that has been
recorded by the **addAction** function. The output after the 3 sample inputs (above) would be:

    [
      {"action":"jump","avg":150},
      {"action":"run","avg":75}
    ]

### Example Usage

Below is an example of how to use the **Action Statistics Library**:

    package main

    import (
        "github.com/cbusboom/jc/pkg"
        "fmt"
    )

    func main() {
        s := actionstats.NewActionStats()

        s.AddAction(`{"action":"jump","time":100}`)
        s.AddAction(`{"action":"run","time":75}`)
        s.AddAction(`{"action":"jump","time":200}`)

        fmt.Println(s.GetStats())
    }

## How to download code and run it

**NOTE: These instructions assume you have a github account and have go installed on your computer**

On your computer create a directory to download the code to:

    mkdir ~/scratch

Download the code via git clone:

    cd ~/scratch
    git clone https://github.com/cbusboom/jc.git

Run the example program:

    cd ~/scratch/jc
    go run main.go

Run the test cases for actionstats:

    cd ~/scratch/jc/lib
    go test -v
