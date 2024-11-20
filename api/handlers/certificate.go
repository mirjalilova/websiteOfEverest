package handlers

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/websiteOfEverest/internal/genproto/proto"
	"github.com/mirjalilova/websiteOfEverest/pkg/utils"
)

type CertificateHandler struct {
	clients *pb.CertificateServiceClient
}

func NewCertificateHandler(clients *pb.CertificateServiceClient) *CertificateHandler {
	return &CertificateHandler{clients: clients}
}

// CreateCertificate godoc
// @Summary Create a certificate
// @Description Create a new certificate
// @Tags certificates
// @Accept  json
// @Produce  json
// @Param certificate body proto.CreateCertificate true "Create certificate"
// @Success 200 {string} string "Certificate created successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /certificates/create [post]
func (h *Handler) CreateCertificate(c *gin.Context) {
	req := &pb.CreateCertificate{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	if _, err := h.Clients.Certificate.Create(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error creating certificate", "err", err)
		return
	}

	slog.Info("Certificate created successfully")
	c.JSON(200, gin.H{"message": "Certificate created successfully"})
}

// UpdateCertificate godoc
// @Summary Update a certificate
// @Description Update certificate details
// @Tags certificates
// @Accept  json
// @Produce  json
// @Param id query string true "Certificate ID"
// @Param certificate body proto.CreateCertificate true "Update certificate"
// @Success 200 {string} string "Certificate updated successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /certificates/update [put]
func (h *Handler) UpdateCertificate(c *gin.Context) {
	id := c.Query("id")

	reqBody := &pb.CreateCertificate{}
	if err := c.BindJSON(reqBody); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing request body", "err", err)
		return
	}

	req := &pb.UpdateCertificate{
		Id:             id,
		IeltsScore:     reqBody.IeltsScore,
		CefrLevel:      reqBody.CefrLevel,
		Description:    reqBody.Description,
		CertificateUrl: reqBody.CertificateUrl,
	}

	if _, err := h.Clients.Certificate.Update(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error updating certificate", "err", err)
		return
	}

	slog.Info("Certificate updated successfully")
	c.JSON(200, gin.H{"message": "Certificate updated successfully"})
}

// DeleteCertificate godoc
// @Summary Delete a certificate
// @Description Delete a certificate by ID
// @Tags certificates
// @Accept  json
// @Produce  json
// @Param id query string true "Certificate ID"
// @Success 200 {string} string "Certificate deleted successfully"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /certificates/delete [delete]
func (h *Handler) DeleteCertificate(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{Id: id}
	if _, err := h.Clients.Certificate.Delete(context.Background(), req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error deleting certificate", "err", err)
		return
	}

	slog.Info("Certificate deleted successfully")
	c.JSON(200, gin.H{"message": "Certificate deleted successfully"})
}

// GetCertificateById godoc
// @Summary Get certificate by ID
// @Description Get a certificate by its ID
// @Tags certificates
// @Accept  json
// @Produce  json
// @Param id query string true "Certificate ID"
// @Success 200 {object} proto.CertificateRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /certificates/get [get]
func (h *Handler) GetCertificateById(c *gin.Context) {
	id := c.Query("id")

	req := &pb.ById{Id: id}
	res, err := h.Clients.Certificate.GetById(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error retrieving certificate by ID", "err", err)
		return
	}

	slog.Info("Certificate retrieved successfully")
	c.JSON(200, res)
}

// GetCertificateList godoc
// @Summary Retrieve a list of certificates
// @Description Fetch a list of certificates with optional filters for name, IELTS score, CEFR level, and pagination parameters.
// @Tags certificates
// @Accept json
// @Produce json
// @Param name query string false "Filter by certificate name"
// @Param ielts_score query number false "Filter by IELTS score"
// @Param cefr_level query string false "Filter by CEFR level"
// @Param limit query int false "Number of records to retrieve (default is 10)"
// @Param offset query int false "Number of records to skip for pagination (default is 0)"
// @Success 200 {object} proto.GetListCertificateRes "A list of certificates matching the filters"
// @Failure 400 {object} map[string]string "Invalid request parameters"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /certificates/list [get]
func (h *Handler) GetCertificateList(c *gin.Context) {
	name := c.Query("name")
	ieltsScoreStr := c.Query("ielts_score")
	cefrLevel := c.Query("cefr_level")

	limit, offset, err := utils.ParsePaginationParams(c, c.Query("limit"), c.Query("offset"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		slog.Error("Error parsing pagination params", "err", err)
		return
	}

	ieltsScore, err := strconv.ParseFloat(ieltsScoreStr, 64)
	if err != nil {
		slog.Error("Invalid ielts_score value", "err", err)
		c.JSON(400, gin.H{"error": "Invalid ielts_score value"})
		return
	}

	req := &pb.GetListCertificateReq{
		Name:       name,
		IeltsScore: float32(ieltsScore),
		CefrLevel:  cefrLevel,
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Clients.Certificate.GetList(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		slog.Error("Error retrieving certificate list", "err", err)
		return
	}

	slog.Info("Certificate list retrieved successfully")
	c.JSON(200, res)
}
