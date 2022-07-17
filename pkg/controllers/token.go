package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotipath/pkg/token"
)

type Request struct {
	Uri string `json:"uri"`
}

func GetToken(c *gin.Context) {

	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uri": token.GenerateFinalUrl(request.Uri),
	})
}

//2022-17-07-060850
//2022-17-07:06:08:50
