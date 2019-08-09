//------------------------------------------------------------------------------
// Test cases for the ActionStats library
//
// To run tests issue command: go test -v in directory containing ActionStats code
//------------------------------------------------------------------------------
package actionstats

import (
    "testing"
    "sync"
)

func TestInitialState(t *testing.T) {
    s := NewActionStats()
    actual := s.GetStats()
    expected := "[]"
    if actual != expected {
        t.Fatalf("GetStats() error, expected:'%s', actual:'%s'", expected, actual)
    }
}

func TestAddActionInvalidJson(t *testing.T) {
    s := NewActionStats()

    json := `{"action":invalid,"time":10}`
    err := s.AddAction(json)
    if (err == nil) {
        t.Fatalf("Expected error for AddAction('%s')", json)
    }

    json = `{"action":"run","time":invalid}`
    err = s.AddAction(json)
    if (err == nil) {
        t.Fatalf("Expected error for AddAction('%s')", json)
    }
}

func TestAddActionConcurrency(t *testing.T) {
    s := NewActionStats()
    actual := s.GetStats()
    expected := "[]"
    if actual != expected {
        t.Fatalf("GetStats() error, expected:'%s', actual:'%s'", expected, actual)
    }

    wg := sync.WaitGroup{}

    wg.Add(1)
    go func() {
        s.AddAction(`{"action":"add","time":10}`)
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        s.AddAction(`{"action":"add","time":20}`)
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        s.AddAction(`{"action":"add","time":30}`)
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        s.AddAction(`{"action":"sub","time":110}`)
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        s.AddAction(`{"action":"sub","time":120}`)
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        s.AddAction(`{"action":"sub","time":130}`)
        wg.Done()
    }()

    wg.Wait()

    expected1 := `[{"action":"add","avg":20},{"action":"sub","avg":120}]`
    expected2 := `[{"action":"sub","avg":120},{"action":"add","avg":20}]`
    actual = s.GetStats()
    if (actual != expected1) && (actual != expected2) {
        t.Fatalf("GetStats() error: expected:'%s' or '%s', actual:'%s'", expected1, expected2, actual)
    }
}
