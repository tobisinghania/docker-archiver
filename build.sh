#!/bin/sh
set -e

echo "---------------------------------------------------------------------------"
echo "Building backend"
cd src
go build -o ../dist/backupServer
cd ..
chmod +x ./dist/backupServer

echo "Done"
echo ""
echo "---------------------------------------------------------------------------"
echo "Building fontend"

cd ui
ng build --prod --outputPath ../dist/html --base-href /backupManager/

echo "Done"
echo ""
