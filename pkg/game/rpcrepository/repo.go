package rpcrepository

import (
	context "context"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

func NewServer(repo game.Repository) RepoServer {
	return &repoServer{
		repo: repo,
	}
}

type repoServer struct {
	repo game.Repository
}

func (r *repoServer) New(ctx context.Context, req *NewRequest) (*NewResponse, error) {
	var resp *NewResponse = &NewResponse{}
	id, err := r.repo.New()
	if err != nil {
		return nil, err
	}
	resp.ID = string(id)
	return resp, nil
}

func (r *repoServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	resp := &ListResponse{}
	ids, err := r.repo.List(req.Max, req.Offset)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		resp.ID = append(resp.ID, string(id))
	}
	return resp, nil
}

func (r *repoServer) Exists(ctx context.Context, req *ExistsRequest) (*ExistsResponse, error) {
	resp := &ExistsResponse{}
	exists, err := r.repo.Exists(game.ID(req.ID))
	if err != nil {
		return nil, err
	}
	resp.Exists = exists
	return resp, nil
}
