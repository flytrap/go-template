package api

import (
	"context"

	"github.com/flytrap/go-template/internal/services"
	"github.com/flytrap/go-template/pb/v1"
	"github.com/mitchellh/mapstructure"
)

func NewLogApi(logService services.LogService) pb.LogServiceServer {
	return &LogApi{logService: logService}
}

type LogApi struct {
	logService services.LogService
	pb.UnimplementedLogServiceServer
}

func (s *LogApi) ListLog(ctx context.Context, req *pb.QueryRequest) (*pb.QueryLogResp, error) {
	results := []*pb.Log{}
	n, err := s.logService.List(req.Q, req.Page, req.Size, req.Order, &results)
	if err != nil {
		return &pb.QueryLogResp{Ret: &pb.RetInfo{Status: false, Msg: err.Error()}}, err
	}
	return &pb.QueryLogResp{Ret: &pb.RetInfo{Status: true}, Data: results, Total: n}, nil
}

func (s *LogApi) CreateLog(ctx context.Context, req *pb.Log) (*pb.RetInfo, error) {
	info := map[string]interface{}{}
	mapstructure.Decode(&req, info)
	err := s.logService.Create(ctx, info)
	if err != nil {
		return &pb.RetInfo{Status: false, Msg: err.Error()}, err
	}
	return &pb.RetInfo{Status: true}, nil
}

func (s *LogApi) UpdateLog(ctx context.Context, req *pb.Log) (*pb.RetInfo, error) {
	info := map[string]interface{}{}
	mapstructure.Decode(&req, info)
	err := s.logService.Update(ctx, info)
	if err != nil {
		return &pb.RetInfo{Status: false, Msg: err.Error()}, err
	}
	return &pb.RetInfo{Status: true}, nil
}

func (s *LogApi) DeleteLog(ctx context.Context, req *pb.DeleteIds) (*pb.RetInfo, error) {
	err := s.logService.Delete(ctx, req.Ids)
	if err != nil {
		return &pb.RetInfo{Status: false, Msg: err.Error()}, err
	}
	return &pb.RetInfo{Status: true}, nil
}
