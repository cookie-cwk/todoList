package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"log"
	"task/model"
	service "task/services"
)

//创建备忘录，生产出备忘录信息，然后放入消息队列中
func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()

	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
	}
	log.Println("rpc Task service")
	log.Printf("%v", req)
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	body, _ := json.Marshal(req)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:     "application/json",
		Body:            body,
	})
	if err != nil {
		err = errors.New("rabbitMQ publish err:" + err.Error())
	}
	return nil
}

func (*TaskService) GetTasksList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	var count int64
	err := model.DB.Offset(int(req.Start)).Limit(int(req.Limit)).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find:" + err.Error())
	}

	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)
	var taskRes []*service.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return nil
}

func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}

func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content
	model.DB.Save(&taskData)
	resp.TaskDetail = BuildTask(taskData)
	return nil
}

func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).Delete(&model.Task{})
	return nil
}