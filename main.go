package main

import (
    "tunacollector/src"
)

func main() {
    heap := gc.NewHeap(20)
    a := heap.Alloc()
    b := heap.Alloc()
    heap.AddRoot(a)
    a.Ref(b)

    heap.Collect()
    heap.Stats()

    heap.RemoveRoot(a)
    heap.Collect()
    heap.Stats()
}