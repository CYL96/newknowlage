package core

import "sync"

type MyExtPoll struct {
	A  map[string]int
	Ha struct {
		A int
		B int
	}
	C string
	D string
	B []string
}

var MyPool *MyPoolExt

type MyPoolExt struct {
	*sync.Pool
}

func (this *MyPoolExt) Get() interface{} {
	st := this.Pool.Get().(*MyExtPoll)
	st.A = make(map[string]int)
	st.Ha.A = 0
	st.Ha.B = 0
	st.C = ""
	st.D = ""
	return st
}
func (this *MyPoolExt) GetSt() *MyExtPoll {
	st := this.Pool.Get().(*MyExtPoll)
	st.A = make(map[string]int)
	st.Ha.A = 0
	st.Ha.B = 0
	st.C = ""
	st.D = ""
	return st
}

// func (this *MyPoolExt) Put(inf interface{}) {
// 	st := inf.(*MyExtPoll)
// 	st.A = make(map[string]int)
// 	st.Ha.A = 0
// 	st.Ha.B = 0
// 	st.C = ""
// 	st.D = ""
// }

func InitMyPool() {
	MyPool = new(MyPoolExt)
	MyPool.Pool = &sync.Pool{
		New: func() interface{} {
			data := new(MyExtPoll)
			data.A = make(map[string]int)
			return data
		},
	}

}
