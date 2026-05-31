package gc

// Object is a single heap-allocated node.
// In a real language runtime the payload would hold
// actual values — ints, strings, etc. Here it's just
// an ID so we can track it in tests.
type Object struct {
    id      int
    marked  bool
    refs    []*Object // outgoing references to other objects
}

func newObject(id int) *Object {
    return &Object{id: id}
}

// Ref adds a reference from this object to another.
// This is how you build the object graph manually in tests.
func (o *Object) Ref(other *Object) {
    o.refs = append(o.refs, other)
}

func (o *Object) ID() int { return o.id }