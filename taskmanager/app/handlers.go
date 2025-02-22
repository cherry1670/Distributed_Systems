// Description: Defines HTTP request handlers and encapsulates dependencies for API functionality.

package app

import (
	"net/http"

	"manager/models"

	"github.com/gin-gonic/gin"
)

// Search handles the Papertrail search API call.
func (handler *Handler) CreateTask(c *gin.Context) {
	var task models.Task

	// Bind and validate the JSON request
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := handler.operations.ProcessCreate(task)
	if resp.StatusCode != 200 {
		c.JSON(resp.StatusCode, gin.H{"message": resp.Message})
		return
	}
	c.JSON(resp.StatusCode, resp.Data)

}
