package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
    ID       int     `json:"id"`
    Type     string  `json:"type"`     // 收入/支出
    Category string  `json:"category"` // 分类
    Amount   float64 `json:"amount"`   // 金额
}

func main() {
        // 创建默认 Gin 引擎
    r := gin.Default()

    
	// 跨域 (CORS)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})


	// 模拟数据
	transactions := []Transaction{
		{ID: 1, Type: "income", Category: "Salary", Amount: 5000},
		{ID: 2, Type: "expense", Category: "Food", Amount: 120},
		{ID: 3, Type: "expense", Category: "Transport", Amount: 80},
	}

	// GET /api/transactions
	r.GET("/api/transactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, transactions)
	})

	// POST /api/transactions
	r.POST("/api/transactions", func(c *gin.Context) {
		var t Transaction
		if err := c.BindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		t.ID = len(transactions) + 1
		transactions = append(transactions, t)
		c.JSON(http.StatusCreated, t)
	})


	r.Run(":8080") // AWS Elastic Beanstalk 默认端口
}