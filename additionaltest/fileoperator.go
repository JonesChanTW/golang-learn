package additionaltest

import (
	"fmt"
	"os"
	"sync"
)

// FileOpenOrCreateAndWrite 檔案操作
func FileOpenOrCreateAndWrite(data string) {
	file, fOpenErr := os.OpenFile("fileOpt.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if fOpenErr != nil {
		fmt.Println("file open error : ", fOpenErr)
		return
	}
	defer file.Close()
	res, err := file.WriteString(data)
	if err != nil {
		fmt.Println("file open error : ", err)
		return
	}
	fmt.Println("file open and write success, and res = ", res)
}

func FileOpenOrCreateAndWriteWithMutex(data string) {
	var mtx sync.Mutex

	mtx.Lock()
	defer mtx.Unlock()
	file, fOpenErr := os.OpenFile("fileOpt.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if fOpenErr != nil {
		fmt.Println("file open error : ", fOpenErr)
		return
	}
	defer file.Close()
	res, err := file.WriteString(data)
	if err != nil {
		fmt.Println("file open error : ", err)
		return
	}
	fmt.Println("file open and write success, and res = ", res)

}
