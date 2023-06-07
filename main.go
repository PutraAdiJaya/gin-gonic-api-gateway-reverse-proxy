package main

import (
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")


		c.Next()
	}
}
const CoreUrl = "http://localhost:8082"
const TranscUrl = "http://localhost:8083"
const AuthUrl = "http://localhost:8084"
const NotifUrl = "http://localhost:8085"

var AuthService  = []string{
	"/v1/auth/login",
	"/v1/auth/refresh",
	"/v1/auth/logout",
	"/v1/auth/backoffice/login",
	"/v1/auth/register-email/before",
	"/v1/auth/register/confirm",
	"/v1/auth/register",
}

var NotifService  = []string{
	"/v1/devices",
	"/v1/notifications",
	"/v1/notifications/summary",
	"/v1/notif-config",
	"/v1/me/count/notifications",
}

var CoreService = []string{
	"/v1/profile-basic",
	"/v1/product-favorites",
	"/v1/profile-address",
	"/v1/profile-address/:id",
	"/v1/categories",
	"/v1/cities",
	"/v1/banners",
	"/v1/facebook-groups",
	"/v1/products-homepages",
	"/v1/profile-personal-info",
	"/v1/profile-config-delivery",
	"/v1/profile-config-payment",
	"/v1/profile-general-info",
	"/v1/profile-identity",
	"/v1/profile-banks",
	"/v1/profile-banks/:id",
	"/v1/me/count/product-favorites",
	"/v1/search-info",
	"/v1/search-suggestions",
	"/v1/products",
	"/v1/products/:id",
	"/v1/internal/products",
}



var TranscService = []string{
	"/v1/checkout",
	"/v1/payment/listener",
	"/v1/my-sales",
	"/v1/my-sales/:id",
	"/v1/my-purchases",
	"/v1/my-purchases/:id",
	"/v1/cart",
	"/v1/cart/count",
	"/v1/cart/:id",
	"/v1/cart/checkout/:id",
	"/v1/backoffice/withdraws",
	"/v1/backoffice/withdraws/:id",
	"/v1/backoffice/transactionsv2",
	"/v1/backoffice/transaction-filters",
	"/v1/delivery-vendors",
	"/v1/backoffice/transactions/summary",
	"/v1/backoffice/download/transaction",
}



var DocsService = []string{
	"/v1/documents",
}




func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(cors.Default())
	r.Use(CORSMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	for i := 0; i < len(AuthService ); i++ {
		r.OPTIONS(AuthService[i], ReverseProxy(AuthUrl))
		r.POST(AuthService[i], ReverseProxy(AuthUrl))
		r.GET(AuthService[i], ReverseProxy(AuthUrl))
	}

	for i := 0; i < len(CoreService ); i++ {
		r.OPTIONS(CoreService[i], ReverseProxy(CoreUrl))
		r.POST(CoreService[i], ReverseProxy(CoreUrl))
		r.GET(CoreService[i], ReverseProxy(CoreUrl))
	}
	for i := 0; i < len(TranscService ); i++ {
		r.OPTIONS(TranscService[i], ReverseProxy(TranscUrl))
		r.POST(TranscService[i], ReverseProxy(TranscUrl))
		r.PUT(TranscService[i], ReverseProxy(TranscUrl))
		r.GET(TranscService[i], ReverseProxy(TranscUrl))
		r.DELETE(TranscService[i], ReverseProxy(TranscUrl))
	}
	for i := 0; i < len(NotifService ); i++ {
		r.OPTIONS(NotifService[i], ReverseProxy(NotifUrl))
		r.POST(NotifService[i], ReverseProxy(NotifUrl))
		r.GET(NotifService[i], ReverseProxy(NotifUrl))
	}
	for i := 0; i < len(DocsService ); i++ {
		r.POST(DocsService[i], ReverseProxy(NotifUrl))
		r.GET(DocsService[i], ReverseProxy(NotifUrl))
	}


	r.Run(":9000")
}

func ReverseProxy(target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	checkErr(err)
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func checkErr(err error) {

}
