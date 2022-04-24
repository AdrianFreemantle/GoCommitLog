package server

import (
	api "github.com/AdrianFreemantle/GoCommitLog/api/v1"
)

type Config struct {
	CommitLog CommitLog
}

type grpcServer struct {
	api.UnimplementedLogServer
	*Config
}

var _ api.LogServer = (*grpcServer)(nil)

func newgrpcServer(config *Config) (srv *grpcServer, err error) {
	srv = &grpcServer{
		Config: config,
	}
	return srv, nil
}

type CommitLog interface {
	Append(*api.Record) (uint64, error)
	Read(uint64) (*api.Record, error)
}
