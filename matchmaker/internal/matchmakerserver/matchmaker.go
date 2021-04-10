package matchmakerserver

import (
	"context"

	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker"
	"github.com/theothertomelliott/tic-tac-toverengineered/matchmaker/pkg/rpcmatchmaker"
)

var _ rpcmatchmaker.MatchMakerServer = &matchMaker{}

func NewServer(m matchmaker.MatchMaker) rpcmatchmaker.MatchMakerServer {
	return &matchMaker{
		m: m,
	}
}

type matchMaker struct {
	m matchmaker.MatchMaker
}

func (m *matchMaker) Request(ctx context.Context, req *rpcmatchmaker.RequestRequest) (*rpcmatchmaker.RequestResponse, error) {
	reqID, err := m.m.Request(ctx)
	if err != nil {
		return nil, err
	}
	return &rpcmatchmaker.RequestResponse{
		RequestId: string(reqID),
	}, nil
}

func (m *matchMaker) Check(ctx context.Context, req *rpcmatchmaker.CheckRequest) (*rpcmatchmaker.CheckResponse, error) {
	match, err := m.m.Check(ctx, matchmaker.RequestID(req.RequestId))
	if err != nil {
		return nil, err
	}
	if match == nil {
		return &rpcmatchmaker.CheckResponse{}, nil
	}
	return &rpcmatchmaker.CheckResponse{
		Match: &rpcmatchmaker.Match{
			GameId: string(match.Game),
			Mark:   match.Mark,
			Token:  string(match.Token),
		},
	}, nil
}
