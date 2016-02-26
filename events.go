package main

type wsEvent struct {
    Action string
    Data interface{}
}

type data struct {
    Type string
    Content interface{}
}
