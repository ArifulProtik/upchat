package controller

import (
	"net/http"
	"strings"

	"ArifulProtik/UpChat/internal/controller/middleware"
	"ArifulProtik/UpChat/internal/data"
	"ArifulProtik/UpChat/internal/ent"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// SignUp creates a new user account.
//
//	@Summary	SignUp
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		body	body		data.UserCreateBody	true	"User registration data"
//	@Success	200		{object}	data.UserResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/auth/signup [post]
func (c *Controller) SignUp(ctx *gin.Context) {
	var body data.UserCreateBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  data.FormatValidationErrors(err, body),
		})
		return
	}

	user, err := c.service.CreateUser(ctx.Request.Context(), body)
	if err != nil {
		if ent.IsConstraintError(err) {
			ctx.JSON(http.StatusBadRequest, &ErrorResponse{
				Status: http.StatusBadRequest,
				Error:  "User already exists.",
			})
			return
		}
		c.logger.Error("[Auth_Controller]", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  "Something went wrong.",
		})
		return
	}

	var response data.UserResponse
	if err := copier.Copy(&response, user); err != nil {
		c.logger.Error("copier error", "error", err.Error())
	}

	ctx.JSON(http.StatusOK, response)
}

// LogIn authenticates a user and returns a session token.
//
//	@Summary	Log in
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		body	body		data.LoginBody	true	"Login credentials"
//	@Success	200		{object}	data.LoginResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/auth/login [post]
func (c *Controller) LogIn(ctx *gin.Context) {
	var body data.LoginBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  gin.H{"errors": data.FormatValidationErrors(err, body)},
		})
		return
	}
	account, err := c.service.FindAccountByEmail(ctx.Request.Context(), body.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &ErrorResponse{
			Status: http.StatusUnauthorized,
			Error:  "Wrong email or password.",
		})
		return
	}
	err = c.service.VerifyPassword(body.Password, account.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &ErrorResponse{
			Status: http.StatusUnauthorized,
			Error:  "Wrong email or password.",
		})
		return
	}
	loginData, err := c.service.Login(ctx.Request.Context(), account, body)
	if err != nil {
		c.logger.Error("[Auth_Controller]", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  "Something went wrong",
		})
		return
	}
	var loginResponse data.LoginResponse
	if err := copier.Copy(&loginResponse, loginData); err != nil {
		c.logger.Error("copier error", "error", err.Error())
	}
	ctx.JSON(http.StatusOK, loginResponse)
}

// LogOut invalidates the current session token.
//
//	@Summary	Log out
//	@Tags		auth
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	400	{object}	ErrorResponse
//	@Security	BearerAuth
//	@Router		/auth/logout [post]
func (c *Controller) LogOut(ctx *gin.Context) {
	token_header := ctx.GetHeader("Authorization")
	token := strings.Split(token_header, " ")[1]
	if token_header == "" {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "Token is required",
		})
		return
	}
	err := c.service.DeleteSession(ctx.Request.Context(), token)
	if err != nil {
		c.logger.Error("[Auth_Controller]", "error", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// GetSession returns the currently authenticated user.
//
//	@Summary	Get current session
//	@Tags		auth
//	@Produce	json
//	@Success	200	{object}	data.GetSessionResponse
//	@Failure	500	{object}	ErrorResponse
//	@Security	BearerAuth
//	@Router		/auth/get-session [get]
func (c *Controller) GetSession(ctx *gin.Context) {
	user_id := ctx.MustGet(middleware.UserIDKey).(string) //nolint:errcheck
	c.logger.Info(user_id)
	userdata, err := c.service.FindAccountByID(ctx.Request.Context(), user_id)
	if err != nil {
		c.logger.Error("[Auth_Controller]", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  "Something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, &data.GetSessionResponse{
		User:    Copy[data.UserResponse](userdata.Edges.User),
		Account: Copy[data.AccountResponse](userdata),
	})
}
