package controller

import (
	"database/sql"
	"go-contacts-api/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	if _, err := h.db.Query("INSERT INTO users (email, password) VALUES ($1, $2) ", user.Email, string(hashedPassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Sign Up successfully"})

}

func (h *Handler) SignIn(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.QueryRow("SELECT password FROM users where email=$1", user.Email)

	stored := models.User{}

	err := result.Scan(&stored.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("id", user.Id)
	session.Set("email", user.Email)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "User Sign In successfully"})

}

func (h *Handler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "User Sign out successfully"})
}
