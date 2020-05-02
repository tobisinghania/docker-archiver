package main

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"syscall"
)

func main() {

	log.Printf("Backup script is %s", config.BackupCmd())
	r := gin.New()

	contextGroup := r.Group("/backupManager")
	{

		contextGroup.Static("/static/backups", config.BackupDir())
		r.Use(static.Serve("/backupManager", static.LocalFile(config.UiDir(), false)))

		apiGroup := contextGroup.Group("/api")
		{
			apiGroup.POST("/backups", createBackup)
			apiGroup.POST("/backups/restore/:backupName", restoreBackup)
			apiGroup.GET("/diskStats", getAvailableDiskSpace)
			apiGroup.GET("/backups", listBackups)
			apiGroup.GET("/version", getVersion)
			apiGroup.DELETE("/backups/:backupName", deleteBackup)
			apiGroup.POST("/uploadBackup", uploadBackup)
		}

	}

	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Port()))

}

type DiskStats struct {
	TotalBytes     uint64 `json:"totalBytes"`
	AvailableBytes uint64 `json:"availableBytes"`
}

type Version struct {
	Version string `json:"version"`
}

type CreateBackupOutput struct {
	Log string `json:"log"`
}

type Backup struct {
	Name         string `json:"name"`
	SizeBytes    int64  `json:"sizeBytes"`
	LastModified int64  `json:"lastModified"`
	DownloadLink string `json:"downloadLink"`
}

func getAvailableDiskSpace(c *gin.Context) {

	var stat syscall.Statfs_t
	syscall.Statfs(config.BackupDir(), &stat)

	stats := DiskStats{
		TotalBytes:     stat.Blocks * uint64(stat.Bsize),
		AvailableBytes: stat.Bavail * uint64(stat.Bsize),
	}
	c.JSON(200, stats)

}

func listBackups(c *gin.Context) {
	backups, err := ListBackups()

	if writeServerErr(err, c) {
		return
	}

	c.JSON(200, backups)
}

func getVersion(c *gin.Context) {
	c.JSON(200, Version{config.Version()})
}

func createBackup(c *gin.Context) {
	out, err := CreateBackup()

	if writeServerErr(err, c) {
		return
	}

	c.JSON(200, CreateBackupOutput{out})
}

func restoreBackup(c *gin.Context) {

	out, err := RestoreBackup(c.Param("backupName"))

	if writeServerErr(err, c) {
		return
	}

	c.JSON(200, CreateBackupOutput{out})

}

func deleteBackup(c *gin.Context) {
	backupName := c.Param("backupName")

	if writeServerErr(DeleteBackup(backupName), c) {
		return
	}

	c.JSON(204, nil)
}

func uploadBackup(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	if err := c.SaveUploadedFile(file, config.BackupDir()+"/"+file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func writeServerErr(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(500, c.Error(err))
		return true
	}
	return false
}
