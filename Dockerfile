FROM docker.io/enchant/ubuntu

COPY dist /workingDir
RUN apt-get update && apt-get install -y mariadb-client
RUN mkdir /backupDir

ENV backupScript /scripts/backup
ENV restoreBackupScript /scripts/restore_backup
ENV backupDir /backupDir
ENV backupScript /scripts/backup
ENV backupScript /scripts/backup

WORKDIR /workingDir


CMD ./backupServer
