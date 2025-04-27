# pastey

- Generate a windows executable
```
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -o pastey.exe -ldflags "-H windowsgui" ./cmd/main.go
```

- Generate a linux executable
```
go build -o pastey -ldflags "-s -w" ./cmd/main.go
chmod +x pastey
./pastey-linux
```