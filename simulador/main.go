// package main

// import (
// 	"fmt"
// 	route2 "github.com/codeedu/imersaofsfc2-simulator/application/route"
// 	"github.com/joho/godotenv"
// 	"log"
// )
// func init(){
// err := godotenv.Load()
// if err != nil{
// 	log.Fatal( v...: "")
// }
// }

// func main(){
// 	route := route2.Route{
// 		ID:				 "1",
// 		ClientID: "1",
// 		Positions: nil,
// 	}

// 	route.LoadPositions()
// 	stringjson, _:= route.ExportJsonPositions()

// fmt.Println(stringjson[1])
// }

package main

import (
	"fmt"
	kafka2 "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}