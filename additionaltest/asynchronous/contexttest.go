package asynchronous

import (
	"context"
	"fmt"
	"time"
)

func tesk(ctx context.Context, searial int) {
	var x int64 = 0
	for {
		x++ ///do anything...
		select {
		case <-ctx.Done():
			// reschan <- int64(searial)
			fmt.Printf("task %d has been cancled\n", searial)
			return
		default:
		}
	}
}

// ContextWithCancelTest cancle測試
func ContextWithCancelTest() {
	var routingNum int = 5 ///畢竟只是要展示cancle的用法,所以gorouting 別放太多,不然cancle的時候有機會造成主程式跑完了 其他的卻還正在結束
	///先基本的要一個可以取消的context
	ctx, cancleFunc := context.WithCancel(context.Background())

	for i := 0; i < routingNum; i++ {
		go tesk(ctx, i)
	}

	time.Sleep(time.Second)
	cancleFunc()
	// resVal := <-result
	fmt.Println("Finish...")
}
