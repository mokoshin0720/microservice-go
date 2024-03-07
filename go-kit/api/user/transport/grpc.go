package transport

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	endpoints "github.com/mokoshin0720/microservice-go/go-kit/endpoint"
	"github.com/mokoshin0720/microservice-go/go-kit/pb"
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
				extractCorrelationID,
			),
			gt.ServerBefore(
				displayServerRequestHeaders,
			),
			gt.ServerAfter(
				injectResponseHeader,
				injectResponseTrailer,
				injectConsumedCorrelationID,
			),
			gt.ServerAfter(
				displayServerResponseHeaders,
				displayServerResponseTrailers,
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
