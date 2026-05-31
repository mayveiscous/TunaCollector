package main

import (
    "fmt"
    "tunacollector/src"
)

func main() {
    heap := gc.NewHeap(20)

    fmt.Println("Program Starting")

    config := heap.Alloc()
    logger := heap.Alloc()
    db := heap.Alloc()

    config.Ref(logger)
    config.Ref(db)

    heap.AddRoot(config)

    heap.Stats()
    heap.Collect()
    fmt.Println("After Collect:")
    heap.Stats()

    fmt.Println("Function Call")

    localA := heap.Alloc()
    localB := heap.Alloc()

    localA.Ref(db)

    heap.AddRoot(localA)
    heap.AddRoot(localB)

    heap.Stats()
    heap.Collect()
    fmt.Println("After Collect:")
    heap.Stats()

    heap.RemoveRoot(config)

    heap.Collect()
    fmt.Println("After Collect:")
    heap.Stats()
}