package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if student.Name == "" {
		respondWithError(c, http.StatusBadRequest, "Name is required")
		return
	}
	if student.GPA < 0 || student.GPA > 4 {
		respondWithError(c, http.StatusBadRequest, "GPA must be between 0.00 and 4.00")
		return
	}

	if err := h.Service.UpdateStudent(id, student); err != nil {
		respondWithError(c, http.StatusNotFound, "Student not found")
		return
	}

	student.Id = id
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteStudent(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, models.ErrorResponse{
		Error: message,
	})
}
