package webserver

import (
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/webserver/handler"
	"github.com/melodiez14/meiko/src/webserver/handler/assignment"
	"github.com/melodiez14/meiko/src/webserver/handler/bot"
	"github.com/melodiez14/meiko/src/webserver/handler/course"
	"github.com/melodiez14/meiko/src/webserver/handler/file"
	"github.com/melodiez14/meiko/src/webserver/handler/information"
	"github.com/melodiez14/meiko/src/webserver/handler/place"
	"github.com/melodiez14/meiko/src/webserver/handler/rolegroup"
	"github.com/melodiez14/meiko/src/webserver/handler/user"
)

// Load returns all routing of this server
func loadRouter(r *httprouter.Router) {

	// Home Handler
	r.GET("/", handler.HelloHandler)

	// ========================== User Handler ==========================
	// User section
	r.POST("/api/v1/user/register", auth.OptionalAuthorize(user.SignUpHandler))
	r.POST("/api/v1/user/verify", auth.OptionalAuthorize(user.EmailVerificationHandler))
	r.POST("/api/v1/user/signin", auth.OptionalAuthorize(user.SignInHandler))
	r.POST("/api/v1/user/forgot", auth.OptionalAuthorize(user.ForgotHandler))
	r.POST("/api/v1/user/signout", auth.MustAuthorize(user.SignOutHandler)) // delete
	r.POST("/api/v1/user/profile", auth.MustAuthorize(user.UpdateProfileHandler))
	r.GET("/api/v1/user/profile", auth.MustAuthorize(user.GetProfileHandler))
	r.POST("/api/v1/user/changepassword", auth.MustAuthorize(user.ChangePasswordHandler))

	// Admin section
	r.GET("/api/admin/v1/user", auth.MustAuthorize(user.ReadHandler))
	r.POST("/api/admin/v1/user", auth.MustAuthorize(user.CreateHandler))
	r.GET("/api/admin/v1/user/:id", auth.MustAuthorize(user.DetailHandler))
	r.POST("/api/admin/v1/user/:id", auth.MustAuthorize(user.UpdateHandler))              // patch
	r.POST("/api/admin/v1/user/:id/activate", auth.MustAuthorize(user.ActivationHandler)) // patch
	r.POST("/api/admin/v1/user/:id/delete", auth.MustAuthorize(user.DeleteHandler))       // delete
	// ======================== End User Handler ========================

	// ======================== Rolegroup Handler =======================
	// Admin section
	r.GET("/api/v1/role", auth.OptionalAuthorize(auth.OptionalAuthorize(rolegroup.GetPrivilege)))
	// ====================== End Rolegroup Handler =====================

	// ========================== File Handler ==========================
	// User section
	r.GET("/api/v1/files/:payload/:filename", file.GetFileHandler)
	r.GET("/api/v1/image/:payload", auth.MustAuthorize(file.GetProfileHandler))
	r.POST("/api/v1/image/profile", auth.MustAuthorize(file.UploadProfileImageHandler))
	r.POST("/api/v1/file/assignment", auth.MustAuthorize(file.UploadAssignmentHandler))
	// ======================== End File Handler ========================

	// ========================= Course Handler =========================
	// User section
	r.GET("/api/v1/course", auth.MustAuthorize(course.GetHandler))
	r.GET("/api/v1/course/assistant", auth.MustAuthorize(course.GetAssistantHandler))
	// Admin section
	r.POST("/api/admin/v1/course", auth.MustAuthorize(course.CreateHandler))
	r.GET("/api/admin/v1/course", auth.MustAuthorize(course.ReadHandler))
	r.GET("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.ReadDetailHandler))                      //read
	r.GET("/api/admin/v1/course/:schedule_id/parameter", auth.MustAuthorize(course.ReadScheduleParameterHandler)) //read
	r.POST("/api/admin/v1/course/:schedule_id", auth.MustAuthorize(course.UpdateHandler))                         //patch
	r.POST("/api/admin/v1/course/:schedule_id/delete", auth.MustAuthorize(course.DeleteScheduleHandler))          //delete
	r.GET("/api/admin/v1/list/course/parameter", auth.MustAuthorize(course.ListParameterHandler))
	r.GET("/api/admin/v1/list/course/search", auth.MustAuthorize(course.SearchHandler))
	// ======================== End Course Handler ======================

	// =========================== Bot Handler ==========================
	// User section
	r.GET("/api/v1/bot", auth.MustAuthorize(bot.LoadHistoryHandler))
	r.POST("/api/v1/bot", auth.MustAuthorize(bot.BotHandler))
	// ========================= End Bot Handler ========================

	// ========================= Assignment Handler ========================
	r.POST("/api/admin/v1/assignment/create", auth.MustAuthorize(assignment.CreateHandler))
	r.GET("/api/admin/v1/assignment/:id", auth.MustAuthorize(assignment.DetailHandler))
	r.GET("/api/admin/v1/assignment", auth.MustAuthorize(assignment.GetAllAssignmentHandler))
	r.POST("/api/admin/v1/assignment/update/:id", auth.MustAuthorize(assignment.UpdateHandler))
	r.POST("/api/admin/v1/assignment/delete/:assignment_id", auth.MustAuthorize(assignment.DeleteAssignmentHandler))
	r.GET("/api/admin/v1/assignment/:id/:assignment_id", auth.MustAuthorize(assignment.GetUploadedAssignmentByAdminHandler))
	r.POST("/api/v1/assignment/create", auth.MustAuthorize(assignment.CreateHandlerByUser))
	r.GET("/api/v1/assignment/:schedule_id/:assignment_id", auth.MustAuthorize(assignment.GetUploadedAssignmentByUserHandler))

	// r.GET("/api/v1/assignment/summary", auth.MustAuthorize(assignment.GetSummaryHandler))
	// ========================= End Assignment Handler ========================
	// // Attendance Handler
	// r.GET("/api/v1/attendance/summary", auth.MustAuthorize(attendance.GetSummaryHandler))
	// r.GET("/api/v1/notification", auth.MustAuthorize(notification.GetHandler))

	// ======================= Information Handler ======================
	// User section
	r.GET("/api/v1/information", auth.MustAuthorize(information.GetSummaryHandler))
	// ===================== End Information Handler ====================

	// ========================== Place Handler =========================
	// Public section
	r.GET("/api/v1/place/search", place.SearchHandler)
	// ======================== End Place Handler =======================

	// Catch
	// r.NotFound = http.RedirectHandler("/", http.StatusPermanentRedirect)
	// r.MethodNotAllowed = http.RedirectHandler("/", http.StatusPermanentRedirect)
}
