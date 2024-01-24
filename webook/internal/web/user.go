package web

import (
	"github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	EmailRegexp    *regexp2.Regexp
	PasswordRegexp *regexp2.Regexp
}

func NewUserHandler() *UserHandler {
	emailRegex := regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, regexp2.None)
	passwordRegex := regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`, regexp2.None)
	return &UserHandler{
		EmailRegexp:    emailRegex,
		PasswordRegexp: passwordRegex,
	}

}

func (u *UserHandler) RegisterRouter(ug *gin.RouterGroup) {
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}
func (u *UserHandler) RegisterRouter1(server *gin.Engine) {
	ug := server.Group("/user")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}
func (u *UserHandler) Profile(ctx *gin.Context) {

}
func (u *UserHandler) Signup(ctx *gin.Context) {
	//要大写
	type SignupReq struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var req SignupReq
	//传指针才能修改值
	if err := ctx.Bind(&req); err != nil {
		return
	}
	isEmail, err := u.EmailRegexp.MatchString(req.Name)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
	}
	if !isEmail {
		ctx.String(http.StatusOK, "邮箱错误")
		return
	}
	ctx.String(http.StatusOK, "success")
}
func (u *UserHandler) Login(ctx *gin.Context) {

}
func (u *UserHandler) Edit(ctx *gin.Context) {

}
