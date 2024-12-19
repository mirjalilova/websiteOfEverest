package handlers

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	"github.com/mirjalilova/websiteOfEverest/pkg/utils"
)

type TeacherHandler struct {
	clients *pb.TeacherServiceClient
}

func NewTeacherHandler(clients *pb.TeacherServiceClient) *TeacherHandler {
	return &TeacherHandler{clients: clients}
}

// TeacherCreate godoc
// @Summary Create a teacher
// @Description Create a teacher with details
// @Tags teachers
// @Accept  json
// @Produce  json
// @Param teacher body proto.CreateTeacher true "Create teacher"
// @Success 200 {string} string "Teacher created successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /teachers/create [post]
func (h *Handler) TeacherCreate(c *gin.Context) {
	req := &pb.CreateTeacher{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	if _, err := h.Clients.Teacher.Create(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error creating teacher", "err", err)
		return
	}

	slog.Info("Teacher created successfully")
	c.JSON(200, gin.H{"message": "Teacher created successfully"})
}

// TeacherUpdate godoc
// @Summary Update a teacher
// @Description Update teacher details
// @Tags teachers
// @Accept  json
// @Produce  json
// @Param id query string true "Teacher ID"
// @Param teacher body proto.CreateTeacher true "Update teacher"
// @Success 200 {string} string "Teacher updated successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /teachers/update [put]
func (h *Handler) TeacherUpdate(c *gin.Context) {
	id := c.Query("id")

	reqBody := &pb.CreateTeacher{}
	if err := c.BindJSON(reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	req := &pb.UpdateTeacher{
		Id:                id,
		Name:              reqBody.Name,
		ExperienceYears:   reqBody.ExperienceYears,
		IeltsScore:        reqBody.IeltsScore,
		ProfilePictureUrl: reqBody.ProfilePictureUrl,
		Contact:           reqBody.Contact,
		GraduatedStudents: reqBody.GraduatedStudents,
	}

	if _, err := h.Clients.Teacher.Update(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error updating teacher", "err", err)
		return
	}

	slog.Info("Teacher updated successfully")
	c.JSON(200, gin.H{"message": "Teacher updated successfully"})
}

// TeacherDelete godoc
// @Summary Delete a teacher
// @Description Delete a teacher by ID
// @Tags teachers
// @Accept  json
// @Produce  json
// @Param id query string true "Teacher ID"
// @Success 200 {string} string "Teacher deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /teachers/delete [delete]
func (h *Handler) TeacherDelete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	req := &pb.ById{Id: id}
	if _, err := h.Clients.Teacher.Delete(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error deleting teacher", "err", err)
		return
	}

	slog.Info("Teacher deleted successfully")
	c.JSON(200, gin.H{"message": "Teacher deleted successfully"})
}

// GetTeacherById godoc
// @Summary Get teacher by ID
// @Description Get a teacher by their ID
// @Tags teachers
// @Accept  json
// @Produce  json
// @Param id query string true "Teacher ID"
// @Success 200 {object} proto.TeacherRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /teachers/get [get]
func (h *Handler) GetTeacherById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	req := &pb.ById{Id: id}
	res, err := h.Clients.Teacher.GetById(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error retrieving teacher by ID", "err", err)
		return
	}

	slog.Info("Teacher retrieved successfully")
	c.JSON(200, res)
}

// GetTeacherList godoc
// @Summary Get list of teachers
// @Description Get a list of teachers with optional filters
// @Tags teachers
// @Accept  json
// @Produce  json
// @Param name query string false "Name"
// @Param experience_years_min query int false "Minimum years of experience"
// @Param experience_years_max query int false "Maximum years of experience"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} proto.GetListTeacherRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /teachers/list [get]
func (h *Handler) GetTeacherList(c *gin.Context) {
	language := c.Query("language")
	experienceMinStr := c.Query("experience_years_min")
	experienceMaxStr := c.Query("experience_years_max")

	limit, offset, err := utils.ParsePaginationParams(c, c.Query("limit"), c.Query("offset"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing pagination params", "err", err)
		return
	}

	req := &pb.GetListTeacherReq{
		Language:           language,
		ExperienceYearsMin: experienceMinStr,
		ExperienceYearsMax: experienceMaxStr,
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Clients.Teacher.GetList(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error retrieving teacher list", "err", err)
		return
	}

	slog.Info("Teacher list retrieved successfully")
	c.JSON(200, res)
}
