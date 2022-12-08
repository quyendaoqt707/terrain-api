package routes

import (
	"TerraInnAPI/controller"
	"TerraInnAPI/middleware"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	// Welcome
	// app.Get("/", controller.Welcome)

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
	motelGroup.Post("", controller.CreateMotelGroup)
	motelGroup.Delete(":id<int>", controller.DelMotelGroup)
	motelGroup.Put("", controller.UpdateMotelGroup)

	// //Setting Flow Hyouka (new version):
	// // user.Post("setting/flow_hyouka/createFlow", controller.CreateFlow_Hyouka)
	// // user.Delete("setting/flow_hyouka/deleteFlow", controller.DeleteFlow_Hyouka)
	// // user.Put("setting/flow_hyouka/updateFlow", controller.UpdateFlow_Hyouka)
	// // user.Get("setting/flow_hyouka/getFlowList", controller.GetFlowListOrDetail_Hyouka)

	// //Setting General: Review Type
	// user.Post("setting/general/createReviewType", controller.CreateReviewType)
	// user.Get("setting/general/getReviewTypeList", controller.GetReviewType)
	// user.Put("setting/general/updateReviewType", controller.UpdateReviewType)
	// user.Delete("setting/general/deleteReviewType", controller.DeleteReviewType)

	// //Setting General: Review Element
	// user.Get("setting/general/getReviewElement", controller.GetReviewElement)
	// user.Post("setting/general/createReviewElement", controller.CreateReviewElement)
	// user.Put("setting/general/updateReviewElement", controller.UpdateReviewElement)
	// user.Delete("setting/general/deleteReviewElement", controller.DeleteReviewElement)

	// //Setting General: Slary Rank

	// user.Get("setting/general/getSalaryRank", controller.GetSalaryRank)
	// user.Post("setting/general/createSalaryRank", controller.CreateSalaryRank)
	// user.Put("setting/general/updateSalaryRank", controller.UpdateSalaryRank)
	// user.Delete("setting/general/deleteSalaryRank", controller.DeleteSalaryRank)

	// //Setting Criteria
	// user.Get("setting/getReviewCriteriaList", controller.GetReviewCriteriaList)
	// user.Post("setting/createReviewCriteria", controller.CreateReviewCriteria)
	// user.Put("setting/updateReviewCriteria", controller.UpdateReviewCriteria)
	// user.Delete("setting/deleteReviewCriteria", controller.DeleteReviewCriteria)

	// //Seting Role/Position
	// user.Get("setting/getReviewRoleList", controller.GetReviewRoleList)
	// user.Post("setting/createReviewRole", controller.CreateReviewRole)
	// user.Put("setting/updateReviewRole", controller.UpdateReviewRole)
	// user.Delete("setting/deleteReviewRole", controller.DeleteReviewRole)

	// //Setting Proportion
	// user.Get("setting/getProportion", controller.GetProportion)
	// user.Post("setting/createProportion", controller.CreateProportion)
	// user.Put("setting/updateProportion", controller.UpdateProportion)
	// user.Delete("setting/deleteProportion", controller.DeleteProportion)

	// //Setting Rating Criteria
	// user.Get("setting/getRatingCriteria", controller.GetRatingCriteria)
	// user.Post("setting/createRatingCriteria", controller.CreateRatingCriteria)
	// user.Put("setting/updateRatingCriteria", controller.UpdateRatingCriteria)
	// user.Delete("setting/deleteRatingCriteria", controller.DeleteRatingCriteria)

	// //Employee Salary Rank
	// user.Get("setting/getEmployeeSalaryRank", controller.GetEmployeeSalaryRank)
	// user.Post("setting/createEmployeeSalaryRank", controller.CreateEmployeeSalaryRank)
	// user.Put("setting/updateEmployeeSalaryRank", controller.UpdateEmployeeSalaryRank)
	// user.Delete("setting/deleteEmployeeSalaryRank", controller.DeleteEmployeeSalaryRank)
	// // user.Post("setting/uploadEmployeeSalaryRank", controller.UploadEmployeeSalaryRank)
	// user.Post("setting/uploadEmployeeSalaryRank", controller.ImportEmployeeSalaryRank)
	// // user.Get("setting/exportEmployeeSalaryRank", controller.ExportEmployeeSalaryRank)
	// user.Get("setting/exportEmployeeSalaryRank", controller.PrepareTemplate)

	// //Helper API
	// user.Get("helper/getPositionList", controller.GetPositionList)
	// user.Get("helper/getSalaryRankList", controller.GetSalaryRankList)
	// user.Get("helper/getEmployeeList", controller.GetEmployeeList)
	// user.Get("helper/prepareDataForFlowSetting", controller.PrepareDataForFlowSetting)
	// user.Post("helper/writeLog", controller.WriteWebLog)

	// //Notification API
	// user.Get("getNotification", controller.GetNotification)

}
