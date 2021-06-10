package model

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

//球赛服务部分





//GetAllGames 是获取全部的球赛,    还有一些功能需要实现……
func GetAllGames() (allgames []*ListResponse, err error) {
	if err := DB.Find(&allgames).Error; err != nil {
		return nil, err
	}
	return
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
type MyClaims struct {
	Username string `json:"username"`
	Role int 
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(username string,role int) (string, error) {
	// 创建一个我们自己的声明
	var MySecret = []byte("夏天夏天悄悄过去")
	c := MyClaims{
		Username:"username",
		Role:"role",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

//GetGameInfo 用于根据球员名字获取球赛信息
func GetGameInfo(playername string) (playerGames []*ListResponse) {
	err := DB.Where("playername = ?", playername).Find(&playerGames).Error
	if err != nil {
		return nil
	}
	return
}



//写函数（查询数据）
//先确定要传入的字段和要返回的字段
//然后创建一个结构体用于保存查询得到的数据
//用gorm查询语句查询结果
//如果查询失败进行错误判断
//成功则将所需字段返回

//写函数（插入数据）
//先确定要传入的字段和要返回的字段
//然后创建一个结构体，将要插入的数据在结构体中初始化
//用gorm查询语句进行插入
//如果插入失败进行错误判断
//成功则将一些前端需要的字段返回





//GetRole 是一个获取用户权限的函数
func GetRole(ID int) (int, error) {
	var temp User
	if err := DB.Table("user").Where("id=?", ID).Find(&temp).Error; err != nil {
		log.Println(err)
		return 4, err
	}
	return temp.Role, nil
}














//Registerplayer 是一个球员登记的函数
func Registerplayer(name string, picture string, teambelong string) int {
	var player Player
	if err := DB.Table("player").Create(&player).Error; err != nil {
		return 0
	}
	return player.ID
}

//Registerteam 是一个球队登记的函数
func Registerteam(teamname string, logo string, players *[]Player) string {
	var team Team
	if err := DB.Table("team").Create(&team).Error; err != nil {
		return " "
	}
	return team.Name
}

//Changeteam 是一个球员换队的函数
func Changeteam()

//GetTeamInfoByname 是一个   根据id    获取单个团队信息的函数
func GetTeamInfoByname(team string) (teaminfo Team) {
	var temp Team
	err := DB.Table("team").Where("Name = ?", team).Find(&temp)
	if err != nil {
		fmt.Println("没有这个队伍")
	}
	return temp

}

//GetPlayerInfoByName 这是一个通过球员姓名获取球员信息的函数
func GetPlayerInfoByName(Name string) (Player, error) {
	var temp Player
	if err := DB.Table("player").Where("play_name=?", Name).Find(&temp).Error; err != nil {
		log.Println(err)
		return Player{}, err
	}
	return temp, nil
}
