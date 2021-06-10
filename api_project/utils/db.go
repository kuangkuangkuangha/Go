package utils
import(
	"database/sql"
	_"go-sql-driver/mysql"
)


// Db 是；连接词
var (
	Db *sql.DB
	err error
)
//Init 函数是将数据库初始化
func Init(){
    //这里面的Db就是上面的Db
     Db, err = sql.Open("mysql", "root:zk2824895143@tcp(localhost:3306)/user")
     if err != nil{
	   panic(err.Error())
     }
    //  err =Db.Ping()
    //  if err != nil{
	//    return
	//  }
	//  return
}	


