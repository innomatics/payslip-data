package main

import (
    "testing"
)

func getIntAsserter(t *testing.T, testName string) (func(int, int, string)) {
    return func(actual int, expected int, name string)  {
        if actual != expected {
        t.Errorf("%s: %s was incorrect, got: %d, want: %d.", testName, name, actual, expected)
        }
    }
}