package main

import (
    "fmt"
    "net/http"
    // _ "go-sql-driver/mysql"
    "database/sql"
	
)

//Db 是；连接词
var Db *sql.DB
type user struct{
	id int
	username string
	password string
	role string
}
func connectmysql()(err error){
    //这里面的Db就是上面的Db
     Db, err = sql.Open("mysql", "root:zk2824895143@tcp(localhost:3306)/user")
     if err != nil{
	 return
     }
     err =Db.Ping()
     if err != nil{
	   return
	 }
	 return
}	

func cheakusername (username string ,password string )(*user,error){
//查询单条记录
    var a user
	sqlStr :="select id , username , password, role from login where username = ? & password =?"
	row , _:= Db.Query(sqlStr,username,password)
	//将扫描的结果赋给a，括号里的是（a.username）是赋值的意思
	row.Scan(&a.id,&a.username,&a.password,&a.role)
	fmt.Printf("a:%#v\n",a)
	return &a ,nil
	
}

//去登陆页面
func goweb(w http.ResponseWriter, r *http.Request){
	html:=`<html>
    <head>
        <meta charset="UTF-8"/>
    </head>
    <body>
        <form action="http://localhost:8080/login" method="POST">
            用户名：<input type="text" name="username " /><br/>
            密码：<input type="password" name="password" /><br/>
              <input type="submit"/>
        </form>
    </body>
</html>`
	
	w.Write([]byte(html))
}

//登陆
func login(w http.ResponseWriter, r *http.Request){
	
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
    b,_:=cheakusername(username,password)
   if  a.id=nil{
       fmt.Fprintln(w,"登入成功")
	//    w.Write([]byte("登入成功"))
   }else {
	w.Write([]byte("登入失败"))
   }
//    if err!=nil{
// fmt.Fprintf("err:%v",err)
//    }

}

// //创建处理器函数
//     func goweb(w http.ResponseWriter, r *http.Request){
//         html:=`<html>
//         <head>
//             <meta charset="UTF-8"/>
//         </head>
//         <body>
//             <form action="http://localhost:8080" method="POST">
//                 用户名：<input type="text" name="username " /><br/>
//                 密码：<input type="password" name="password" /><br/>
//                   <input type="submit"/>
//             </form>
//         </body>
//     </html>`
        
//         w.Write([]byte(html))
//     }


func main() {
    err :=connectmysql()
		if err != nil{
			fmt.Printf("init DB failed,err:5v%\n",err)
		}
		fmt.Println("数据库连接成功")
	

    http.HandleFunc("/", goweb)
    http.HandleFunc("/login",login)
    //创建路由
    http.ListenAndServe(":8080", nil)
}