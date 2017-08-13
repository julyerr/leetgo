package main 
// sync/atomic mux to each goroutine
// AddUint64(&var,delta) LoadUint64(&var)
import (
	"fmt"
	"time"
	"sync/atomic"
)

func main(){
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for{
				atomic.AddUint64(&ops,1)
				time.Sleep(time.Millisecond)
			}			
		}
	}
	time.Sleep(time.Second)
	opsFinal:=atomic.LoadUint64(&ops)
	fmt.Println("ops:",opsFinal)
}




