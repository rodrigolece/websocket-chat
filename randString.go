package main

import (
    "math/rand"
    "time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func randString(n int) string {
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}
