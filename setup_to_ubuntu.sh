#!/bin/bash

echo "sudo apt update && sudo apt upgrade -y"
sudo apt update && sudo apt upgrade -y
echo "---------------------------------------------"

# Python3 and pip3
echo "sudo apt install python3 python3-pip -y"
sudo apt install python3 python3-pip -y
echo "---------------------------------------------"

# Python libs
echo "pip3 install flask numpy scipy pandas matplotlib"
pip3 install flask numpy scipy pandas matplotlib
echo "---------------------------------------------"

# Go
echo "sudo apt install golang-go -y"
sudo apt install golang-go -y
echo "---------------------------------------------"

# Go frameworks and utils
echo "Instaling more frameworks and utils"
go get -u github.com/gin-gonic/gin
go get -u github.com/gorilla/mux
go get -u github.com/spf13/cobra
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/go-sql-driver/mysql                # MySQL
go get -u go.mongodb.org/mongo-driver/mongo             # MongoDB
go get -u github.com/bluebreezecf/opentsdb-goclient     # OpenTSDB client Go
go get -u gopkg.in/rethinkdb/rethinkdb-go.v6            # RethinkDB Go
go get -u github.com/lib/pq                             # PostgreSQL Go

# check setup
go list github.com/bluebreezecf/opentsdb-goclient
go list gopkg.in/rethinkdb/rethinkdb-go.v6
go list github.com/lib/pq
echo "---------------------------------------------"

# Add Go bin to PATH
echo 'export PATH=$PATH:/root/go/bin' >> ~/.bashrc
source ~/.bashrc
echo "---------------------------------------------"

# Check version
echo "Check version"
python3 --version
pip3 --version
go version
echo "---------------------------------------------"

echo "End setup!"

