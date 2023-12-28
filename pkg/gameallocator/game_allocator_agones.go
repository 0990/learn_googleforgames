package gameallocator

import (
	pb "agones.dev/agones/pkg/allocation/go"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/0990/avatar-fight-server/msg/smsg"
	"github.com/0990/avatar-fight-server/pkg/matchmaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"strings"
)

type AllocatorAgones struct {
	conn *grpc.ClientConn
}

func NewGameAllocatorAgones(endpoint string, namespace, certFile, keyFile, cacertFile string) (GameAllocator, error) {
	cert, err := os.ReadFile(certFile)
	if err != nil {
		panic(err)
	}
	key, err := os.ReadFile(keyFile)
	if err != nil {
		panic(err)
	}
	cacert, err := os.ReadFile(cacertFile)
	if err != nil {
		return nil, err
	}

	dialOpts, err := createRemoteClusterDialOption(cert, key, cacert)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(endpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return &AllocatorAgones{
		conn: conn,
	}, nil
}

func (a *AllocatorAgones) Allocate(room *matchmaker.Match) (*net.TCPAddr, error) {
	labelUser := &smsg.GameAllocateLabelUser{}

	var userIds []string
	for userId, _ := range room.Users {
		labelUser.Users = append(labelUser.Users, &smsg.GameAllocateLabelUser_User{
			UserId: userId,
		})

		userIds = append(userIds, fmt.Sprintf("%d", userId))
	}
	//metaLabelUser := proto.MarshalTextString(labelUser)

	request := &pb.AllocationRequest{
		Namespace: "default",
		GameServerSelectors: []*pb.GameServerSelector{
			{
				MatchLabels: map[string]string{
					"agones.dev/fleet": "af-test",
				},
				GameServerState: pb.GameServerSelector_READY,
			},
		},
		Metadata: &pb.MetaPatch{
			Annotations: map[string]string{
				"users": strings.Join(userIds, ";"),
			},
		},
	}

	grpcClient := pb.NewAllocationServiceClient(a.conn)
	response, err := grpcClient.Allocate(context.Background(), request)
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%d", response.Addresses[0].Address, response.Ports[0].Port)

	result, err := net.ResolveTCPAddr("tcp", addr)
	return result, err
}

func createRemoteClusterDialOption(clientCert, clientKey, caCert []byte) (grpc.DialOption, error) {
	// Load client cert
	cert, err := tls.X509KeyPair(clientCert, clientKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	if len(caCert) != 0 {
		// Load CA cert, if provided and trust the server certificate.
		// This is required for self-signed certs.
		tlsConfig.RootCAs = x509.NewCertPool()
		if !tlsConfig.RootCAs.AppendCertsFromPEM(caCert) {
			return nil, errors.New("only PEM format is accepted for server CA")
		}
	}

	return grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)), nil
}
