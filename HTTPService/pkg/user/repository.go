package user

import (
	"context"

	erro "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/errors"

	"github.com/fCalixto-Gb/Final-Project/GRPCServiceA/pkg/user/proto"
	ent "github.com/fCalixto-Gb/Final-Project/HTTPService/pkg/entities"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

type gRPCstub struct {
	conn   *grpc.ClientConn
	logger log.Logger
}

func NewGRPClient(log log.Logger, con *grpc.ClientConn) *gRPCstub {
	return &gRPCstub{
		conn:   con,
		logger: log,
	}
}

func (g *gRPCstub) Get(ctx context.Context, req ent.GetUserReq) (ent.GetUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.GetUserReq{
		Id: req.Id,
	}
	gRPCresp, err := client.GetUser(ctx, request)
	if err != nil {
		return ent.GetUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.GetUserResp{
			Name:    gRPCresp.Name,
			Country: gRPCresp.Country,
			Job:     gRPCresp.Job,
			Age:     gRPCresp.Age,
			Id:      gRPCresp.Id,
			Created: gRPCresp.Created,
			Email:   gRPCresp.Email,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.GetUserResp{}, err
	}
}

func (g *gRPCstub) Create(ctx context.Context, req ent.CreateUserReq) (ent.CreateUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.CreateUserReq{
		Name:    req.Name,
		Age:     req.Age,
		Country: req.Country,
		Job:     req.Job,
		Pwd:     req.Pwd,
		Email:   req.Email,
	}
	gRPCresp, err := client.CreateUser(ctx, request)
	if err != nil {
		return ent.CreateUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.CreateUserResp{
			Id:      gRPCresp.Id,
			Created: gRPCresp.Created,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.CreateUserResp{}, err
	}
}

func (g *gRPCstub) Delete(ctx context.Context, req ent.DeleteUserReq) (ent.DeleteUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.DeleteUserReq{
		Id: req.Id,
	}
	gRPCresp, err := client.DeleteUser(ctx, request)
	if err != nil {
		return ent.DeleteUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.DeleteUserResp{
			Deleted: gRPCresp.Deleted,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.DeleteUserResp{}, err
	}
}

func (g *gRPCstub) Update(ctx context.Context, req ent.UpdateUserReq) (ent.UpdateUserResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.UpdateUserReq{
		Name:    req.Name,
		Age:     req.Age,
		Country: req.Country,
		Job:     req.Job,
		Pwd:     req.Pwd,
		Email:   req.Email,
	}
	gRPCresp, err := client.UpdateUser(ctx, request)
	if err != nil {
		return ent.UpdateUserResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.UpdateUserResp{
			Updated: gRPCresp.Updated,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.UpdateUserResp{}, err
	}
}

func (g *gRPCstub) Authenticate(ctx context.Context, req ent.AuthenticateReq) (ent.AuthenticateResp, error) {
	client := proto.NewUserServicesClient(g.conn)
	request := &proto.AuthenticateReq{
		Pwd:   req.Pwd,
		Email: req.Email,
	}
	gRPCresp, err := client.Authenticate(ctx, request)
	if err != nil {
		return ent.AuthenticateResp{}, err
	}
	gRPCcode := gRPCresp.Error.Code
	if gRPCcode == 0 {
		return ent.AuthenticateResp{
			Token: gRPCresp.Token,
		}, nil
	} else {
		err = erro.ErrFromGRPCcode(gRPCcode)
		return ent.AuthenticateResp{}, err
	}
}
