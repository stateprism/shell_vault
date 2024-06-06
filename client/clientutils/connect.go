package clientutils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"sync"

	pb "github.com/stateprism/prisma_ca/rpc/caproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConnection struct {
	lock            sync.Mutex
	token           string
	client          pb.PrismaCaClient
	conn            *grpc.ClientConn
	ctx             context.Context
	isAuthenticated bool
}

func NewClientConnection(ctx context.Context) *ClientConnection {
	return &ClientConnection{
		lock:            sync.Mutex{},
		token:           "",
		client:          nil,
		isAuthenticated: false,
		ctx:             ctx,
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

	client := pb.NewPrismaCaClient(gc)
	cc.client = client

	return nil
}

func (cc *ClientConnection) Close() {
	cc.lock.Lock()
	defer cc.lock.Unlock()

	if cc.client == nil {
		return
	}

	cc.conn.Close()
	cc.client = nil
}

func (cc *ClientConnection) Authenticate(user string, pass string) error {
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
	ctx := metadata.NewOutgoingContext(cc.ctx, md)
	resp, err := cc.client.Authenticate(ctx, &pb.EmptyMsg{})
	if err != nil {
		return err
	}

	cc.token = resp.GetAuthToken()
	cc.isAuthenticated = true
	cc.ctx = metadata.NewOutgoingContext(cc.ctx, metadata.New(map[string]string{"authorization": fmt.Sprintf("encrypted %s", cc.token)}))
	return nil
}

func (cc *ClientConnection) GetToken() string {
	return cc.token
}

func (cc *ClientConnection) RequestCert(publicKey []byte) ([]byte, error) {
	cc.lock.Lock()
	cc.lock.Unlock()

	if cc.client == nil {
		return nil, errors.New("not connected")
	}

	resp, err := cc.client.RequestCert(cc.ctx, &pb.CertRequest{
		PublicKey: publicKey,
	})
	if err != nil {
		return nil, err
	}

	return []byte(resp.GetCert()), err
}
