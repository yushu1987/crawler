package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type Router struct{}
type Common struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
func StartServer()  {
	gin.SetMode(gin.ReleaseMode)




	r := gin.Default()
	r.GET("/search", Search)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8000")
}



func Search(q *gin.Context) {
	word := q.Query("word")
	list:=SearchWord(word)
	q.JSON(http.StatusOK, Common{Code: 0, Message: "success", Data: list})
}