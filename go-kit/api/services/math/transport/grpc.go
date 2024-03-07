package transport

import (
	"context"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/mokoshin0720/microservice-go/go-kit/helper"
	"github.com/mokoshin0720/microservice-go/go-kit/services/math/endpoints"
	"github.com/mokoshin0720/microservice-go/go-kit/services/math/pb"
)

// NewGRPCServer initializes a new gRPC server
type gRPCServer struct {
	add      gt.Handler
	subtract gt.Handler
	pb.UnimplementedMathServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) *gRPCServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoints.Add,
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
		subtract: gt.NewServer(
			endpoints.Subtract,
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

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Subtract(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.subtract.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoints.MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.MathResponse{Result: resp.Result}, nil
}
