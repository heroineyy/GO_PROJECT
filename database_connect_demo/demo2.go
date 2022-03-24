package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//mysql连接信息配置： 用户名、密码、ip、端口、库名、表名
const (
	strUserName  = "root"
	strPassword  = "root123"
	strIP        = "127.0.0.1"
	strPort      = "3306"
	strDBName    = "db1"
	strTableName = "t1"
)

//插入数据量配置
const (
	TOTAL_INSERT_NUM = 1200000 //共插入多少行数据
	PER_INSERT_NUM   = 5000    //单次向mysql插入多少行数据
	MAX_FAILNUM      = 10      //最大容许插入失败次数
)

//mysql整型范围 供预置数据范围使用
const (
	BIGINT_MIN = -9223372036854775808
	BIGINT_MAX = 9223372036854775807

	INT_MIN = -2147483648
	INT_MAX = 2147483647

	MEDIUMINT_MIN = -8388608
	MEDIUMINT_MAX = 8388607

	SMALLINT_MIN = -32768
	SMALLINT_MAX = 32767

	TINYINT_MIN = -128
	TINYINT_MAX = 127
)

//随机字符串种子

const STRCHAR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&*+-./:;<=>?@[]^_{|}~"

var CITYS = []string{"ChengDu", "KunMing", "XiAn", "LaSa", "JiNan", "NanJing", "HangZhou", "FuZhou", "GuangZhou",
	"HaiKou", "HaErBin", "ChangChun", "ShenYang", "ZhengZhou", "HeFei", "WuHan", "ChongQing", "BeiJing", "ShangHai"}

//随机生成一个整型数据

func MakeRandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

//随机生成一个浮点型数据

func MakeRandFloat(base int64) float32 {
	return rand.Float32() * float32(base)
}

//随机生成一个双精度浮点型数据

func MakeRandDouble(base int64) float64 {
	return rand.Float64() * float64(base)
}

//随机生成一个汉字字符串  入参：字符串长度

func MakeChineseString(length int) string {
	a := make([]rune, length)
	for i := range a {
		a[i] = rune(MakeRandInt(19968, 40869))
	}
	return string(a)
}

//随机生成一个字符串（指定字符种子） 入参：（字符串长度，长度是否随机）

func MakeRandString(length int64, bRegular bool) string {
	var size int64
	if bRegular {
		size = length
	} else {
		size = MakeRandInt(1, length)
	}

	str := make([]byte, size)
	for i := 0; i < int(size); i++ {
		index := MakeRandInt(0, int64(len(STRCHAR)))
		str[i] = STRCHAR[index]
	}
	return string(str)
}

//随机生成一个字符串（任意字符） 入参：字符串长度

func MakeRandString2(length int) string {
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		index := MakeRandInt(0, 127)
		str[i] = byte(index)
	}
	return string(str)
}

//随机生成一个日期类型数据

func MakeRandDate() string {
	year := MakeRandInt(1970, 2021)
	month := MakeRandInt(1, 12)
	var day int64
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = MakeRandInt(1, 31)
	case 4, 6, 9, 11:
		day = MakeRandInt(1, 30)
	case 2:
		day = MakeRandInt(1, 28)
	}

	strDate := strconv.FormatInt(year, 10) + "-" + strconv.FormatInt(month, 10) + "-" + strconv.FormatInt(day, 10)
	return strDate
}

//生成一个递增数据
var iValue int64

//初始化递增数据的值 入参：初始值

func InitIncreaseInt(start int64) {
	iValue = start
}

//生成递增数据 入参：递增量

func MakeIncreaseInt(stage int64) int64 {
	iRet := iValue
	iValue += stage
	return iRet
}

//建立数据库连接

func InitDB() *sql.DB {

	source := strings.Join([]string{strUserName, ":", strPassword, "@tcp(", strIP, ":", strPort, ")/", strDBName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err := sql.Open("mysql", source)
	if err != nil {
		panic(fmt.Sprintf("Open Kungate Connection:[%s] failed, error is [%v].", source, err))
	}

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)

	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return nil
	}

	fmt.Println("connnect success")
	return DB
}

//执行插入操作

func DoInsert(dbConn *sql.DB) {
	fmt.Printf("begin insert, total num:[%d]\n", TOTAL_INSERT_NUM)

	startTime := time.Now().Unix()
	failnum := 0

	//需要赋值的字段定义
	var id int64
	var name string
	var birthday string
	var salary float64
	var distence int64
	var description string
	//var city string

	//初始化递增数据
	InitIncreaseInt(-1000)

	//拼接insert语句
	var strInsert string = "insert into " + strTableName + " values"
	var InsertBuf string = strInsert

	for i := 1; i <= TOTAL_INSERT_NUM; i++ {

		//为各个字段制造随机数据
		id = MakeIncreaseInt(1)
		//id = MakeRandInt(1,20)
		name = MakeRandString(20, false)
		//city = CITYS[MakeRandInt(0, int64(len(CITYS)))]
		birthday = MakeRandDate()
		salary = MakeRandDouble(300)
		distence = MakeRandInt(0, SMALLINT_MAX)
		description = MakeRandString(100, false)

		//拼接insert语句中的值
		InsertBuf += fmt.Sprintf(" (%d, '%s','%s',%f, %d, '%s') ", id, name, birthday, salary, distence, description)

		//若达到单次插入行数 执行插入
		if i%PER_INSERT_NUM == 0 {
			InsertBuf += ";"

			_, err := dbConn.Exec(InsertBuf)
			if err != nil {
				fmt.Println(err)
				fmt.Println(InsertBuf)

				failnum++
				if failnum > MAX_FAILNUM {
					return
				}

				InsertBuf = strInsert
				continue
			}

			//重新初始化insert插入语句
			InsertBuf = strInsert
			CurTime := time.Now().Unix()
			fmt.Printf("complete:[%d]   failnum:[%d]  consume:[%ds]\n", i, failnum*PER_INSERT_NUM, CurTime-startTime)
		} else {
			InsertBuf += ","
		}
	}

	endTime := time.Now().Unix()
	fmt.Printf("finished: totalnum:[%d]  consum:[%d]s\n", TOTAL_INSERT_NUM, endTime-startTime)
}

func main() {
	dbConn := InitDB()

	DoInsert(dbConn)
}
