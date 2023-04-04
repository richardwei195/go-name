/bin/sh
SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build src/main.go

nohup ./main > logs/app.log 2>&1 &