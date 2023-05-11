package main

import (
	"encoding/json"
	"net/http"
	"quiz1/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func CreateUser(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	u := new(model.User)
	if err := json.Unmarshal(data, &u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	if err := model.CreateUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func GetUser(c *gin.Context) {
	reslut := make(map[string]interface{})
	sData := make([]interface{}, 0)
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	data, err := model.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": 99999,
		})

		return
	}

	for _, v := range data {
		res := model.User{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		sData = append(sData, res)
	}

	reslut["data"] = sData
	c.JSON(http.StatusOK, reslut)
}

func UpdateUser(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	u := new(model.User)
	if err := json.Unmarshal(data, &u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	if err := model.UpdateUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	if err := model.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": ERROR_CODE,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
