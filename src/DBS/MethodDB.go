package DBS

import "fmt"

func StructInsert(url string,name string,code string){
	sqlStr := "insert into dataset(url,name,code) values(?,?,?)"
	ret,err :=Db.Exec(sqlStr,url,name,code)
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
