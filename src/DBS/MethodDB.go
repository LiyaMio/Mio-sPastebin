package DBS

import (
	"fmt"
	"os"
	"time"
)
type Database struct {
	Username string
	Password string
	Server   string
	MysqlDb  *sql.DB
}
func (db *Database) QuerySingleRow(sqlStr string) *sql.Row {
	rows := db.MysqlDb.QueryRow(sqlStr)
	return rows
}
func (db *Database) QueryMultiRow(sqlStr string) *sql.Rows {
	rows, err := db.MysqlDb.Query(sqlStr)
	if err != nil {
		fmt.Printf("preare failed, err: %v\n", err)
	}
	return rows
}
func (db *Database) InsertRow(sqlStr string) string{
	_, err := db.MysqlDb.Exec(sqlStr)
	fmt.Println(sqlStr)
	if err != nil {
		log.Println("insert failed , err:%v\n", err)
		return "insert failed , err"
	}else{
		return ""
	}
}
func (db *Database) UpdateRow(sqlStr string) {
	ret, err := db.MysqlDb.Exec(sqlStr)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
	}
	n, err := ret.RowsAffected() // 操作收影响的行
	if err != nil {
		fmt.Printf("get RowsAffected failed: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}
func (db *Database) QueryDB(sqlStr string) ([]map[string]string, bool) {
	rows, err := db.MysqlDb.Query(sqlStr) //执行SQL语句，比如select * from users
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i, _ := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得这些内存的指针，因后继的Scan函数只接受指针
	}
	ret := make([]map[string]string, 0) //创建返回值：不定长的map类型切片
	for rows.Next() {
		err := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := make(map[string]string) //用于存放1列的 [键/值] 对
		if err != nil {
			panic(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			b, _ := raw_value.([]byte)
			v := string(b) //将raw数据转换成字符串
			m[colName] = v //colName是键，v是值
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}

	defer rows.Close()

	if len(ret) != 0 {
		return ret, true
	}
	return nil, false
}
func StructInsert(url string,name string,code string,){
	sqlStr := "insert into dataset(url,name,code,nowtime) values(?,?,?,?)"
	now:=time.Now()
	//formatTimeStr:=time.Unix(now,0).Format("20060102150405")
	ret,err :=Db.Exec(sqlStr,url,name,code,now)
	if err !=nil {
		fmt.Printf("insert failed , err:%v\n", err)
		return
	}
	theId,err :=ret.LastInsertId()
	if err !=nil{
		fmt.Printf("get lastinsert Id failed, err: %v\n,", err)
	}
	fmt.Println("insert success, the id is:", theId)
}
func QueryRowByUrl(url string) string{
	sqlStr := "select code from dataset where url = ?"
	var code string
	err := Db.QueryRow(sqlStr,url).Scan(&code)
	if err != nil {
		var str = "scan failed"
		fmt.Printf("scan failed, err: %v\n", err)
		return str
	}
	fmt.Printf("%s",code)
	return code
}
func QueryRowByName(name string) string{
	sqlStr := "select name from dataset where name = ?"
	var code string
	err := Db.QueryRow(sqlStr,name).Scan(&code)
	if err != nil{
		var str = "scan failed"
		fmt.Printf("Scan failed,err:%v/n",err)
		return str
	}
	fmt.Printf("%s",code)
	return code
}
func QueryRowTime()[]string{
	sqlStr := "select url from dataset where SUBDATE(now(),interval 2 hour) >  nowtime"
	var ans []string
	err:=Db.QueryRow(sqlStr).Scan(&ans)
	if err != nil{
		var str []string
		str[0]="scan "
		str[1]="failed"
		fmt.Printf("Scan failed,err:%v/n",err)
		return str
	}
	fmt.Printf("%s",ans)
	fmt.Println("查询时间成功")
	return ans
}
func deleteRowDemo(code string) {
	sqlStr := "delete from dataset where code=?"
	ret, err := Db.Exec(sqlStr, code)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 获取操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected fail, err:%v\n", err)
	}
	fmt.Printf("delete success, affected rows: %v \n", n)

}
func QueryMutiRowTime() {
	sqlStr := "select code from dataset where SUBDATE(now(),interval 5 minute ) > nowtime"
	var timeList []string
	rows,err:= Db.Query(sqlStr)

	if err !=nil{
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ans string
		err := rows.Scan(&ans)
		if err != nil {
			fmt.Println("Scan failed,err:%v\n", err)
			return
		}
		timeList =append(timeList,ans)
		os.Remove(ans)
		deleteRowDemo(ans)
		fmt.Println("url : \n",ans)
	}
}
func UpdateTime(url string){
	sqlStr := "update dataset set nowtime = ? where url = ?"
	ret,err := Db.Exec(sqlStr,time.Now(),url)
	if err !=nil{
		fmt.Println("uoadte failed,err: %v\n",err)
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get RowsAffected failed: %v\n",err)
		return
	}
	fmt.Println("update success,affected rows: \n",n)
}
