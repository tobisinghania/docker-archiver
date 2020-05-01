FROM docker.io/enchant/ubuntu

COPY dist /workingDir

WORKDIR /workingDir

CMD ./backupServer
