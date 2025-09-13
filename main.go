package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json: "status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
type TodoItemCreation struct{
	Id int `json:"-"  gorn:"column:id;"`
	Title string `json:"title" gorn:"column:title;"`
	Description string `json: "description" gorn:"column:description;"`
	Status string `json:"status" gorn:"column:status;"`
}

type TodoItemUpdate struct{
	Id int `json:"-"  gorn:"column:id;"`
	Title string `json:"title" gorn:"column:title;"`
	Description string `json: "description" gorn:"column:description;"`
	Status string `json:"status" gorn:"column:status;"`
}

func (TodoItemCreation) TableName() string {return "todo_items"}
func (TodoItemUpdate) TableName() string {return "todo_items"}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(db)

	now := time.Now().UTC()
	item := TodoItem{1,"hgege","đây là test","Doing",&now, nil,}

	jsonData, err :=json.Marshal(item)

	if err != nil {
		fmt.Print(err)
		return
	}else{
		fmt.Println(string(jsonData))
	}

	var items2 TodoItem

	json.Unmarshal(jsonData,&items2)
	fmt.Println(items2)

	r := gin.Default()
	//CRUD
	// POST /v1/items (create new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items:id (get item detail by id)
	// (PUT || PATCH) /v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delete item by id)

	v1:= r.Group("/v1")
	{
		items:= v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id",GetItem(db))
			items.PATCH("/:id")
			items.DELETE("/:id")

		}

	}


	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":item,
		})
	})
	r.Run() //chạy trên port default 8080
}
func CreateItem(db *gorm.DB) func(*gin.Context){
	return func(c *gin.Context){
		var data TodoItemCreation
		//check xem data có phải dạng json
		if err:= c.ShouldBind(&data);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error" :err.Error(),
		})

			return 
		}
		//gorn insert data vào db
		if err := db.Create(&data).Error; err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error" :err.Error(),
			} )

			return
		}
		// lấy data trả về 
		c.JSON(http.StatusOK, gin.H{
		
			"data":data.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var data TodoItem
		
		// /v1/items/:id
		id, err :=strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusAccepted, gin.H{
				"error":err.Error(),
			})
			return 

		}
		if err := db.Where("id =?",id).First(&data).Error; err != nil{
			ctx.JSON(http.StatusAccepted, gin.H{
				"error": err.Error(),
			})

			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}

func UpdateItem(db *gorm.DB) func(*gin.Context){
	return func(ctx *gin.Context) {
		var data TodoItem
		
		id, err :=strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusAccepted, gin.H{
				"error":err.Error(),
			})
			return 

		}
		if err := db.Where("id =?",id).First(&data).Error; err != nil{
			ctx.JSON(http.StatusAccepted, gin.H{
				"error": err.Error(),
			})

			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	}
}