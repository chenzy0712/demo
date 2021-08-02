package service

import (
	"context"

	"git.kldmp.com/learning/demo/pkg/log"

	pb "git.kldmp.com/learning/demo/bizapi"
)

type BizService struct {
	pb.UnimplementedBizServer
}

func NewBizService() *BizService {
	return &BizService{}
}

func (s *BizService) GetDataTransferProtocolList(ctx context.Context, req *pb.GetDataTransferProtocolListReq) (*pb.GetDataTransferProtocolListResp, error) {
	log.Info("Get GetDataTransferProtocolList request %+v", req)
	var resp pb.GetDataTransferProtocolListResp
	var protocol pb.Protocol
	protocol.ProtocolName = "IEC104"
	protocol.ProtocolVersion = "huauyn"
	protocol.ProtocolConfig = "{url:127.0.0.1:2404,common_addr:1}"
	resp.Protocols = append(resp.Protocols, &protocol)

	return &resp, nil
}
func (s *BizService) GetDaqProtocolList(ctx context.Context, req *pb.GetDaqProtocolListReq) (*pb.GetDaqProtocolListResp, error) {
	return &pb.GetDaqProtocolListResp{}, nil
}
func (s *BizService) CreatePhysicalDevice(ctx context.Context, req *pb.CreatePhysicalDeviceReq) (*pb.Result, error) {
	return &pb.Result{}, nil
}
func (s *BizService) CreateLogicalDevice(ctx context.Context, req *pb.CreateLogicalDeviceReq) (*pb.Result, error) {
	return &pb.Result{}, nil
}
func (s *BizService) CreateSubDeviceWithLogical(ctx context.Context, req *pb.CreateSubDeviceWithLogicalReq) (*pb.Result, error) {
	return &pb.Result{}, nil
}
func (s *BizService) CreateSubDevice(ctx context.Context, req *pb.CreateSubDeviceReq) (*pb.Result, error) {
	return &pb.Result{}, nil
}
func (s *BizService) DeleteDevice(ctx context.Context, req *pb.DeleteDeviceReq) (*pb.Result, error) {
	return &pb.Result{}, nil
}
func (s *BizService) ActivateDevice(ctx context.Context, req *pb.ActivateDeviceReq) (*pb.Result, error) {
	return &pb.Result{}, nil
}
