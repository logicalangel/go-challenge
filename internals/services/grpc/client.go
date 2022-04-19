package grpc

import "google.golang.org/grpc"

func NewClient() *EsServiceClient {
	conn, err := grpc.Dial("127.0.0.1:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	c := NewEsServiceClient(conn)
	return &c
}
