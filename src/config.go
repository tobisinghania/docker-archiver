package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
}

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("BACKUP_SCRIPT", `/scripts/backup`)
	viper.SetDefault("RESTORE_BACKUP_SCRIPT", `/scripts/restore_backup`)
	viper.SetDefault("BACKUP_DIR", "/backupDir")
	viper.SetDefault("uiDir", "./html")
	viper.SetDefault("port", 8888)
	viper.SetDefault("version", "v0.0.1")

}

func (conf *Config) BackupCmd() string {
	return viper.GetString("BACKUP_SCRIPT")
}

func (conf *Config) RestoreBackupCmd() string {
	return viper.GetString("RESTORE_BACKUP_SCRIPT")
}

func (conf *Config) Version() string {
	return viper.GetString("version")
}

func (conf *Config) BackupDir() string {
	return viper.GetString("BACKUP_DIR")
}

func (conf *Config) UiDir() string {
	return viper.GetString("uiDir")
}

func (conf *Config) Port() int32 {
	return viper.GetInt32("port")
}

func (conf *Config) ListBackupsCmd() string {
	if viper.IsSet("LIST_BACKUPS_COMMAND") {
		return viper.GetString("LIST_BACKUPS_COMMAND")
	} else {
		return fmt.Sprintf("ls -d1 %s", config.BackupDir()+"/*")
	}
}
