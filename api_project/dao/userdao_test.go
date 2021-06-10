package dao
import(
	"fmt"
"testing"
)
func TestUser(t *testing.T){
	  fmt.Println("测试userdao中的函数")
	  t.Run("验证用户名和密码",testLogin)
	  t.Run("验证用户名",testregister)
	  t.Run("验证用户名和密码",testsave)
}

func testLogin(t *testing.T){
	 user,_:=ch("curry","88888888")
	 fmt.Println("获取的用户信息是",user)
}
func testregister(t *testing.T){
	user,_:=("curry","88888888")
	fmt.Println("获取的用户信息是",user)
}
func testsave(t *testing.T){
	user,_:=CheckUsernameAndPassword("curry","88888888")
	fmt.Println("获取的用户信息是",user)
}