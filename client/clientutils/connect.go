package clientutils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"sync"

	pb "github.com/stateprism/shell_vault/rpc/caproto"
	pbcommon "github.com/stateprism/shell_vault/rpc/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConnection struct {
	lock            sync.Mutex
	token           string
	client          pb.CertificateAuthorityClient
	conn            *grpc.ClientConn
	isAuthenticated bool
}

func NewClientConnection() *ClientConnection {
	return &ClientConnection{
		lock:            sync.Mutex{},
		token:           "",
		client:          nil,
		isAuthenticated: false,
	}
}

func (cc *ClientConnection) TryConnect(addr string) error {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	if cc.client != nil {
		return nil
	}

	gc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	cc.conn = gc

	client := pb.NewCertificateAuthorityClient(gc)
	cc.client = client

	return nil
}

func (cc *ClientConnection) Close() {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	if cc.client == nil {
		return
	}

	err := cc.conn.Close()
	if err != nil {
		return
	}
	cc.client = nil
}

func (cc *ClientConnection) Authenticate(ctx context.Context, user string, pass string) error {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	if cc.client == nil {
		return errors.New("not connected")
	}
	loginData := make([]byte, len(user)+len(pass)+1)
	copy(loginData, user)
	loginData[len(user)] = 0x1E
	copy(loginData[len(user)+1:], pass)
	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("local %s", base64.StdEncoding.EncodeToString(loginData))})
	metadata.NewOutgoingContext(ctx, md)
	resp, err := cc.client.Authenticate(ctx, &pbcommon.Empty{})
	if err != nil {
		return err
	}

	cc.token = resp.GetAuthToken()
	cc.isAuthenticated = true
	return nil
}

func (cc *ClientConnection) GetToken() string {
	return cc.token
}

func (cc *ClientConnection) addTokenToRequest(parent context.Context) context.Context {
	if !cc.isAuthenticated {
		panic("you must be authenticated to make this request")
	}
	md := metadata.New(map[string]string{"authorization": fmt.Sprintf("encrypted %s", cc.token)})
	return metadata.NewOutgoingContext(parent, md)
}

func (cc *ClientConnection) RequestUserCert(ctx context.Context, publicKey []byte) ([]byte, error) {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	ctx = cc.addTokenToRequest(ctx)

	if cc.client == nil {
		return nil, errors.New("not connected")
	}

	resp, err := cc.client.RequestUserCertificate(ctx, &pb.UserCertRequest{
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}

	return []byte(resp.GetCert()), err
}
func (cc *ClientConnection) RequestHostCert(ctx context.Context, publicKey []byte, hostnames []string) ([]byte, error) {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	ctx = cc.addTokenToRequest(ctx)

	if cc.client == nil {
		return nil, errors.New("not connected")
	}

	resp, err := cc.client.RequestUserCertificate(ctx, &pb.UserCertRequest{
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}

	return []byte(resp.GetCert()), err
}

func (cc *ClientConnection) GetCurrentCert(ctx context.Context) (string, int64, error) {
	r, err := cc.client.GetCurrentKey(ctx, &pbcommon.Empty{})
	if err != nil {
		return "", 0, err
	}

	return r.GetCert(), r.GetValidUntil(), nil
}
