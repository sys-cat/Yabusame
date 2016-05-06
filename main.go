package main
import (
  "os"
  //"fmt"
  //"log"
  "./config"
  "./data"
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
type Calendar struct {
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

func setAccessHeader(c *gin.Context) {
  c.Header("Access-Control-Allow-Origin", "*")
  c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
  c.Header("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS")
}

func main() {
  router := gin.Default()
  var path string
  if con.Server.ENV == "production" {
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
    v1.POST("/calendar", CalendarFunc)
    v1.POST("/category", CategoryFunc)
  }
  router.GET("/call/status/check.json", func(c *gin.Context){
    c.JSON(200,gin.H{
      "status":"200",
      "result":"OK",
      "ESVer":Data.Alive(),
      "ESIndex":Data.ExistsIndex("Master"),
    })
  })
  router.Run(con.Server.PORT)
}

func ItemFunc(c *gin.Context) {
  setAccessHeader(c)
  var req Item
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
func UserFunc(c *gin.Context) {
  setAccessHeader(c)
  var req User
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
func CalendarFunc(c *gin.Context) {
  setAccessHeader(c)
  var req Calendar
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
func CategoryFunc(c *gin.Context) {
  setAccessHeader(c)
  var req Category
  c.BindJSON(&req)
  c.JSON(200, gin.H{"status":"200"})
}
