package asynchronous

import (
	"fmt"
	"time"
)

func task(val chan int) {
	x := 0
	for i := 0; i < 100; i++ {
		x = i
	}
	val <- x
}

// ChannelTest 基本展示channel
func ChannelTest() {
	cht := make(chan int)

	go task(cht)
	res := <-cht

	fmt.Println("res = ", res)
}

func worker1(ch chan int) {
	var data string = ""
	data = fmt.Sprintf("worker1 : Start...\n")
	fmt.Println(data)
	// fmt.Println(data)
	for {
		data = fmt.Sprintf("worker1 : Start loop\n")
		fmt.Println(data)
		select {
		case x := <-ch:
			data = fmt.Sprintf("worker1 : got message x = %d\n", x)
			fmt.Println(data)

			if x < 0 {
				data = fmt.Sprintf("worker1 : going to exit,\n")
				fmt.Println(data)
				ch <- x

				return
			}
			x++
			data = fmt.Sprintf("worker1 : will send message %d\n", x)
			fmt.Println(data)
			ch <- x
			data = fmt.Sprintf("worker1 : send message %d\n", x)
			fmt.Println(data)
		default:
			data = fmt.Sprintf("worker1 : wait for message\n")
			fmt.Println(data)
			time.Sleep(time.Second)
		}
		data = ""
	}
}

func worker2(ch chan int) {
	var data string = ""
	data = fmt.Sprintf("worker2 : Start...\n")
	fmt.Println(data)
	for {
		data = fmt.Sprintf("worker2 : Start loop\n")
		fmt.Println(data)
		select {
		case x := <-ch:
			data = fmt.Sprintf("worker2 : got message x = %d\n", x)
			fmt.Println(data)
			if x < 0 {
				data = fmt.Sprintf("worker2 : going to exit\n")
				fmt.Println(data)
				ch <- x

				return
			}
			x++
			time.Sleep(3 * time.Second)
			data = fmt.Sprintf("worker2 : will send message %d\n", x)
			fmt.Println(data)
			ch <- x
			data = fmt.Sprintf("worker2 : was send message %d\n", x)
			fmt.Println(data)
		default:
			data = fmt.Sprintf("worker2 : wait for message\n")
			time.Sleep(time.Second)
		}
		data = ""
	}
}

// ChannelBufferOrUnbufferTest 測試建立channel當下給予buffer或者讓他使用預設值,運作的狀況
/*
總結:channel在實體化當下
ch := make(chan int) : 這是預設做法,unbuffer,這種做法channel本身內部不會附帶buffer因此不論是讀 <- ch, 或者是寫 ch <- 都需要等到另一個部分的出現
ch := make(chan int, 1) : 這種作法讓channel使體化當下內部就有buffer可以寫資料,因此 ch <- 不需要等待 <- ch
實際上運作方式比較像是
var mtx sync.Mutex
type channel struct {
	buffer BufferElement =  nil
}
func (ch *channel)read() interface {
		mtx.Lock()
		defer mtx.Unlock()
		if ch.buffer == nil {
			ch.buffer = [1]BufferElement
		}
		for ch.buffer is empty {
			wait data
		}
		read data
	}
	func (ch *channel)write( arg ... interface) {
		mtx.Lock()
		defer mtx.Unlock()
		for buffer == nil {
			wait buffr
		}
		write data
	}
*/
func ChannelBufferOrUnbufferTest() {
	x := 0
	ch := make(chan int)
	// ch := make(chan int, 1)

	go worker1(ch)
	go worker2(ch)
	ch <- x

	time.Sleep(30 * time.Second)
	/*回圈內的程式碼在make(chan int),與 make(chan int, 1)的情況下有所不同
	make(chan int) 會造成deadlock,因為寫入資料的 ch <- 要等後面的 <- ch,但實際上根本等不到
	make(chan int, 1) 可以正常執行,但相對的 前面程式的兩個go worker() 就會亂跳,例如某個worker自己輸入資料自己取走,或者各自做一個 >1的次數後交換,或者照原本的想法輪流做一次...
	*/
	/*
		ch <- x
		for {
			select {
			case <-ch:
				x++
				time.Sleep(time.Second)
				ch <- x
			default:
				fmt.Println("Not thing to do")
			}
		}
	*/
	fmt.Println("Notify worker finish")

	<-ch
	ch <- -1
	time.Sleep(20 * time.Second)
	<-ch
	fmt.Println("Finish")
}
