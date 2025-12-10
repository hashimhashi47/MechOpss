package middleware

import (
	"MechOpss/infra/db"
	"MechOpss/internal/src/models"
	"MechOpss/internal/src/utils"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := []byte(os.Getenv("DB_SECRET_KEY"))

		Token, err := c.Cookie("admin_id")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Cookie expired or Missing"})
			c.Abort()
			return
		}

		claim := &utils.Claims{}
		if tokens, err := jwt.ParseWithClaims(Token, claim, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		}); err == nil && tokens.Valid {
			c.Set("admin_id", claim.UserId)
			c.Next()
			return
		}

		Admin_id := claim.UserId
		var Admin models.Admin

		if err := db.DB.Select("refresh_token").Where("id = ? ", Admin_id).First(&Admin).Error; err != nil {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to find Refershtoken"})
			c.Abort()
			return
		}

		Refershclaim := &utils.Claims{}
		if Refersh, err := jwt.ParseWithClaims(Admin.RefreshToken, Refershclaim, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		}); err != nil || !Refersh.Valid {
			c.Redirect(http.StatusSeeOther, "/admin/login")
			c.JSON(http.StatusBadRequest, gin.H{"Error": "RefershToken expired or Missing"})
			c.Abort()
			return
		}

		NewAccessToken, err := utils.AccessToken(Refershclaim.UserId, Refershclaim.Email, Refershclaim.Role)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Unable to Create refersh Token"})
			c.Abort()
			return
		}

		c.SetCookie("admin_id", NewAccessToken, 7*24*3600, "/", "localhost", false, true)
		c.Set("admin_id", Refershclaim.UserId)
		c.Next()

	}
}