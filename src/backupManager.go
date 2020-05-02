package main

import (
	"errors"
	"os"
	"strings"
)

var config Config = Config{}

func init() {

}

func DeleteBackup(name string) error {
	if strings.Contains(name, "..") {
		return errors.New("Directory up identifier is not allowed in backup name")
	}
	filePath := config.BackupDir() + "/" + name
	return os.Remove(filePath)
}

func CreateBackup() (string, error) {
	resultVal, err := execCmd(config.BackupCmd())
	if err != nil {
		return "", err
	}
	return resultVal, nil
}

func RestoreBackup(name string) (string, error) {
	resultVal, err := execCmd(config.RestoreBackupCmd() + " " + config.BackupDir() + "/" + name)
	if err != nil {
		return "", err
	}
	return resultVal, nil
}

func ListBackupPath() ([]string, error) {
	backups, err := execCmd(config.ListBackupsCmd())
	if err != nil {
		return []string{}, nil
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

func ListBackups() ([]Backup, error) {
	paths, err := ListBackupPath()
	if err != nil {
		return nil, err
	}
	result := make([]Backup, len(paths))

	for i, path := range paths {
		fi, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		result[i] = Backup{
			Name:         fi.Name(),
			SizeBytes:    fi.Size(),
			LastModified: fi.ModTime().Unix(),
			DownloadLink: `/static/backups/` + fi.Name(),
		}
	}

	return result, nil
}
