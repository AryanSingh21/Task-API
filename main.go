package main

import (
	connect "github.com/AryanSingh21/Task-API/database"
	"github.com/gin-gonic/gin"
)


func init()  {
	connect.ConnectDatabse()
}

func deleteTask(ctx *gin.Context){
	
	id := ctx.Param("id")

	tasks := []connect.Task{}
	result := connect.DB.Delete(&tasks , id)


	//if we do not find the id
	if result.RowsAffected == 0 {
		ctx.JSON(400 , gin.H{
			"message" : "Id not found",
		})
		return
	}
	//if we find the error and delete it successfully
	ctx.JSON(200 , gin.H{
		"Success" : "Task deleted successfully",
	})
}
func singleTask(ctx *gin.Context){
	// getting id from url
	id := ctx.Param("id")

	task := connect.Task{}
	result := connect.DB.First(&task , id)

	if result.Error != nil {
		ctx.JSON(400 , gin.H{
			"message" : "Error while fetching task",
		})
		return
	}
	ctx.JSON(200 , gin.H{
		"Task" : task,
	})
}
func allTask(ctx *gin.Context){
	task := []connect.Task{}
	result := connect.DB.Find(&task)

	if result.Error != nil {
		ctx.JSON(400 , gin.H{
			"message" : "Error while fetching tasks",
		})
		return
	}
	ctx.JSON(200 , gin.H{
		"Tasks" : task,
	})
}
func CreateTask(ctx *gin.Context)  {

	var body struct {
	ID          uint 	 
	Title       string  
	Description string
	Due_Date    string
	Status string

	}
	ctx.Bind(&body)
	if body.Title == "" || body.Description == "" || body.Due_Date == "" || body.Status == "" {
		ctx.JSON(400 , gin.H{
			"error" : "missing fields",
		})
		return
	} 
	
	

	
	
	task := connect.Task{ID: body.ID , Title: body.Title , Description: body.Description , Due_Date: body.Due_Date , Status: body.Status}

	result := connect.DB.Create(&task)

	if result.Error != nil {
		ctx.Status(400)
		return
	}
	ctx.JSON(200 ,  gin.H{
			"task" : task,
		},
	)
}

func updateTask(ctx *gin.Context) {
	//get id
	id := ctx.Param("id")
	//get data from body
	var body struct {
		ID          uint
		Title       string  
		Description string
		Due_Date    string
		Status string 		
	
	}
	ctx.Bind(&body)
	//find the post
	task := connect.Task{}
	result := connect.DB.First(&task , id)
	if result.Error != nil {
		ctx.JSON(400 , gin.H{
			"message" : "Error while upadting",
		})
		return
	}
	//update the post
	connect.DB.Model(&task).Updates(connect.Task{ID: body.ID , Title: body.Title , Description: body.Description , Due_Date: body.Due_Date , Status: body.Status})
	//response

	ctx.JSON(200 , gin.H{
		"Task" : task,
	})


}

func main() {
	router := gin.Default()

	router.POST("/" , CreateTask)
	router.GET("/tasks" , allTask)
	router.GET("/tasks/:id" , singleTask)
	router.DELETE("/tasks/:id" , deleteTask)
	router.PUT("/tasks/:id" , updateTask)
	router.Run()

}