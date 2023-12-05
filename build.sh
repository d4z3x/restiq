GOOS=linux GOARCH=amd64 go build
upx -9v restiq
scp ./restiq root@titan.local:/mnt/cache/restic
