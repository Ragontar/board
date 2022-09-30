package server

import (
	"account-service/crud"
	pb "account-service/grpc/AccountService"
	"account-service/telegram"
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var pendingConfirmation = make(map[string]*telegram.TelegramSession)

type grpcAccountServiceServer struct {
	pb.UnimplementedAccountServiceServer
}

func (*grpcAccountServiceServer) UserRegistration(ctx context.Context, req *pb.RegistrationRequest) (*pb.SessionResponse, error) {
	u := crud.User{
		Email:       req.GetEmail(),
		Credentials: req.GetCredentials(),
	}

	if u.Email == "" || u.Credentials == "" {
		return nil, &ErrorBadRegistrationRequest{}
	}

	s, err := crud.UserRegistration(u)
	if err != nil {
		return nil, err
	}

	resp := pb.SessionResponse{
		User: &pb.User{
			UserId: s.User.UserID,
			Email:  s.User.Email,
		},
		Token:  s.Token,
		Secret: s.Secret,
	}

	return &resp, nil
}

func (*grpcAccountServiceServer) UserAuthentication(ctx context.Context, req *pb.UserAuthenticationRequest) (resp *pb.SessionResponse, err error) {
	u := crud.User{
		Credentials: req.Credentials,
	}
	s, err := crud.AuthenticateUser(u)
	if err != nil {
		return
	}
	resp = &pb.SessionResponse{
		User: &pb.User{
			UserId: s.User.UserID,
			Email:  s.User.Email,
		},
		Token:  s.Token,
		Secret: s.Secret,
	}

	return
}

func (*grpcAccountServiceServer) RequestAuthentication(ctx context.Context, req *pb.RequestAuthenticationRequest) (*pb.RequestAuthenticationResponse, error) {
	s := crud.Session{
		Token: req.GetToken(),
	}

	if s.Token == "" {
		return nil, &ErrorBadAuthenticationRequest{}
	}

	if err := s.GetSecret(); err != nil {
		return nil, err
	}

	resp := crud.AuthenticateRequest(crud.AuthRequest{
		Base64Body:      req.GetBase64Body(),
		Session:         s,
		Base64Signature: req.GetBase64Signature(),
	})

	grpcResp := pb.RequestAuthenticationResponse{
		Authenticated: resp.Authenticated,
	}
	return &grpcResp, nil
}

func (*grpcAccountServiceServer) LinkTelegram(ctx context.Context, req *pb.TelegramSession) (*pb.TelegramConfirmationCode, error) {
	println("LinkTelegram")
	ts, err := telegram.NewTelegramSession(req.User.UserId, req.Phone)
	if err != nil {
		return nil, err
	}

	err = ts.SendCode()
	if err != nil {
		return nil, err
	}

	// key := uuid.New().String()
	key := ts.Storage.SessionID
	pendingConfirmation[key] = ts

	return &pb.TelegramConfirmationCode{Key: key}, nil
}

func (*grpcAccountServiceServer) LinkTelegramConfirmCode(ctx context.Context, req *pb.TelegramConfirmationCode) (*pb.LinkStatus, error) {
	if ts, ok := pendingConfirmation[req.Key]; ok {
		if ts == nil {
			return nil, errors.New("confirm code: nil session")
		}
		_, err := ts.ConfirmCode(req.PhoneCode)
		if err != nil {
			return nil, err
		}

		model := crud.TelegramSession{}
		model.SessionID = ts.Storage.SessionID
		model.UserID = ts.UserID
		if err != nil {
			return &pb.LinkStatus{}, err
		}

		println("INSERTING")
		err = model.Insert(ctx)
		//! Possible race
		log.Printf("Deleting pending key: %s", req.Key)
		delete(pendingConfirmation, req.Key)
		//! Possible race
		return &pb.LinkStatus{Ok: true, PasswordRequired: false}, err
	}

	return &pb.LinkStatus{Ok: false, PasswordRequired: false}, fmt.Errorf("[LinkTelegramConfirmCode] no key %v is pending", req.Key)
}

func Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAccountServiceServer(grpcServer, &grpcAccountServiceServer{})

	log.Println("[gRPC]: starting server...")
	return grpcServer.Serve(lis)
}
