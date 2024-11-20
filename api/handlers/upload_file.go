package handlers

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"log/slog"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {object} string
// @Router /file-upload [post]
func (h *Handler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		slog.Error("Error uploading file", "err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}
	defer file.Close()

	fileName := header.Filename

	tempFilePath := filepath.Join("./internal/media", fileName)
	out, err := os.Create(tempFilePath)
	if err != nil {
		slog.Error("Error creating temporary file", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create temporary file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		slog.Error("Error saving file", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	minioURL, err := h.MinIO.UploadFile(c, fileName, tempFilePath)
	if err != nil {
		slog.Error("Error uploading to MinIO", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to MinIO"})
		return
	}

	os.Remove(tempFilePath)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully upload",
		"Url":     minioURL,
	})

}
