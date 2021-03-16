package service

import (
	"github.com/SmartDuck9000/go-tradex/src/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api StatService) saveStat(c *gin.Context) {
	var stat db.SavedStat
	err := c.BindJSON(&stat)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = api.db.CreateStat(&stat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.Status(http.StatusOK)
	}
}

func (api StatService) getStat(c *gin.Context) {
	fromDate := c.Query("from")
	if fromDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't find from field in query parameters",
		})
	}

	toDate := c.Query("to")
	if toDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't find to field in query parameters",
		})
	}

	orderBy := c.Query("ordered_by")
	if orderBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Can't find order_by field in query parameters",
		})
	}

	stat := api.db.GetStat(fromDate, toDate, orderBy)
	c.JSON(http.StatusOK, stat)
}

func (api StatService) deleteStat(c *gin.Context) {
	api.db.DeleteStat()
	c.Status(http.StatusOK)
}
