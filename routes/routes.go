package routes

import (
    "github.com/gin-gonic/gin"
    "task-golang-db/handlers"
)

// SetupRoutes initializes all the routes for the application
func SetupRoutes() *gin.Engine {
    router := gin.Default()

    // Transaction Category Routes
    transactionCategoryGroup := router.Group("/transaction-category")
    {
        transactionCategoryGroup.POST("/", handlers.CreateTransactionCategory)
        transactionCategoryGroup.GET("/", handlers.ListTransactionCategories)
        // Add PUT and DELETE handlers here
    }

    // Account Routes
    router.POST("/account/topup", handlers.TopUpAccount)
    router.POST("/account/balance", handlers.GetAccountBalance)
    // Add more account-related routes as needed

    return router
}
