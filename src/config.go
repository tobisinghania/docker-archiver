package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
}

func init() {

	viper.SetDefault("backupScript", `/scripts/backup`)
	viper.SetDefault("restoreBackupScript", `/scripts/restore_backup`)
	viper.SetDefault("backupDir", "/backupDir")
	viper.SetDefault("uiDir", "./html")
	viper.SetDefault("port", 8888)
	viper.SetDefault("version", "v0.0.1")
}

func (conf *Config) BackupCmd() string {
	return viper.GetString("backupScript")
}

func (conf *Config) RestoreBackupCmd() string {
	return viper.GetString("restoreBackupScript")
}

func (conf *Config) Version() string {
	return viper.GetString("version")
}

func (conf *Config) BackupDir() string {
	return viper.GetString("backupDir")
}

func (conf *Config) UiDir() string {
	return viper.GetString("uiDir")
}

func (conf *Config) Port() int32 {
	return viper.GetInt32("port")
}

func (conf *Config) ListBackupsCmd() string {
	if viper.IsSet("listBackupsCmd") {
		return viper.GetString("listBackupsCmd")
	} else {
		return fmt.Sprintf("ls -d1 %s", viper.GetString("backupDir")+"/*")
	}
}
