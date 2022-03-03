package core

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestInitMyPool(t *testing.T) {
	InitMyPool()
	ts := time.Now().UnixNano()
	a := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		a.Add(1)
		go func() {
			for i := 0; i <= 10; i++ {
				a := new(MyExtPoll)
				a.A = make(map[string]int)
				a.A["hahah"] = i
				a.A["hahah1"] = i
				a.A["hahah2"] = i
				a.A["hahah3"] = i
				a.A["hahah4"] = i
				a.A["hahah5"] = i
				a.B = []string{"fdsafdsa", "24rdfsgrfdsf", "24rdffdsgrfdsf", "24rdfsgrqerfdsf", "24rdfsgrewqrfdsf", "24rdf432143sgrfdsf"}
				// fmt.Println(&a)
			}
			a.Done()
		}()

	}

	fmt.Println(time.Now().UnixNano() - ts)
	ts2 := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		a.Add(1)
		go func() {
			for i := 0; i <= 10; i++ {
				a := MyPool.GetSt()
				a.A = make(map[string]int)
				a.A["hahah"] = i
				a.A["hahah1"] = i
				a.A["hahah2"] = i
				a.A["hahah3"] = i
				a.A["hahah4"] = i
				a.A["hahah5"] = i
				a.B = []string{"fdsafdsa", "24rdfsgrfdsf", "24rdffdsgrfdsf", "24rdfsgrqerfdsf", "24rdfsgrewqrfdsf", "24rdf432143sgrfdsf"}
				MyPool.Put(a)
				// fmt.Println(&a)
			}
			a.Done()
		}()
	}
	fmt.Println(time.Now().UnixNano() - ts2)
	a.Wait()

}
