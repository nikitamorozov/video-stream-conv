env GOOS=linux GOARCH=amd64 GOARM=7 go build -o video-stream-conv main.go
env GOOS=linux GOARCH=amd64 GOARM=7 go build worker.go