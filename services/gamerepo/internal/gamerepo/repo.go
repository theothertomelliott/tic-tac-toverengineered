package gamerepo

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game/rpcrepository"
)

func NewServer(repo game.Repository) rpcrepository.RepoServer {
	return &repoServer{
		repo: repo,
	}
}

type repoServer struct {
	rpcrepository.UnimplementedRepoServer

	repo game.Repository
}

func (r *repoServer) New(ctx context.Context, req *rpcrepository.NewRequest) (*rpcrepository.NewResponse, error) {
	var resp *rpcrepository.NewResponse = &rpcrepository.NewResponse{}
	id, err := r.repo.New(ctx)
	if err != nil {
		return nil, err
	}
	resp.ID = string(id)
	return resp, nil
}

func (r *repoServer) List(ctx context.Context, req *rpcrepository.ListRequest) (*rpcrepository.ListResponse, error) {
	resp := &rpcrepository.ListResponse{}
	ids, err := r.repo.List(ctx, req.Max, req.Offset)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		resp.ID = append(resp.ID, string(id))
	}
	return resp, nil
}

func (r *repoServer) Exists(ctx context.Context, req *rpcrepository.ExistsRequest) (*rpcrepository.ExistsResponse, error) {
	resp := &rpcrepository.ExistsResponse{}
	exists, err := r.repo.Exists(ctx, game.ID(req.ID))
	if err != nil {
		return nil, err
	}
	resp.Exists = exists
	return resp, nil
}
