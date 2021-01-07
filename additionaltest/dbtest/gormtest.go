package dbtest

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// programTestTable 此結構對應表格原先並不存在,交由gorm自己去創造
// 此類表格 `gorm:"column:ID"` 這種寫法就沒有很強烈的建議要寫,畢竟多半情況下它自己能對映到
// 此處有一點特別注意,文件上說零值的欄位會去讀取預設值,其實是指我們設定給GORM的預設值,實際上恐怕不是從db中取得的預設值,因此這裡如果忘記寫 default:xxx , 那麼gorm會依照golang的設定去給預設
// 字串為 "" 數字為0...,這將造成 資料庫中允許NULL,且唯一值的欄位會發生我們預想是應該被填寫NULL,且可以寫入,但實際上它卻寫寫入 "" 而發生值重複無法寫入的錯誤
type programTestTable struct {
	//gorm.Model
	ID           int    `gorm:"primary_key:ID; AUTO_INCREMENT:number"`
	Name         string `gorm:"type:varchar(50); not null;default:''"`
	HomeAddress  string `gorm:"type:varchar(100);not null;default:''"`
	CerditCardNO string `gorm:"type:varchar(20);unique_index;default:NULL"`
}

// TestTable  原始測試表,非GORM建立
// 由於非gorm創建的表格,因此表的欄位名稱與結構的欄位對應建議自己寫上會比較妥當,避免因為認定上不同造成gorm找不到對應的欄位而沒把撈到的資料寫回來
// 重要的一點, golang中或許不會限制你變數名稱非得開頭大寫,但是GORM中這是必須,開頭小寫的變數即使後面的`gorm:"..." 中的內容都對,但就是會發生有撈取到資料但不會寫入回傳的變數中的情況`
type TestTable struct {
	//gorm.Model
	ID   int    `gorm:"column:ID;primary_key:ID; AUTO_INCREMENT:number"`
	Name string `gorm:"column:Name;type:varchar(50)"`
}

// type TestTable struct {
// 	//gorm.Model
// 	ID   int    `gorm:"primary_key:ID; AUTO_INCREMENT:number"`
// 	Name string `gorm:"type:varchar(50)"`
// }

// GormTest GORM使用
func GormTest() {
	db, err := gorm.Open("mysql", "root:a13675452@tcp(192.168.100.12:3306)/LearnDB")
	if err != nil {
		fmt.Println("DB Connection fail, err = ", err)
		return
	}
	///連線成功先做
	defer db.Close()
	///開啟單數表名,防止它幫我們把表明稱後面+ s
	db.SingularTable(true)

	///==========gorm 建立表格測試區================
	///建立表
	if !db.HasTable(&programTestTable{}) { ///先判斷對應結構的表是否存在
		///不存在...創建
		if err := db.CreateTable(&programTestTable{}).Error; err != nil { ///判斷錯誤
			fmt.Println("Create fail Error = ", err)
		}
		fmt.Println("Create Table Success")
	} else {
		///存在,寫LOG就好,不用再做了
		fmt.Println("Table was already exist, do not need to create!!!")
	}
	///由於交由gorm建立的表格ID欄位是自動增量,因此這裡一律寫 0,讓他再新增時忽略該欄位的值得寫入,當新增成功後我們就可以從原本變數中ID欄位獲得新的自動增量值了
	/*
		data := programTestTable{ID: 0, Name: "bbb"}
		if err := db.Debug().Create(&data).Error; err != nil {
			fmt.Println("Insert fail Error = ", err)
		} else {
			fmt.Println("insert Data Success, new data = ", data)
		}
	*/

	resData := programTestTable{}
	if err := db.Debug().First(&resData).Error; err != nil {
		fmt.Println("Query fail Err = ", err)
	} else {
		fmt.Println("Query success, resData = ", resData)
	}

	///Where 條件式,內容用FirstOrInit()找到的,或者找不到就用Assign(...)的去設定變數內容,但無論如何都不會寫入DB
	qUser := programTestTable{}
	qUser.Name = "cccc"
	db.Debug().Where(qUser).Assign(programTestTable{ID: -1, Name: "Not exist"}).FirstOrInit(&qUser)
	fmt.Println("qUser = ", qUser)

	///找Name = John 如果找的到就把他的資料撈出來,不然就創一筆,然後撈出他的資料
	newUser := programTestTable{Name: "John"}
	db.Debug().FirstOrCreate(&newUser, newUser)
	fmt.Println("new User = ", newUser)

	///刪除
	delUser := programTestTable{ID: 5}
	db.Debug().Delete(delUser)
	///==========gorm 建立表格測試區================

	///==========非 gorm 建立的表格測試區=============
	testData := make([]TestTable, 1)
	// db.CreateTable(&pct)
	///實在不清楚到底用了甚麼語句查詢 可以先在前面加入 Debug(). 後面再接上原本的查詢動作,以此看清楚GORM到底送了什麼查詢語法給資料庫,不使用記得拿掉
	resDB := db.Debug().Table("TestTable").Find(&testData)

	if resDB.Error != nil {
		fmt.Println("Error : ", resDB.Error)
	} else {
		fmt.Println("testData = ", testData, "resDB = ", resDB)
	}
	///==========非 gorm 建立的表格測試區=============
}
