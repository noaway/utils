package pool

import (
	"fmt"
)

type ObjPool struct {
	obj chan interface{}
	New func() interface{}
}

func NewObjPool() *ObjPool {
	return &ObjPool{
		obj: make(chan interface{}, 200),
	}
}

func (p *ObjPool) Get() interface{} {
	select {
	case o := <-p.obj:
		return o
	default:
		return p.New()
	}
}

func (p *ObjPool) Put(o interface{}) {
	select {
	case p.obj <- o:
	default:
	}
}

// func main() {
// 	p := NewObjPool()
// 	p.New = func() interface{} {
// 		return new(Student)
// 	}
// 	sync.Pool
// 	s := p.Get().(*Student)
// 	s.Age = 1
// 	s.Id = 1
// 	s.Name = "周杰伦"

// 	fmt.Println(*s)

// 	p.Put(s)

// 	fmt.Println(*s)
// }
