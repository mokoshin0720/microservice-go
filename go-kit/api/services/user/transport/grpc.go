package transport

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/mokoshin0720/microservice-go/go-kit/helper"
	"github.com/mokoshin0720/microservice-go/go-kit/services/user/endpoints"
	"github.com/mokoshin0720/microservice-go/go-kit/services/user/pb"
)

// NewGRPCServer initializes a new gRPC server
type gRPCServer struct {
	add gt.Handler
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) *gRPCServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoints.Get,
			decodeMathRequest,
			encodeMathResponse,
			gt.ServerBefore(
				helper.ExtractCorrelationID,
			),
			gt.ServerBefore(
				helper.DisplayServerRequestHeaders,
			),
			gt.ServerAfter(
				helper.InjectResponseHeader,
				helper.InjectResponseTrailer,
				helper.InjectConsumedCorrelationID,
			),
			gt.ServerAfter(
				helper.DisplayServerResponseHeaders,
				helper.DisplayServerResponseTrailers,
			),
		),
	}
}

func (s *gRPCServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetRequest)
	return endpoints.GetReq{ID: req.Id}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetResp)
	return &pb.GetResponse{Name: resp.Name}, nil
}
