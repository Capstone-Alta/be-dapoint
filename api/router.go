package api

import (
	"dapoint-api/api/middleware"
	auth "dapoint-api/api/v1/auth"
	contentV1 "dapoint-api/api/v1/content"
	"dapoint-api/api/v1/user"
	voucherController "dapoint-api/api/v1/voucher"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	ContentV1Controller *contentV1.Controller
	UserController      *user.Controller
	AuthController      *auth.Controller
	MiddlewareJwt       middleware.JWTService
	VoucherController   *voucherController.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	contentV1 := e.Group("/v1/content")
	// contentV1.Use(middleware.JWTMiddleware())
	contentV1.GET("", controller.ContentV1Controller.GetAll)
	contentV1.POST("", controller.ContentV1Controller.Create)

	auth := e.Group("/v1/auth")
	auth.POST("", controller.AuthController.Auth)

	user := e.Group("/users")
	user.POST("/login", controller.UserController.Login)                                          // login
	user.GET("", controller.UserController.GetAll, controller.MiddlewareJwt.AdminJwtMiddleware()) // get all users
	user.POST("", controller.UserController.Create)                                               // register users
	user.GET("/:id", controller.UserController.GetByID)                                           // get users by id
	user.PUT("/:id", controller.UserController.Modify)                                            // update users
	// TODO Delete users
	user.DELETE("/:id", controller.UserController.Delete)

	// TODO Create vouchers
	voucher := e.Group("/vouchers")
	voucher.POST("", controller.VoucherController.Create)
	voucher.GET("", controller.VoucherController.GetAll)
	// TODO Update Vouchers by id
	// TODO Delete vouchers
}
