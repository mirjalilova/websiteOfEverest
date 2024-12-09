package handlers

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	"github.com/mirjalilova/websiteOfEverest/internal/storage/repository"
	"github.com/mirjalilova/websiteOfEverest/pkg/utils"
)

type BranchesHandler struct {
	repo *repository.BranchesRepository
}

func NewBranchesHandler(repo *repository.BranchesRepository) *BranchesHandler {
	return &BranchesHandler{repo: repo}
}

// BranchesCreate godoc
// @Summary Create branches
// @Description Create branches
// @Tags branches
// @Accept  json
// @Produce  json
// @Param branches body proto.CreateBranches true "Create branches"
// @Success 200 {string} string "Create branches successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /branches/create [post]
func (h *Handler) BranchesCreate(c *gin.Context) {
	req := &pb.CreateBranches{}

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	_, err = h.Clients.Branches.Create(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error creating branches", "err", err)
		return
	}

	slog.Info("Branches created successfully")
	c.JSON(200, gin.H{"message": "Branches created successfully"})
}

// BranchesUpdate godoc
// @Summary Update branches
// @Description Update branches
// @Tags branches
// @Accept  json
// @Produce  json
// @Param id query string false "Branches ID"
// @Param category body proto.CreateBranches true "Branches Update Details"
// @Success 200 {string} string "Branches updated successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /branches/update [put]
func (h *Handler) BranchesUpdate(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "id is required"})
		return
	}
	
	reqBody := &pb.CreateBranches{}

	err := c.BindJSON(reqBody)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	req := &pb.UpdateBranches{
		Id:          id,
		Name:        reqBody.Name,
		Description: reqBody.Description,
		GoogleUrl:   reqBody.GoogleUrl,
		YandexUrl:   reqBody.YandexUrl,
		ImgUrl:      reqBody.ImgUrl,
		Contact:     reqBody.Contact,
	}

	_, err = h.Clients.Branches.Update(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error updating branches", "err", err)
		return
	}

	slog.Info("Branches updated successfully")
	c.JSON(200, gin.H{"message": "Branches updated successfully"})
}

// BranchesDelete godoc
// @Summary Delete a branches
// @Description Delete a branches by ID
// @Tags branches
// @Accept  json
// @Produce  json
// @Param id query string false "Branches ID"
// @Success 200 {string} string "Branches deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /branches/delete [delete]
func (h *Handler) BranchesDelete(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{
		Id: id,
	}

	_, err := h.Clients.Branches.Delete(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error deleting branches", "err", err)
		return
	}

	slog.Info("Branches deleted successfully")
	c.JSON(200, gin.H{"message": "Branches deleted successfully"})
}

// GetBranchesById godoc
// @Summary Get branches by ID
// @Description Get an branches by their ID
// @Tags branches
// @Accept  json
// @Produce  json
// @Param id query string false "Branches ID"
// @Success 200 {object} proto.BranchesRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /branches/get [get]
func (h *Handler) GetBranchesById(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{
		Id: id,
	}

	res, err := h.Clients.Branches.GetById(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error getting branches by id", "err", err)
		return
	}

	slog.Info("Branches retrieved successfully")
	c.JSON(200, res)
}

// GetBranchesList godoc
// @Summary Get all branches
// @Description Get all branches with optional filtering
// @Tags branches
// @Accept  json
// @Produce  json
// @Param name query string false "Name"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} proto.GetListBranchesRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /branches/list [get]
func (h *Handler) GetBranchesList(c *gin.Context) {
	name := c.Query("name")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, offset, err := utils.ParsePaginationParams(c, limitStr, offsetStr)
	if err != nil {
		slog.Error("Error parsing pagination params from request body: ", "err", err)
		return
	}

	req := &pb.GetListBranchesReq{
		Name: name,
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Clients.Branches.GetList(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error getting branches list", "err", err)
		return
	}

	slog.Info("Branches list retrieved successfully")
	c.JSON(200, res)
}
