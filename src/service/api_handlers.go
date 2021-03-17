package service

import (
	"github.com/SmartDuck9000/go-tradex/src/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (api StatService) saveStat(c *gin.Context) {
	var stat data.SavedStat
	err := c.BindJSON(&stat)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.validator.ValidateStruct(stat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.db.CreateStat(&stat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusOK)
	}
}

func (api StatService) getStat(c *gin.Context) {
	var filterData data.FilterData
	err := c.BindQuery(&filterData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = api.validator.ValidateStruct(filterData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stat, getErr := api.db.GetStat(filterData.From, filterData.To, filterData.OrderBy)
	if getErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getErr.Error()})
	} else {
		c.JSON(http.StatusOK, stat)
	}
}

func (api StatService) deleteStat(c *gin.Context) {
	err := api.db.DeleteStat()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}
