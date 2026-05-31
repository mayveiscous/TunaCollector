package gc

import "fmt"

type Heap struct {
    objects  []*Object
    roots    []*Object
    nextID   int
    maxObjs  int
}

func NewHeap(maxObjs int) *Heap {
    return &Heap{
        objects: make([]*Object, 0, maxObjs),
        maxObjs: maxObjs,
    }
}

func (h *Heap) Alloc() *Object {
    if len(h.objects) >= h.maxObjs {
        h.Collect()
        if len(h.objects) >= h.maxObjs {
            panic("heap exhausted")
        }
    }
    obj := newObject(h.nextID)
    h.nextID++
    h.objects = append(h.objects, obj)
    return obj
}

func (h *Heap) AddRoot(obj *Object) {
    h.roots = append(h.roots, obj)
}

func (h *Heap) RemoveRoot(obj *Object) {
    for i, r := range h.roots {
        if r == obj {
            h.roots = append(h.roots[:i], h.roots[i+1:]...)
            return
        }
    }
}

func (h *Heap) Collect() {
    h.mark()
    h.sweep()
}

func (h *Heap) mark() {
    for _, root := range h.roots {
        h.markObject(root)
    }
}

func (h *Heap) markObject(obj *Object) {
    if obj == nil || obj.marked {
        return
    }
    obj.marked = true
    for _, ref := range obj.refs {
        h.markObject(ref)
    }
}

func (h *Heap) sweep() {
    live := h.objects[:0]
    for _, obj := range h.objects {
        if obj.marked {
            obj.marked = false
            live = append(live, obj)
        } else {
            fmt.Printf("[GC] swept object %d\n", obj.id)
        }
    }
    h.objects = live
}

func (h *Heap) Stats() {
    fmt.Printf("heap: %d objects live\n", len(h.objects))
    for _, obj := range h.objects {
        fmt.Printf("  obj#%d refs=%d\n", obj.id, len(obj.refs))
    }
}