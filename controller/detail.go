package controller

import (
	"firstapp/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type DetailJSON struct {
	Category string `json:"category"`
	Cost     int    `json:"cost"`
	Date     string `json:"date"`
}

func GetDetails(c *gin.Context) {
	claims, _ := c.Get("claims")
	user, err := model.GetUserByName(claims.(jwt.MapClaims)["username"].(string))
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	details, err := model.GetAllDetailsOfUser(user.ID)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, details)
	}
}

func CreateDetail(c *gin.Context) {
	var detailJSON DetailJSON
	if err := c.BindJSON(&detailJSON); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(detailJSON)

	claims, _ := c.Get("claims")
	user, err := model.GetUserByName(claims.(jwt.MapClaims)["username"].(string))
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	category, err := model.GetCategoryByName(detailJSON.Category)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "category not found"})
		return
	}

	detail := model.Detail{
		UserID:     user.ID,
		CategoryID: category.ID,
		Cost:       detailJSON.Cost,
		Date:       detailJSON.Date,
	}
	createdDetail, err := detail.Save()
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": "unable to create detail"})
	} else {
		c.JSON(201, createdDetail)
	}
}

func GetDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	detail, err := model.GetDetail(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, detail)
	}
}

func DeleteDetail(c *gin.Context) {
	id := c.Params.ByName("id")
	err := model.DeleteDetail(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, gin.H{"message": "detail deleted"})
	}
}