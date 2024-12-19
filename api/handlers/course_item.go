package handlers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	"github.com/mirjalilova/websiteOfEverest/pkg/utils"
)

type CourseItemHandler struct {
	clients *pb.CourseItemServiceClient
}

func NewCourseItemHandler(clients *pb.CourseItemServiceClient) *CourseItemHandler {
	return &CourseItemHandler{clients: clients}
}

// CreateCourseItem godoc
// @Summary Create a course item
// @Description Create a new course item
// @Tags courseItems
// @Accept  json
// @Produce  json
// @Param course_item body proto.CreateCourseItem true "Create course item"
// @Success 200 {string} string "Course item created successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courseItems/create [post]
func (h *Handler) CreateCourseItem(c *gin.Context) {
	req := &pb.CreateCourseItem{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	if _, err := h.Clients.CourseItem.Create(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error creating course item", "err", err)
		return
	}

	slog.Info("Course item created successfully")
	c.JSON(200, gin.H{"message": "Course item created successfully"})
}

// UpdateCourseItem godoc
// @Summary Update a course item
// @Description Update course item details
// @Tags courseItems
// @Accept  json
// @Produce  json
// @Param id query string true "Course Item ID"
// @Param course_item body proto.CreateCourseItem true "Update course item"
// @Success 200 {string} string "Course item updated successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courseItems/update [put]
func (h *Handler) UpdateCourseItem(c *gin.Context) {
	id := c.Query("id")
	reqBody := &pb.CreateCourseItem{}

	if err := c.BindJSON(reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	req := &pb.UpdateCourseItem{
		Id:            id,
		CourseId:      reqBody.CourseId,
	    Descritption:  reqBody.Descritption,
		Price:         reqBody.Price,
		DaysPerWeek:   reqBody.DaysPerWeek,
		LessonHours:   reqBody.LessonHours,
		WeekDays: reqBody.WeekDays,
	}

	if _, err := h.Clients.CourseItem.Update(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error updating course item", "err", err)
		return
	}

	slog.Info("Course item updated successfully")
	c.JSON(200, gin.H{"message": "Course item updated successfully"})
}

// DeleteCourseItem godoc
// @Summary Delete a course item
// @Description Delete a course item by ID
// @Tags courseItems
// @Accept  json
// @Produce  json
// @Param id query string true "Course Item ID"
// @Success 200 {string} string "Course item deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courseItems/delete [delete]
func (h *Handler) DeleteCourseItem(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	req := &pb.ById{Id: id}
	if _, err := h.Clients.CourseItem.Delete(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error deleting course item", "err", err)
		return
	}

	slog.Info("Course item deleted successfully")
	c.JSON(200, gin.H{"message": "Course item deleted successfully"})
}

// GetCourseItemById godoc
// @Summary Get course item by ID
// @Description Get a course item by its ID
// @Tags courseItems
// @Accept  json
// @Produce  json
// @Param id query string true "Course Item ID"
// @Success 200 {object} proto.CourseItemRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courseItems/get [get]
func (h *Handler) GetCourseItemById(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{Id: id}
	res, err := h.Clients.CourseItem.GetById(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error retrieving course item by ID", "err", err)
		return
	}

	slog.Info("Course item retrieved successfully")
	c.JSON(200, res)
}

// GetCourseItemList godoc
// @Summary Get list of course items
// @Description Get a list of course items with optional filters
// @Tags courseItems
// @Accept  json
// @Produce  json
// @Param course_id query string false "Course ID"
// @Param max_price query number false "Maximum price"
// @Param min_price query number false "Minimum price"
// @Param days_per_week query int false "Days Per Week"
// @Param lesson_hours query int false "Lesson Hours"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} proto.GetListCourseItemRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /courseItems/list [get]
func (h *Handler) GetCourseItemList(c *gin.Context) {
	courseID := c.Query("course_id")
	maxPriceStr := c.Query("max_price")
	minPriceStr := c.Query("min_price")
	daysPerWeekStr := c.Query("days_per_week")
	lessonHoursStr := c.Query("lesson_hours")

	limit, offset, err := utils.ParsePaginationParams(c, c.Query("limit"), c.Query("offset"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing pagination params", "err", err)
		return
	}

	var daysPerWeek, lessonHours int
	var maxPrice, minPrice float64

	if daysPerWeekStr != "" {
		daysPerWeek, err = strconv.Atoi(daysPerWeekStr)
		if err != nil {
			slog.Error("Invalid days_per_week value", "err", err)
			c.JSON(400, gin.H{"error": "Invalid days_per_week value"})
			return
		}
	}
	if lessonHoursStr != "" {
		lessonHours, err = strconv.Atoi(lessonHoursStr)
		if err != nil {
			slog.Error("Invalid lesson_hours value", "err", err)
			c.JSON(400, gin.H{"error": "Invalid lesson_hours value"})
			return
		}
	}
	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			slog.Error("Invalid max_price value", "err", err)
			c.JSON(400, gin.H{"error": "Invalid max_price value"})
			return
		}
	}
	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			slog.Error("Invalid min_price value", "err", err)
			c.JSON(400, gin.H{"error": "Invalid min_price value"})
			return
		}
	}
	
	req := &pb.GetListCourseItemReq{
		CourseId:    courseID,
		MaxPrice:    float32(maxPrice),
		MinPrice:    float32(minPrice),
		DaysPerWeek: int32(daysPerWeek),
		LessonHours: int32(lessonHours),
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Clients.CourseItem.GetList(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error retrieving course item list", "err", err)
		return
	}

	slog.Info("Course item list retrieved successfully")
	c.JSON(200, res)
}
