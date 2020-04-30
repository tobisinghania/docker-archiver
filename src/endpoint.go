package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"syscall"
)

func main() {

	log.Printf("Backup script is %s", config.BackupCmd())
	r := gin.Default()
	r.Static("/static/backups", config.BackupDir())

	r.POST("/backup", createBackup)
	r.GET("/diskStats", getAvailableDiskSpace)
	r.GET("/backups", listBackups)
	r.Run("localhost:8888")

}

type DiskStats struct {
	TotalBytes     uint64 `json:"totalBytes"`
	AvailableBytes uint64 `json:"availableBytes"`
}

type CreateBackupOutput struct {
	Log string `json:"log"`
}

func getAvailableDiskSpace(c *gin.Context) {

	var stat syscall.Statfs_t

	wd, err := os.Getwd()

	if err != nil {
		log.Printf("Error loading disk stats: %s", err)
		c.JSON(500, c.Error(err))
	}
	syscall.Statfs(wd, &stat)

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

func createBackup(c *gin.Context) {
	out, err := CreateBackup()

	if writeServerErr(err, c) {
		return
	}

	c.JSON(200, CreateBackupOutput{out})
}

func writeServerErr(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(500, c.Error(err))
		return true
	}
	return false
}
