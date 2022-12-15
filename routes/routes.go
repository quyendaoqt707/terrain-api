package routes

import (
	"TerraInnAPI/controller"
	"TerraInnAPI/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	// Welcome
	app.Get("/", controller.Wellcome)

	/** Without AppAuthen**/
	api := app.Group("/api")
	api.Post("/login", controller.Login)
	api.Post("/register", controller.InsertUser)
	api.Get("/motel-group", controller.GetModelGroup)
	api.Get("/motel", controller.GetModelDetail)

	user := api.Group("/user", middleware.AppAuthen)
	// user.Put("/updateLanguageAndTheme", controller.UpdateLanguageAndTheme)

	user.Post("/logout", controller.Logout)
	user.Get("", controller.GetUser)
	user.Put("/update", controller.UpdateProfile)
	user.Put("/change-password", controller.ChangePassword)

	// Motel
	motel := api.Group("/motel", middleware.AppAuthen)
	motel.Post("", controller.CreateMotel)
	motel.Delete(":id<int>", controller.DelMotel)
	motel.Put("", controller.UpdateMotel)

	// Motel Group
	motelGroup := api.Group("/motel-group", middleware.AppAuthen)
	motelGroup.Get("/list", controller.GetGroupList)
	motelGroup.Post("", controller.CreateMotelGroup)
	motelGroup.Delete(":id<int>", controller.DelMotelGroup)
	motelGroup.Put("", controller.UpdateMotelGroup)

	// Invoice
	invoice := api.Group("/invoice")
	invoice.Get("", controller.GetInvoice)
	invoice.Post("", controller.CreateInvoice)
	invoice.Delete(":id<int>", controller.DelInvoice)
	invoice.Put("", controller.UpdateInvoice)

	// Request
	request := api.Group("/request", middleware.AppAuthen)
	request.Get("", controller.GetRequest)
	request.Post("", controller.CreateRequest)
	request.Delete(":id<int>", controller.DelRequest)
	request.Put("", controller.UpdateRequest)

	//Helper
	api.Post("/asset", controller.UploadImg)

	// //Helper API
	// user.Get("helper/getPositionList", controller.GetPositionList)
	// user.Get("helper/getSalaryRankList", controller.GetSalaryRankList)
	// user.Get("helper/getEmployeeList", controller.GetEmployeeList)
	// user.Get("helper/prepareDataForFlowSetting", controller.PrepareDataForFlowSetting)
	// user.Post("helper/writeLog", controller.WriteWebLog)

	// //Notification API
	// user.Get("getNotification", controller.GetNotification)

}
