package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	repository2 "project1/src/infra/repository"
	usecase2 "project1/src/usecase"

	kafka "project1/src/infra/kafka"

	_ "github.com/go-sql-driver/mysql"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "app:app@tcp(mysql:3306)/app")
	if err != nil {
		log.Fatalln(err)
	}
	repository := repository2.CourseMySQLRepository{Db: db}
	usecase := usecase2.CreateCourse{CourseRepository: &repository}

	var message = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
		"group.id":          "app_go",
	}

	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)
	go consumer.Consume(message)

	for msg := range message {
		var input usecase2.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)

		if err != nil {
			fmt.Println("error: ", err)
		} else {
			fmt.Println("output: ", output)
		}
	}
}
