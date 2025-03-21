package routes

import (
	"backend/pkg/assignment"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas da API
func SetupRouter(r *gin.Engine, assignmentService *assignment.AssignmentService) {
	assignmentHandler := assignment.NewAssignmentHandler(assignmentService)

	api := r.Group("/api")
	{
		assignments := api.Group("/assignments")
		{
			assignments.POST("", assignmentHandler.CreateAssignment)
			assignments.GET("/meeting/:meeting_id", assignmentHandler.GetAssignmentsByMeeting)
			assignments.GET("/subsession/:subsession_id", assignmentHandler.GetAssignmentsBySubsession)
			assignments.GET("/user/:user_id", assignmentHandler.GetAssignmentsByUserID)
			assignments.GET("", assignmentHandler.GetAssignmentsPaginated)
			assignments.PUT("", assignmentHandler.UpdateAssignment)
			assignments.DELETE("/:id", assignmentHandler.DeleteAssignment)
		}
	}

}
