package main

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"

	"html/template"
	"net/http"
)

// Db 是连接池
var Db *sql.DB

type user struct {
	id       string
	username string
	password string
	role     string
}

func connectmysql() (err error) {
	Db, err = sql.Open("mysql", "root:zk2824895143@tcp(localhost:3306)/user")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

//CheckUsernameAndPassword 函数
func checkUsernameAndPassword(username string, password string) *user {
	sqlStr := "select id , username , password, role from login where username = ? and password =?"
	row := Db.QueryRow(sqlStr, username, password)
	a := &user{}
	row.Scan(&a.id, &a.password, &a.role, &a.username)
	return a
}

//CheckUser 函数
func checkUser(username string) (a *user) {
	sqlStr := "select id , username , password, role from login where username = ?"
	row := Db.QueryRow(sqlStr, username)
	a = &user{}
	row.Scan(&a.id, &a.password, &a.role, &a.username)
	return a
}

// 登陆的处理器终于写好了！！！
func login(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	ret := checkUsernameAndPassword(username, password)
	if ret.id != "" {
		fmt.Fprintln(w, "登陆成功")

	} else {
		fmt.Fprintln(w, "登陆失败")
	}
}

// func login(w http.ResponseWriter,r *http.Request){
// 	b := r.PostFormValue("password")
// 	if b != "qqq" {
// 		  fmt.Fprintln(w,"密码不是qqq")
// 	}else{
// 		fmt.Fprintln(w,"密码是qqq")
// 	}

// }

func addUser(id string, username string, password string, role string) {
	// sqlStr :="insert into login(username,password,role) values(?,?,?,?)"
	// _,err:=Db.Exec(sqlStr,id,username,password,role)
	// if err != nil{
	// return
	// }else{
	// 	fmt.Println("插入成功")
	// }
	sqlStr := "insert into login(id,username,password,role)values(?,?,?,?)"
	ret, _ := Db.Exec(sqlStr, id, username, password, role)
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)

		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)

}

//注册的函数也写好啦！
func register(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	role := r.PostFormValue("role")
	a := checkUser(username)
	if a.id != " " {
		addUser(id, username, password, role)
		fmt.Fprintln(w, "注册完成")
		t, _ := template.ParseFiles("heool.html")
		t.Execute(w, "欢迎成为书城的一元")
	} else {
		fmt.Fprintln(w, "用户名已存在")
	}

}

// //去登入界面
// func goweb(w http.ResponseWriter, r *http.Request){
// 	//解析模版
// 	t,_:= template.ParseFiles("heool.html")
//    //执行模版
//    t.Execute(w, "hahahahaha")
// }

func main() {
	err := connectmysql()
	if err != nil {
		fmt.Printf("init DB failed,err:5v%\n", err)
	}
	fmt.Println("数据库连接成功")

	// sqlStr :="insert into login(id,username,password,role)values(?,?,?,?)"
	//   ret,_:=Db.Exec(sqlStr,"salt","harden","13","1")
	//   theID, err := ret.LastInsertId() // 新插入数据的id
	// 	if err != nil {
	// 		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
	// 		return
	// 	}
	// 	fmt.Printf("insert success, the id is %d.\n", theID)

	// addUser("3d","klay","11","0")
	// addUser("dyrant","5443","1")

	http.HandleFunc("/http", login)
	http.HandleFunc("/register", register)
	// http.HandleFunc("/goweb",goweb)
	http.ListenAndServe(":8090", nil)

}

// ret,_:=cheakusername("curry","88888888")
// fmt.Println(*ret)
// sqlStr :="select id , username , password, role from login where id = ?"
// row := Db.QueryRow(sqlStr,888)
//将扫描的结果赋给a，括号里的是（a.username）是赋值的意思
// var a user
// row.Scan(&a.id,&a.username,&a.password,&a.role)
// fmt.Printf("a:%#v\n",a)
// ret := checkusername("888")
//    fmt.Printf("ret:%#v\n",ret)

//查询单条记录
// sqlStr :="select id , username , password, role from login where id = ?"
// row , _:= Db.Query(sqlStr,666)
// //将扫描的结果赋给a，括号里的是（a.username）是赋值的意思
// var a
// row.Scan(&a.id,&a.username,&a.password,&a.role)
// fmt.Printf("a:%#v\n",a)

// //去登陆页面
// func goweb(w http.ResponseWriter, r *http.Request){
// 	html:=`<html>
//     <head>
//         <meta charset="UTF-8"/>
//     </head>
//     <body>
//         <form action="http://localhost:8080/login" method="POST">
//             用户名：<input type="text" name="username " /><br/>
//             密码：<input type="password" name="password" /><br/>
//               <input type="submit"/>
//         </form>
//     </body>
// </html>`

// 	w.Write([]byte(html))
// }

//登陆
// func login(w http.ResponseWriter, r *http.Request){

// 	username := r.FormValue("username")
// 	password := r.FormValue("password")
//     b,err:=cheakusername(username,password)
//    if  b!=nil{
// 	   w.Write([]byte("登入成功"))
//    }else {
// 	w.Write([]byte("登入失败"))
//    }
//    if err!=nil{

//    }

// }

//    //查询多条记录
//    sqlStr :="select id ,username ,password, role from login where id > ?"
//    rows , _:= Db.Query(sqlStr,0)
// 	defer rows.Close()
// 	//调用Next方法
// 	for rows.Next() {
// 		var u user
// 		//调用Scan方法
// 		err := rows.Scan(&u.id, &u.username,&u.password, &u.role)
// 		if err != nil {
// 			fmt.Printf("scan failed, err:%v\n", err)
// 			return
// 		}
// 		fmt.Printf("id:%d name:%s password:%s age:%s\n", u.id, u.username,u.password, u.role)
// 	}
// }

// //插入数据
//   sqlStr :="insert into login(id,username,password,role)values(?,?,?,?)"
//   ret,_:=Db.Exec(sqlStr,	888,"curry",88888888,1)
//   theID, err := ret.LastInsertId() // 新插入数据的id
// 	if err != nil {
// 		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("insert success, the id is %d.\n", theID)

// }

//删除数据
// sqlStr := "delete from login where id =999"
// //执行
// ret,err:=Db.Exec(sqlStr)

// if err !=nil{
// 	fmt.Printf("插入失败，err:%v",err)
// 	return
// }
// fmt.Printf("删除成功%v",ret)
// }

//修改数据
// sqlStr := "update login set username ='james'where id = 12345"
// ret,err:=Db.Exec(sqlStr)
// if err != nil {
// 	fmt.Println("修改失败")
// 	return
// }
// fmt.Printf("修改成功，%v",ret)
// }
// http.HandleFunc("/hello",goweb)
//http.HandleFunc("/login",login)
// http.ListenAndServe("8000",nil)
// }
