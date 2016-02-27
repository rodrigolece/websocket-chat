package main

type wsEvent struct {
    Action string
    Data []data
}

type data struct {
    Type string
    Content interface{}
}
