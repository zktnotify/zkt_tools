package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/web_zktnotify/internal/app/model"
	"time"
)

//var Secret = os.Getenv("JWT_SECRET")
var Secret = "abc123"

func Handler(ctx iris.Context) {
	//如果解密成功，将会进入这里,获取解密了的token
	token := ctx.Values().Get("jwt").(*jwt.Token)
	//或者这样
	//userMsg :=ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	// userMsg["id"].(float64) == 1
	// userMsg["nick_name"].(string) == iris
	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	//可以了解一下token的数据结构
	ctx.Writef("%s", token.Signature)
	ctx.Next()
}

func JWTMiddleware() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(Secret), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(context context.Context, e error) {
			context.Redirect("/login")
		},
		//debug 模式
		//Debug: bool
	})
}

func CheckJWTMiddleware(ctx iris.Context) {
	jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(Secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(context context.Context, e error) {
			context.Next()
		},
	})
}

func GetSessionKeyFromJWT(ctx iris.Context) {
	token := ctx.Values().Get("jwt").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	if sessionKey, ok := claims["sessionKey"]; ok {
		ctx.Values().Set("sessionKey", sessionKey)
	}
	ctx.Next()
}

func CreateJWTToken(user *model.User) string {
	// 生成加密串过程
	claims := jwt.MapClaims{
		// 自定义参数
		"userId":   user.UserID,
		"userName": user.UserName,
		// 必须
		"iss": "Iris", // 发行人
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), // 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//  把token已约定的加密方式和加密秘钥加密，当然也可以使用不对称加密
	tokenString, _ := token.SignedString([]byte(Secret))
	//  登录时候，把tokenString返回给客户端，然后需要登录的页面就在header上面附此字符串
	//  eg: header["Authorization"] = "bearer "+tokenString
	//ctx.Header("Authorization", fmt.Sprintf("bearer %s", tokenString))
	//ctx.Next()
	return fmt.Sprintf("bearer %s", tokenString)
}

//
//func HttpProxy(ctx iris.Context) {
//	host := os.Getenv("HOST_ZKT_NOTIFY")
//	HostReverseProxy(ctx.ResponseWriter(), ctx.Request(), &TargetHost{
//		Host:    host,
//		IsHttps: false,
//		CAPath:  "",
//	})
//	ctx.Next()
//}
//
//type TargetHost struct {
//	Host    string
//	IsHttps bool
//	CAPath  string
//}
//
//func HostReverseProxy(w http.ResponseWriter, req *http.Request, targetHost *TargetHost) {
//	host := ""
//	if targetHost.IsHttps {
//		host = host + "https://"
//	} else {
//		host = host + "http://"
//	}
//	remote, err := url.Parse(host + targetHost.Host)
//	if err != nil {
//		fmt.Errorf("err:%s", err)
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	req.URL.Path = strings.TrimPrefix(req.RequestURI,"/forward/zkt")
//	proxy := httputil.NewSingleHostReverseProxy(remote)
//	if targetHost.IsHttps {
//		tls, err := GetVerTLSConfig(targetHost.CAPath)
//		if err != nil {
//			fmt.Errorf("https crt error: %s", err)
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		var pTransport http.RoundTripper = &http.Transport{
//			Dial: func(netw, addr string) (net.Conn, error) {
//				c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(5))
//				if err != nil {
//					return nil, err
//				}
//				return c, nil
//			},
//			ResponseHeaderTimeout: time.Second * time.Duration(5),
//			TLSClientConfig:       tls,
//		}
//		proxy.Transport = pTransport
//	}
//	proxy.ServeHTTP(w, req)
//}
//
//var TlsConfig *tls.Config
//
//func GetVerTLSConfig(CAPath string) (*tls.Config, error) {
//	caData, err := ioutil.ReadFile(CAPath)
//	if err != nil {
//		fmt.Errorf("read wechat ca fail", err)
//		return nil, err
//	}
//	pool := x509.NewCertPool()
//	pool.AppendCertsFromPEM(caData)
//	TlsConfig = &tls.Config{
//		RootCAs: pool,
//	}
//	return TlsConfig, nil
//}
