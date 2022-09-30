package grpcclients

import (
	"api/dry"
	pb "api/grpc/AccountService"
	"api/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var timeoutAuth = 5 * time.Second
var accountService *AccountServiceClient
var accountServiceAddr string

type AccountServiceClient struct {
	serverAddr *string
	Conn       *grpc.ClientConn // useless?
	client     pb.AccountServiceClient
}

func (asc *AccountServiceClient) Initialize(addr string) (err error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	asc.serverAddr = &addr
	asc.Conn, err = grpc.Dial(addr, opts...)
	asc.client = pb.NewAccountServiceClient(asc.Conn)
	return
}

func (asc *AccountServiceClient) AuthorizeRequest(r *http.Request) (bool, error) {
	var opts grpc.CallOption
	ctx, cancel := context.WithTimeout(context.Background(), timeoutAuth)
	defer cancel()

	body, err := io.ReadAll(r.Body)
	grpcReq := pb.RequestAuthenticationRequest{
		Base64Body: base64.StdEncoding.EncodeToString(body),
	}

	if err != nil {
		return false, err
	}
	for _, cookie := range r.Cookies() {
		if cookie.Name == "Token" {
			grpcReq.Token = cookie.Value
		}
	}
	grpcReq.Base64Signature = r.Header.Get("Signature")

	resp, err := asc.client.RequestAuthentication(ctx, &grpcReq, opts)
	if err != nil {
		return false, err
	}

	return resp.Authenticated, nil
}

func (asc *AccountServiceClient) RegisterUser(r *http.Request) (s models.AuthenticatedUser, err error) {
	u := models.NewUser{}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return
	}

	if !dry.CheckNonNil(u.Email, u.Credentials) {
		return models.AuthenticatedUser{}, fmt.Errorf("[RegisterUser] models.NewUser has nil pointer")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutAuth)
	defer cancel()
	resp, err := asc.client.UserRegistration(
		ctx,
		&pb.RegistrationRequest{
			Email:       *u.Email,
			Credentials: *u.Credentials,
		},
	)
	if err != nil {
		return
	}

	userID := resp.GetUser().GetUserId()
	email := resp.GetUser().GetEmail()
	token := resp.GetToken()
	secret := resp.GetSecret()

	s = models.AuthenticatedUser{
		User: &models.User{
			UserID: &userID,
			Email:  &email,
		},
		Token:  &token,
		Secret: &secret,
	}

	return
}

func (asc *AccountServiceClient) AuthorizeUser(r *http.Request) (s models.AuthenticatedUser, err error) {
	u := models.NewUser{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return
	}

	if !dry.CheckNonNil(u.Credentials) {
		return models.AuthenticatedUser{}, fmt.Errorf("[AuthorizeUser] models.NewUser has nil pointer")
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutAuth)
	defer cancel()
	resp, err := asc.client.UserAuthentication(
		ctx,
		&pb.UserAuthenticationRequest{
			Credentials: *u.Credentials,
		},
	)
	if err != nil {
		return
	}

	userID := resp.GetUser().GetUserId()
	email := resp.GetUser().GetEmail()
	token := resp.GetToken()
	secret := resp.GetSecret()

	s = models.AuthenticatedUser{
		User: &models.User{
			UserID: &userID,
			Email:  &email,
		},
		Token:  &token,
		Secret: &secret,
	}

	return
}

func NewAccountServiceClient(addr string) (*AccountServiceClient, error) {
	asc := AccountServiceClient{}
	err := asc.Initialize(addr)
	return &asc, err
}

func GetAccountServiceClient() *AccountServiceClient {
	var err error
	if accountService == nil {
		if accountService, err = NewAccountServiceClient(accountServiceAddr); err != nil {
			panic(fmt.Sprintf("[gRPC] account service client cannot establish connection to server, error: %v", err))
		}
	}

	return accountService
}
