package handlers

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
	"github.com/mirjalilova/websiteOfEverest/pkg/utils"
)

type CourseHandler struct {
	repo *repository.CourseRepository
}

func NewCourseHandler(repo *repository.CourseRepository) *CourseHandler {
	return &CourseHandler{repo: repo}
}

// CourseCreate godoc
// @Summary Create a course
// @Description Create a course
// @Tags courses
// @Accept  json
// @Produce  json
// @Param course body proto.CreateCourse true "Create course"
// @Success 200 {string} string "Course created successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courses/create [post]
func (h *Handler) CourseCreate(c *gin.Context) {
	req := &pb.CreateCourse{}

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	_, err = h.Clients.Course.Create(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error creating course", "err", err)
		return
	}

	slog.Info("Course created successfully")
	c.JSON(200, gin.H{"message": "Course created successfully"})
}

// CourseUpdate godoc
// @Summary Update a course
// @Description Update a course by ID
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id query string true "Course ID"
// @Param course body proto.CreateCourse true "Update course details"
// @Success 200 {string} string "Course updated successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courses/update [put]
func (h *Handler) CourseUpdate(c *gin.Context) {
	id := c.Query("id")
	reqBody := &pb.CreateCourse{}

	err := c.BindJSON(reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	req := &pb.UpdateCourse{
		Id:       id,
		Name:     reqBody.Name,
		Duration: reqBody.Duration,
		PictureUrl: reqBody.PictureUrl,
	}

	_, err = h.Clients.Course.Update(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error updating course", "err", err)
		return
	}

	slog.Info("Course updated successfully")
	c.JSON(200, gin.H{"message": "Course updated successfully"})
}

// CourseDelete godoc
// @Summary Delete a course
// @Description Delete a course by ID
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id query string true "Course ID"
// @Success 200 {string} string "Course deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courses/delete [delete]
func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{
		Id: id,
	}

	_, err := h.Clients.Course.Delete(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error deleting course", "err", err)
		return
	}

	slog.Info("Course deleted successfully")
	c.JSON(200, gin.H{"message": "Course deleted successfully"})
}

// GetCourseById godoc
// @Summary Get a course by ID
// @Description Get a course by its ID
// @Tags courses
// @Accept  json
// @Produce  json
// @Param id query string true "Course ID"
// @Success 200 {object} proto.CourseRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courses/get [get]
func (h *Handler) GetCourseById(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{
		Id: id,
	}

	res, err := h.Clients.Course.GetById(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error getting course by id", "err", err)
		return
	}

	slog.Info("Course retrieved successfully")
	c.JSON(200, res)
}

// GetCoursesList godoc
// @Summary Get all courses
// @Description Get a list of courses with optional filtering
// @Tags courses
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} proto.GetListCourseRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courses/list [get]
func (h *Handler) GetCoursesList(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, offset, err := utils.ParsePaginationParams(c, limitStr, offsetStr)
	if err != nil {
		slog.Error("Error parsing pagination params", "err", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req := &pb.GetListCourseReq{
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Clients.Course.GetList(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error getting courses list", "err", err)
		return
	}

	slog.Info("Courses list retrieved successfully")
	c.JSON(200, res)
}
