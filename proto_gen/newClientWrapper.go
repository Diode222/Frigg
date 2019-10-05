package proto_gen

import "google.golang.org/grpc"

func NewWordSplitServiceClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewWordSplitServiceClient(cc)
}
