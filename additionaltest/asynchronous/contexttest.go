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
			t := time.Now()
			timeStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
			// reschan <- int64(searial)
			fmt.Printf("%s task %d has been cancled\n", timeStr, searial)
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

// ContextTimeoutTest timeout 測試,設定等候多久後逾時
func ContextTimeoutTest() {
	ctx, cancleFunc := context.WithTimeout(context.Background(), 5*time.Second) ///5秒逾時

	defer cancleFunc()
	for i := 0; i < 5; i++ {
		tesk(ctx, i)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("ContextTimeoutTest Finish...")
}

// ContextDeadlineTest 做到甚麼時候為止,給定一個時間,例如幾點幾分...,而不是多長的時間
func ContextDeadlineTest() {
	dl := time.Now().Add(30 * time.Second)
	ctx, cancleFunc := context.WithDeadline(context.Background(), dl)

	defer cancleFunc()
	t := dl
	str := fmt.Sprintf("test print %d", 6)
	fmt.Println(str)
	timeStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Println("goroutings will be stop at ", timeStr)
	for i := 0; i < 5; i++ {
		go tesk(ctx, i)
	}

	time.Sleep(60 * time.Second)
}
