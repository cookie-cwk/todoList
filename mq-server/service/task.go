package service

import (
	"encoding/json"
	"log"
	"mq-server/model"
)
func CreateTask() {
	ch, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	ch.Qos(1, 0, false)
	msgs, err := ch.Consume(q.Name, "", false,false, false, false, nil)
	go func() {
		for d := range msgs {
			var t model.Task
			err := json.Unmarshal(d.Body, &t)
			if err != nil {
				panic(err)
			}
			model.DB.Create(&t)
			log.Println("Done")
			_ = d.Ack(false)
		}
	}()
}
