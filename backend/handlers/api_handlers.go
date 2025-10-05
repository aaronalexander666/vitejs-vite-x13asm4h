package handlers

import (
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
)

// StatusHandler returns server status
func StatusHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
        "message": "Go backend is running!",
        "timestamp": time.Now().Format(time.RFC3339),
    })
}

// ProcessHandler processes data from frontend
func ProcessHandler(c *gin.Context) {
    var request struct {
        Prompt string `json:"prompt" binding:"required"`
    }
    
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request payload",
        })
        return
    }
    
    // Here you would process the request with your business logic
    c.JSON(http.StatusOK, gin.H{
        "result": "Processed: " + request.Prompt,
        "timestamp": time.Now().Format(time.RFC3339),
    })
}