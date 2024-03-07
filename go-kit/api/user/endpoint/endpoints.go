package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mokoshin0720/microservice-go/go-kit/service"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	Get endpoint.Endpoint
}

type GetReq struct {
	ID string
}

type GetResp struct {
	Name string
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Get: makeGetEndpoint(s),
	}
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetReq)
		result, _ := s.Get(ctx, req.ID)
		return GetResp{Name: result}, nil
	}
}
