package handlers

import (
	"api-gateway/pkg/utils"
	services "api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)
func GetTaskList(c *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))

	taskService := c.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskResp, err := taskService.GetTasksList(context.Background(),&taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	c.JSON(200, gin.H{
		"data": gin.H{
			"task": taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

func CreateTask(c *gin.Context)  {
	var taskReq services.TaskRequest
	log.Println("CreateTask")
	PanicIfTaskError(c.Bind(&taskReq))
	log.Printf("%v", taskReq)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	log.Printf("claim:%v", c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskService := c.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": taskRes.TaskDetail})
}
func GetTaskDetail(c *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(c.BindUri(&taskReq))
	//从gin.keys取出服务实例
	claim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id,_ := strconv.Atoi(c.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	productService := c.Keys["taskService"].(services.TaskService)
	productRes, err := productService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": productRes.TaskDetail})
}

func UpdateTask(c *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	id,_ := strconv.Atoi(c.Param("id"))
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claim.Id)
	taskService := c.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": taskRes.TaskDetail})
}
func DeleteTask(c *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id,_ := strconv.Atoi(c.Param("id"))
	taskReq.Id = uint64(id)
	taskService := c.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": taskRes.TaskDetail})
}



