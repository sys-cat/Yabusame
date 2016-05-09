package Data
import (
  "fmt"
  //"log"
  "../config"
  "gopkg.in/olivere/elastic.v3"
)

type Item struct {
  Title     string `json:"Title"`
  Body      string `json:"Body"`
  Category  int    `json:"Category"`
  Created   string `json:"CreatedAt"`
  Updated   string `json:"UpdatedAt"`
}
type User struct {
  SID       string `json:"SystemId"`
  Device    string `json:"Device"`
  Sub       string `json:"Sub"`
  Created   string `json:"CreatedAt"`
  Updated   string `json:"UpdatedAt"`
}
type Calendar struct {
  Date      string  `json:"Date"`
  Item      string  `json:"ItemId"`
  Users     string  `json:"Users"`
  Device    string  `json:"Device"`
  Created   string  `json:"CreatedAt"`
  Updated   string  `jsno:"UpdatedAt"`
}
type Category struct {
  Name      string  `json:"Name"`
  Created   string  `json:"CreatedAt"`
  Updated   string  `json:"UpdatedAt"`
}

var con Config.Base = Config.Get()
var Index string = "Push"
func Client()(client *elastic.Client) {
  client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("%s:%s", con.Elasticsearch.URL, con.Elasticsearch.BASE_PORT)))
  if err != nil {panic(err)}
  return client
}

func Alive()(version string){
  client := Client()
  ver, err := client.ElasticsearchVersion(fmt.Sprintf("%s:%s", con.Elasticsearch.URL, con.Elasticsearch.BASE_PORT))
  if err != nil { panic(err) }
  return ver
}
func ExistsIndex(Name string)(result bool) {
  client := Client()
  exists, err := client.IndexExists(Name).Do()
  if err != nil {panic(err)}
  return exists
}
func CreateIndex(Index string)(result bool) {
  return true
}
