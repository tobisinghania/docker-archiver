FROM docker.io/enchant/ubuntu

COPY dist /workingDir
RUN mkdir /backupDir

ENV backupScript /scripts/backup
ENV restoreBackupScript /scripts/restore_backup
ENV backupDir /backupDir
ENV backupScript /scripts/backup
ENV backupScript /scripts/backup

WORKDIR /workingDir

RUN apt-get update && apt-get install -y software-properties-common

RUN apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0xF1656F24C74CD1D8 \
    && add-apt-repository "deb [arch=amd64,arm64,ppc64el] http://mariadb.mirror.liquidtelecom.com/repo/10.4/ubuntu $(lsb_release -cs) main"

RUN apt-get update && apt-get install -y mariadb-client

CMD ./backupServer
