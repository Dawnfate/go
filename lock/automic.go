package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var x int32
	num := 1000
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
}

//atomic_amd64.s
//8行 Lock，这个是cpu级别的锁，是硬件锁（CPU提供的能力）

/*TEXT ·Xaddint32(SB), NOSPLIT, $0-20
JMP  ·Xadd(SB)

TEXT ·Xadd(SB), NOSPLIT, $0-20
MOVQ  ptr+0(FP), BX
MOVL  delta+8(FP), AX
MOVL  AX, CX
LOCK
XADDL  AX, 0(BX)
ADDL  CX, AX
MOVL  AX, ret+16(FP)
RET*/
