package assignment

import (
	"net/http"
	"strconv"

	board "backend/db/sqlc/out"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AssignmentHandler representa os handlers para assignments
type AssignmentHandler struct {
	assignmentService *AssignmentService
}

// NewAssignmentHandler cria um novo handler para assignments
func NewAssignmentHandler(assignmentService *AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{assignmentService: assignmentService}
}

// 📌 Criar um assignment
func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	var req CreateAssignmentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, 400, "Invalid JSON format", err)
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, 400, "Validation failed", err)
		return
	}

	mappedReq := ToCreateAssignmentParams(req)
	assignment, err := h.assignmentService.CreateAssignment(c, mappedReq)
	if err != nil {
		utils.SendError(c, 500, "Failed to create assignment", err)
		return
	}

	utils.SendSuccess(c, "Assignment created successfully", ToAssignmentResponse(assignment))
}

// 📌 Buscar assignments por reunião
func (h *AssignmentHandler) GetAssignmentsByMeeting(c *gin.Context) {
	meetingID, err := strconv.ParseInt(c.Param("meeting_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meeting ID"})
		return
	}

	assignments, err := h.assignmentService.GetAssignmentsByMeeting(c, meetingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assignments", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

// 📌 Buscar assignments por subsessão
func (h *AssignmentHandler) GetAssignmentsBySubsession(c *gin.Context) {
	subsessionID, err := strconv.ParseInt(c.Param("subsession_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subsession ID"})
		return
	}

	assignments, err := h.assignmentService.GetAssignmentsBySubsession(c, subsessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assignments", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

// 📌 Buscar assignments por usuário
func (h *AssignmentHandler) GetAssignmentsByUserID(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	assignments, err := h.assignmentService.GetAssignmentsByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch assignments", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignments)
}

// 📌 Buscar assignments com paginação
func (h *AssignmentHandler) GetAssignmentsPaginated(c *gin.Context) {
	page, pageSize := utils.GetPaginationParams(c)
	offset := (page - 1) * pageSize

	assignments, total, err := h.assignmentService.GetAssignmentsPaginated(c, pageSize, offset)
	if err != nil {
		utils.SendError(c, 500, "Failed to fetch assignments", err)
		return
	}

	pagination := utils.BuildPagination(page, pageSize, total)
	response := ToAssignmentResponseList(assignments)

	utils.SendPaginatedSuccess(c, "Assignments fetched successfully", response, pagination)
}

// 📌 Atualizar um assignment
func (h *AssignmentHandler) UpdateAssignment(c *gin.Context) {
	var req board.UpdateAssignmentParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.assignmentService.UpdateAssignment(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update assignment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment updated successfully"})
}

// 📌 Deletar um assignment
func (h *AssignmentHandler) DeleteAssignment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.assignmentService.DeleteAssignment(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete assignment", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted successfully"})
}
