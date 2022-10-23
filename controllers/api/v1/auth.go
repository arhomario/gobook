package v1

import (
	"gobook/auth"
	"gobook/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Postlogin(ctx *gin.Context) {
	var user models.User
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	db := ctx.MustGet("db").(*gorm.DB)

	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User not found!",
		})
		return
	} else if err := models.VerifyPassword(user.Password, password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Password is wrong!",
		})
		return
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Create Token Failed",
		})
		return
	}

	user.SetToken(token)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create Token Success",
		"data":    user,
	})

}
