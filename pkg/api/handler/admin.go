package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rganes5/maanushi_earth_e-commerce/pkg/auth"
	"github.com/rganes5/maanushi_earth_e-commerce/pkg/domain"
	"github.com/rganes5/maanushi_earth_e-commerce/pkg/support"
	services "github.com/rganes5/maanushi_earth_e-commerce/pkg/usecase/interface"
	utils "github.com/rganes5/maanushi_earth_e-commerce/pkg/utils"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

// type Response struct {
// 	ID   uint   `copier:"must"`
// 	Name string `copier:"must"`
// }

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// Variable declared containing type as Admin which is already initialiazed in domain folder.
var signUp_admin domain.Admin

func (cr *AdminHandler) AdminSignUp(c *gin.Context) {
	//Binding
	if err := c.BindJSON(&signUp_admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Check the email format
	if err := support.Email_validator(signUp_admin.Email); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Check the phone number format
	if err := support.MobileNum_validator(signUp_admin.PhoneNum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Check whether such email already exits
	if _, err := cr.adminUseCase.FindByEmail(c.Request.Context(), signUp_admin.Email); err == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Hash the password and sign up
	signUp_admin.Password, _ = support.HashPassword(signUp_admin.Password)
	if err := cr.adminUseCase.SignUpAdmin(c.Request.Context(), signUp_admin); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Admin Sign up": "Success",
	})

}

// admin login
func (cr *AdminHandler) AdminLogin(c *gin.Context) {
	//Checks whether a jwt token already exists
	_, err := c.Cookie("admin-token")
	if err == nil {
		// c.JSON(http.StatusAlreadyReported, gin.H{
		// 	"admin": "Already logged in",
		// })
		c.Redirect(http.StatusFound, "/admin/home")
		return
	}
	//binding
	var Login_admin utils.LoginBody
	if err := c.BindJSON(&Login_admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check email
	admin, err := cr.adminUseCase.FindByEmail(c.Request.Context(), Login_admin.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check the password
	if err := support.CheckPasswordHash(Login_admin.Password, admin.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Create a jwt token and store it in cookie
	tokenstring, err := auth.GenerateJWT(admin.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetCookie("admin-token", tokenstring, int(time.Now().Add(60*time.Minute).Unix()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Admin": "Success",
	})
}

//admin logout

func (cr *AdminHandler) Logout(c *gin.Context) {
	c.SetCookie("admin-token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Logout": "success",
	})
}

//home handler

func (cr *AdminHandler) HomeHandler(c *gin.Context) {
	email, ok := c.Get(("admin-email"))
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized admin",
		})
		return
	}

	admin, err := cr.adminUseCase.FindByEmail(c.Request.Context(), email.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, admin)
}

//list users

// func (cr *AdminHandler) ListUsers(c *gin.Context) {
// 	users, err := cr.adminUseCase.ListUsers(c.Request.Context())
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"users": users,
// 	})
// }
