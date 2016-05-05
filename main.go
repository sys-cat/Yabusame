package main
import (
  "os"
  //"fmt"
  //"log"
  "./config"
  "github.com/gin-gonic/gin"
)

// Request Parameters
type Item struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  Category int `json:"category"`
}
type User struct {
  Id int `json:"id"`
  SysId int `json:"system_id"`
  Device string `json:"device"`
}
type Calender struct {
  Id int `json:"id"`
  Date string `json:"date"`
  ItemId int `json:"item_id"`
  Users string `json:"users"`
  Device string `json:"device"`
}
type Category struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

var con Config.Base = Config.Get()
func main() {
  router := gin.Default()
  var path string
  if con.Serv.ENV == "production" {
    path = "./production.log"
  } else {
    path = "./staging.log"
  }
  f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil { panic(err) }
  gin.DefaultWriter = f
  v1 := router.Group("/v1")
  {
    v1.POST("/item", ItemFunc)
    v1.POST("/user", UserFunc)
    v1.POST("/calendar", CalenderFunc)
    v1.POST("/category", CategoryFunc)
  }
  router.GET("/call/status/check.json", func(c *gin.Context){
    c.JSON(200,gin.H{
      "status":"200",
      "result":"OK",
    })
  })
  router.Run(con.Serv.PORT)
}

func ItemFunc(c *gin.Context) {
  var req Item
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
func UserFunc(c *gin.Context) {
  var req User
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
func CalenderFunc(c *gin.Context) {
  var req Calender
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
func CategoryFunc(c *gin.Context) {
  var req Category
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
