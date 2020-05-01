package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
}

func init() {

	viper.SetDefault("backupScript", "ls -alh")
	viper.SetDefault("backupDir", "./backups")
	viper.SetDefault("port", 8888)
}

func (conf *Config) BackupCmd() string {
	return viper.GetString("backupScript")
}

func (conf *Config) BackupDir() string {
	return viper.GetString("backupDir")
}

func (conf *Config) Port() int32 {
	return viper.GetInt32("port")
}

func (conf *Config) ListBackupsCmd() string {
	if viper.IsSet("listBackupsCmd") {
		return viper.GetString("listBackupsCmd")
	} else {
		return fmt.Sprintf("ls -1 %s", viper.GetString("backupDir"))
	}
}
