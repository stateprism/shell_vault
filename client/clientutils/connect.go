package clientutils

import (
	"context"
	"errors"
	"sync"

	pb "github.com/stateprism/prisma_ca/rpc/caproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConnection struct {
	lock            sync.Mutex
	token           []byte
	client          pb.PrismaCaClient
	conn            *grpc.ClientConn
	ctx             context.Context
	isAuthenticated bool
}

func NewClientConnection(ctx context.Context) *ClientConnection {
	return &ClientConnection{
		lock:            sync.Mutex{},
		token:           nil,
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

	req := NewUsrPwRequest(user, pass)

	b := req.ToBytes()
	resp, err := cc.client.Authenticate(cc.ctx, &pb.AuthRequest{AuthRequest: string(b)})
	if err != nil {
		return err
	}

	cc.token = []byte(resp.GetAuthToken())
	return nil
}

func (cc *ClientConnection) GetToken() []byte {
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

	return resp.GetCert(), err
}
