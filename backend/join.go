
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type JoinRequest struct {
    Tables    []string `json:"tables"`
    Condition string   `json:"condition"`
}

func generateJoinQuery(c *gin.Context) {
    var req JoinRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid join parameters"})
        return
    }

    if len(req.Tables) < 2 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "At least two tables required"})
        return
    }

    query := "SELECT * FROM " + req.Tables[0]
    for i := 1; i < len(req.Tables); i++ {
        query += " JOIN " + req.Tables[i] + " ON " + req.Condition
    }

    c.JSON(http.StatusOK, gin.H{
        "join_query": query,
    })
}
