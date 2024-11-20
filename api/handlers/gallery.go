package handlers

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	"github.com/mirjalilova/websiteOfEverest/pkg/utils"
)

type GalleryHandler struct {
	Clients *pb.GalleryServiceClient
}

func NewGalleryHandler(clients *pb.GalleryServiceClient) *GalleryHandler {
	return &GalleryHandler{Clients: clients}
}

// GalleryCreate godoc
// @Summary Create gallery
// @Description Create a new gallery entry
// @Tags gallery
// @Accept  json
// @Produce  json
// @Param gallery body proto.CreateGallery true "Create gallery"
// @Success 200 {string} string "Gallery created successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /gallery/create [post]
func (h *Handler) GalleryCreate(c *gin.Context) {
	req := &pb.CreateGallery{}

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	_, err = h.Clients.Gallery.Create(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error creating gallery", "err", err)
		return
	}

	slog.Info("Gallery created successfully")
	c.JSON(200, gin.H{"message": "Gallery created successfully"})
}

// GalleryUpdate godoc
// @Summary Update gallery
// @Description Update gallery entry
// @Tags gallery
// @Accept  json
// @Produce  json
// @Param id query string false "Gallery ID"
// @Param gallery body proto.CreateGallery true "Gallery Update Details"
// @Success 200 {string} string "Gallery updated successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /gallery/update [put]
func (h *Handler) GalleryUpdate(c *gin.Context) {
	id := c.Query("id")
	reqBody := &pb.CreateGallery{}

	err := c.BindJSON(reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	req := &pb.UpdateGallery{
		Id:         id,
		PictureUrl: reqBody.PictureUrl,
	}

	_, err = h.Clients.Gallery.Update(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error updating gallery", "err", err)
		return
	}

	slog.Info("Gallery updated successfully")
	c.JSON(200, gin.H{"message": "Gallery updated successfully"})
}

// GalleryDelete godoc
// @Summary Delete gallery
// @Description Delete a gallery entry by ID
// @Tags gallery
// @Accept  json
// @Produce  json
// @Param id query string false "Gallery ID"
// @Success 200 {string} string "Gallery deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /gallery/delete [delete]
func (h *Handler) GalleryDelete(c *gin.Context) {
	id := c.Query("id")
	req := &pb.ById{Id: id}

	_, err := h.Clients.Gallery.Delete(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error deleting gallery", "err", err)
		return
	}

	slog.Info("Gallery deleted successfully")
	c.JSON(200, gin.H{"message": "Gallery deleted successfully"})
}

// GetGalleryById godoc
// @Summary Get gallery by ID
// @Description Get a gallery entry by ID
// @Tags gallery
// @Accept  json
// @Produce  json
// @Param id query string false "Gallery ID"
// @Success 200 {object} proto.GalleryRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /gallery/get [get]
func (h *Handler) GetGalleryById(c *gin.Context) {
	id := c.Query("id")
	req := &pb.ById{Id: id}

	res, err := h.Clients.Gallery.GetById(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error getting gallery by id", "err", err)
		return
	}

	slog.Info("Gallery retrieved successfully")
	c.JSON(200, res)
}

// GetGalleryList godoc
// @Summary Get gallery list
// @Description Get all galleries with optional filters
// @Tags gallery
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} proto.GetListGalleryRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /gallery/list [get]
func (h *Handler) GetGalleryList(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, offset, err := utils.ParsePaginationParams(c, limitStr, offsetStr)
	if err != nil {
		slog.Error("Error parsing pagination params", "err", err)
		c.JSON(400, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	req := &pb.GetListGalleryReq{
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Clients.Gallery.GetList(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error getting gallery list", "err", err)
		return
	}

	slog.Info("Gallery list retrieved successfully")
	c.JSON(200, res)
}
