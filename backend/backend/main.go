package main

import (
    "log"
    "net/http"
    "os"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/yourusername/vite-go-app/backend/handlers"
)

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using system env variables")
    }
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    // Set up Gin
    router := gin.Default()
    
    // CORS middleware
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })
    
    // API routes
    api := router.Group("/api")
    {
        api.GET("/status", handlers.StatusHandler)
        api.POST("/process", handlers.ProcessHandler)
        // Add more API endpoints as needed
    }
    
    // Serve frontend in production
    if os.Getenv("ENV") == "production" {
        router.Static("/", "../frontend/dist")
        router.NoRoute(func(c *gin.Context) {
            c.File("../frontend/dist/index.html")
        })
    }
    
    log.Printf("Server starting on port %s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatal("Server failed to start: ", err)
    }
}