package main

import (
	"strings"
)

var config Config = Config{}

func init() {

}

func CreateBackup() (string, error) {
	resultVal, err := execCmd(config.BackupCmd())
	if err != nil {
		return "", err
	}
	return resultVal, nil
}

func ListBackups() ([]string, error) {
	backups, err := execCmd(config.ListBackupsCmd())
	if err != nil {
		return nil, err
	}
	lines := strings.Split(backups, "\n")
	slc := []string{}

	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			slc = append(slc, line)
		}
	}
	return slc, nil
}
