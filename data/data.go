package Data
import (
  "fmt"
  "../config"
  "gopkg.in/olivere/elastic.v3"
)

func Alive(ES Config.Elasticsearch)(version string){
  client, err := elastic.NewClient()
  if err != nil { panic(err) }
  ver, err := client.ElasticsearchVersion(fmt.Sprintf("%s:%s", ES.URL, ES.BASE_PORT))
  if err != nil { panic(err) }
  return ver
}
