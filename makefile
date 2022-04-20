Es:
	- CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o es ./cmd/es
Uss:
	- CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o uss ./cmd/uss
