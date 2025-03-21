package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

// 💬 Resposta de sucesso genérica
type SuccessResponse[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// 💬 Resposta de sucesso paginada
type SuccessPaginatedResponse[T any] struct {
	Message    string     `json:"message"`
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// ❌ Resposta de erro
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// 📄 Estrutura de paginação
type Pagination struct {
	Page       int32 `json:"page"`
	PageSize   int32 `json:"pageSize"`
	TotalItems int32 `json:"totalItems"`
	TotalPages int32 `json:"totalPages"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
}

// 🧠 Função para calcular paginação
func BuildPagination(page, pageSize, totalItems int32) Pagination {
	if pageSize <= 0 {
		pageSize = 10
	}
	if page <= 0 {
		page = 1
	}
	totalPages := (totalItems + pageSize - 1) / pageSize

	return Pagination{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}

// 📤 Envia resposta de sucesso simples
func SendSuccess[T any](c *gin.Context, message string, data T) {
	c.JSON(http.StatusOK, SuccessResponse[T]{
		Message: message,
		Data:    data,
	})
}

// 📤 Envia resposta de sucesso com paginação
func SendPaginatedSuccess[T any](c *gin.Context, message string, data []T, pagination Pagination) {
	c.JSON(http.StatusOK, SuccessPaginatedResponse[T]{
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// ❌ Envia resposta de erro padronizada
func SendError(c *gin.Context, status int, message string, err error) {
	c.JSON(status, ErrorResponse{
		Message: message,
		Error:   err.Error(),
	})
}

// GetPaginationParams retorna os valores de page e pageSize com fallback e validação
func GetPaginationParams(c *gin.Context) (page int32, pageSize int32) {
	// Pega page e pageSize da query string
	pageStr := c.DefaultQuery("page", strconv.Itoa(DefaultPage))
	pageSizeStr := c.DefaultQuery("pageSize", strconv.Itoa(DefaultPageSize))

	// Converte para int
	pageInt, err := strconv.Atoi(pageStr)
	if err != nil || pageInt <= 0 {
		pageInt = DefaultPage
	}

	pageSizeInt, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSizeInt <= 0 {
		pageSizeInt = DefaultPageSize
	} else if pageSizeInt > MaxPageSize {
		pageSizeInt = MaxPageSize // previne requisições pesadas
	}

	return int32(pageInt), int32(pageSizeInt)
}
