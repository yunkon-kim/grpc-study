package handler

import (
	context "context"
	pb "github.com/hermitkim1/grpc-study/protos"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// UnimplementedConfigStoreServer is representation of protobuf ApiServer
type UnimplementedConfigStoreServer struct {
}

// GetHello implements api.proto.ApiServer.GetHello
func (s *UnimplementedConfigStoreServer) Get(ctx context.Context, in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func (UnimplementedConfigStoreServer) mustEmbedUnimplementedGreeterServer() {}