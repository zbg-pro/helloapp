package test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"helloapp/models"
	_ "helloapp/routers"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	//mysql
	/*orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		"root:iPYDU0o3MRQOreEW@tcp(172.16.100.130:3306)/c2c?charset=utf8")*/

	//sqlite
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", getCurrentAbPath()+"/sqlite3.db")

	orm.Debug = true
	orm.RegisterModel(new(ActivityBill))
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestDateTime(t *testing.T) {
	u := "user_" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	fmt.Println(u)
	Pwd1 := "admin"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	Pwd1 = hex.EncodeToString(cipherStr)
	fmt.Println(Pwd1)
	fmt.Println(string(cipherStr))

	orm.RegisterDataBase("default", "mysql", "root:iPYDU0o3MRQOreEW@tcp@tcp(172.16.100.130:3306)/c2c?charset=utf8")
	o := orm.NewOrm()

	//orm.NewOrmUsingDB("dbname")
	var user1 models.User
	o.QueryTable("user").Filter("username", "").One(&user1, "Id")
}

func TestMysqlConn(t *testing.T) {
	//https://www.cnblogs.com/qidaii/articles/15633605.html

	o := orm.NewOrm()

	o.Raw("create table user(`id` int(20), `user_name` varchar(100), primary key (`id`) )").Exec()
	o.Raw("insert into user values(1, \"zl239\");").Exec()
	var maps []orm.Params
	num, err := o.Raw("select * from user").Values(&maps)
	if err == nil {
		fmt.Println(num)
		fmt.Println(maps)
	}

	var list []orm.ParamsList
	num1, err1 := o.Raw("select * from user").ValuesList(&list)
	if err1 == nil && num1 > 0 {
		fmt.Println(list) // slene
	}
	/*a := o.Raw("select * from apikey where 1=1")
	fmt.Println(a)

	var ak apikey
	o.QueryTable("apikey").Filter("username", "122").One(&ak, "Id")*/

}

type ActivityBill struct {
	Id               int64   `orm:"column(id);pk"`
	TotalChargeMoney string  `orm:"column(totalChargeMoney)"`
	TotalTradeMoney  float64 `orm:"column(totalTradeMoney)"`
	UserId           int     `orm:"column(userId)"`
	ActivitySetId    int64   `orm:"column(activitySetId)"`
	UserName         string  `orm:"column(userName)"`
}

// table activity_bill [[1 12.00000000 13.00000000 1234 22 zl239]]
func TestInsertMysql(t *testing.T) {
	//beego 插入的要点：1 需要建立对应的类 2 需要注册对应的类 3 需要指定类里面的id作为主键，因为插入后要返回id 4 需要在strust指定类的属性对应的表的字段名称
	var bill ActivityBill
	bill.ActivitySetId = 1
	bill.Id = time.Now().UnixMilli()
	bill.UserId = 100
	bill.UserName = "zl2391"

	o := orm.NewOrm()
	num, err := o.Insert(&bill)
	fmt.Println("num:", num, " err:", err)

}

func TestUpdateMysql(t *testing.T) {
	o := orm.NewOrm()
	var bill ActivityBill
	bill.ActivitySetId = 12
	bill.Id = 1 //time.Now().UnixMilli()
	bill.UserId = 200
	dec, _ := decimal.NewFromString("123.33333777")
	bill.TotalChargeMoney = dec.String()
	f, _ := dec.Float64()
	bill.TotalTradeMoney = f
	bill.UserName = "zl2391"
	fmt.Println(bill)
	eff, err := o.InsertOrUpdate(&bill)
	fmt.Println("eff:", eff, " err:", err)
}

func TestDeleteMysql(t *testing.T) {
	o := orm.NewOrm()
	var bill ActivityBill
	bill.Id = 0
	o.Delete(&bill)
}

func TestQueryMysql(t *testing.T) {
	o := orm.NewOrm()

	//分为2种：转对象的，转普通valueList的
	query := o.QueryTable("activity_bill")
	var bill ActivityBill
	query.Filter("id", 1).One(&bill)
	fmt.Println(bill)

	//value或者valueList
	var list []orm.ParamsList
	x := o.Raw("select * from activity_bill where   username like ?", "%zl239%")
	x.ValuesList(&list)
	fmt.Println(list)
	for i := 0; i < len(list); i++ {
		var t = list[i]
		for j := 0; j < len(t); j++ {
			fmt.Print(", ", t[j])
		}
		fmt.Println("")
	}

	var maps []orm.Params
	num, err := x.Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps)
		for i := 0; i < len(maps); i++ {
			fmt.Println(maps[i]["id"])
			for k, v := range maps[i] {
				fmt.Println(k, v)
			}
			fmt.Println("=============")
		}
	}

}

func TestInsertSqlite(t *testing.T) {

}

func TestUpdateSqlite(t *testing.T) {

}

func TestDeleteSqlite(t *testing.T) {

}

func TestQuerySqlite(t *testing.T) {

}

func TestRedis(t *testing.T) {

}

func TestPath(t *testing.T) {
	fmt.Println(getCurrentAbPath())
	fmt.Println(getCurrentAbPathByExecutable())
	fmt.Println(getCurrentAbPathByCaller())
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
