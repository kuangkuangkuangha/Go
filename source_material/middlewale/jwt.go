package middlewale

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)



var jwtkey = []byte("nihao")

//Myclaims 是自定义的结构体
type Myclaims struct{
	Username string `json:"name"`
	jwt.StandardClaims
}
//SetToken 生成Token
func SetToken(username string)(string){
	setclaim := Myclaims{
		Username :username ,
		StandardClaims: jwt.StandardClaims{
			Issuer: "teammanage",
		},
	}
	//选择加密方式和签名参数
	//返回的是一个*Token，他是一个结构体
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodES256, setclaim)

	token,err := reqClaims.SignedString(jwtkey)

	if err != nil {
			return ""
	}
	
	return token

}


//CheckToken 验证Token
func CheckToken(token string)(*Myclaims,){
	settoken, _ := jwt.ParseWithClaims(token,&Myclaims{}, func(token *jwt.Token) (interface{},error){
		return jwtkey, nil
	})

	if key, ok:=settoken.Claims.(*Myclaims);ok && settoken.Valid{
		return key
	}else {
		return nil
	}

}


//jwt中间件
func JwtToken() gin.HandlerFunc{
	
}