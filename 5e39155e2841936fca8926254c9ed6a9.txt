package DBS

import (
	"fmt"
	"os"
	"time"
)
func StructInsert(url string,name string,code string){
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
	sqlStr := "select code from dataset where SUBDATE(now(),interval 1 hour ) > nowtime"
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
		deleteRowDemo(ans)
		os.Remove(ans)
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
	fmt.Println("update success,affected rows: %d\n",n)
}