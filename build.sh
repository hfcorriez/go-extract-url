if [ "$1" = "linux" ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api.linux *.go
else
  go build -o bin/api *.go
fi