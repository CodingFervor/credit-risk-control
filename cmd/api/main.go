package main

import (
	"log"
	"github.com/gin-gonic/gin"

	"github.com/CodingFervor/credit-risk-control/internal/database"
)

func main() {
	r := gin.Default()
	r.POST("/api/v1/auth/login", func(c *gin.Context) { c.JSON(200, gin.H{"token": "mock"}) })
	api := r.Group("/api/v1")
	{
		// Loan Applications
		api.POST("/loans/apply", func(c *gin.Context) { c.JSON(201, gin.H{"data": gin.H{"status": "pending", "risk_level": "medium"}}) })
		api.GET("/loans", func(c *gin.Context) { c.JSON(200, gin.H{"data": []interface{}{}}) })
		api.GET("/loans/:id", func(c *gin.Context) { c.JSON(200, gin.H{"data": gin.H{}}) })
		api.PUT("/loans/:id/approve", func(c *gin.Context) { c.JSON(200, gin.H{"message": "approved"}) })
		api.PUT("/loans/:id/reject", func(c *gin.Context) { c.JSON(200, gin.H{"message": "rejected"}) })
		api.POST("/loans/:id/disburse", func(c *gin.Context) { c.JSON(200, gin.H{"message": "disbursed"}) })

		// Credit Accounts
		api.GET("/credit/:user_id", func(c *gin.Context) { c.JSON(200, gin.H{"data": gin.H{"limit": 50000, "used": 12000}}) })
		api.PUT("/credit/:user_id/adjust", func(c *gin.Context) { c.JSON(200, gin.H{"message": "adjusted"}) })
		api.PUT("/credit/:user_id/freeze", func(c *gin.Context) { c.JSON(200, gin.H{"message": "frozen"}) })

		// Risk Rules
		api.POST("/risk/rules", func(c *gin.Context) { c.JSON(201, gin.H{"message": "created"}) })
		api.GET("/risk/rules", func(c *gin.Context) { c.JSON(200, gin.H{"data": []interface{}{}}) })
		api.PUT("/risk/rules/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "updated"}) })
		api.POST("/risk/evaluate", func(c *gin.Context) { c.JSON(200, gin.H{"data": gin.H{"score": 720, "level": "low"}}) })

		// Blacklist
		api.POST("/blacklist", func(c *gin.Context) { c.JSON(201, gin.H{"message": "added"}) })
		api.GET("/blacklist", func(c *gin.Context) { c.JSON(200, gin.H{"data": []interface{}{}}) })
		api.DELETE("/blacklist/:id", func(c *gin.Context) { c.JSON(200, gin.H{"message": "removed"}) })
		api.GET("/blacklist/check", func(c *gin.Context) { c.JSON(200, gin.H{"hit": false}) })

		// Repayment
		api.GET("/repayment/:loan_id", func(c *gin.Context) { c.JSON(200, gin.H{"data": []interface{}{}}) })
		api.POST("/repayment/:loan_id/pay", func(c *gin.Context) { c.JSON(200, gin.H{"message": "paid"}) })

		// Collection
		api.GET("/collection/tasks", func(c *gin.Context) { c.JSON(200, gin.H{"data": []interface{}{}}) })
		api.PUT("/collection/:id/assign", func(c *gin.Context) { c.JSON(200, gin.H{"message": "assigned"}) })
		api.PUT("/collection/:id/result", func(c *gin.Context) { c.JSON(200, gin.H{"message": "updated"}) })

		// Risk Dashboard
		api.GET("/dashboard/overview", func(c *gin.Context) { c.JSON(200, gin.H{"data": gin.H{"total_loans": 0, "overdue_rate": 0, "approval_rate": 0}}) })
		api.GET("/dashboard/risk-events", func(c *gin.Context) { c.JSON(200, gin.H{"data": []interface{}{}}) })
	}
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	log.Println("Credit Risk Control starting on :8080")
	addr := ":" + strconv.Itoa(8080)
	srv := &http.Server{Addr: addr, Handler: r}
	go func() {
		logger.Info("server listening", "port", 8080)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("forced shutdown", "error", err)
	}
	logger.Info("server exited")
}
