package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewHealthCheck(
	gridServerTarget,
	checkerServerTarget,
	turnControllerServerTarget,
	repoServerTarget,
	matchMakerServerTarget string,
) (*HealthChecker, error) {

	grid, err := connectHealthClient(gridServerTarget)
	if err != nil {
		return nil, err
	}
	checker, err := connectHealthClient(checkerServerTarget)
	if err != nil {
		return nil, err
	}
	turnController, err := connectHealthClient(turnControllerServerTarget)
	if err != nil {
		return nil, err
	}
	repo, err := connectHealthClient(repoServerTarget)
	if err != nil {
		return nil, err
	}
	matchMaker, err := connectHealthClient(matchMakerServerTarget)
	if err != nil {
		return nil, err
	}

	return &HealthChecker{
		grid:           grid,
		checker:        checker,
		turnController: turnController,
		repo:           repo,
		matchMaker:     matchMaker,
	}, nil
}

func connectHealthClient(target string) (grpc_health_v1.HealthClient, error) {
	conn, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(monitoring.UnaryClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}
	return grpc_health_v1.NewHealthClient(conn), nil
}

type HealthChecker struct {
	grid           grpc_health_v1.HealthClient
	checker        grpc_health_v1.HealthClient
	turnController grpc_health_v1.HealthClient
	repo           grpc_health_v1.HealthClient
	matchMaker     grpc_health_v1.HealthClient

	healthy      *bool
	healthyMutex sync.RWMutex
}

func (h *HealthChecker) Run(d time.Duration, close <-chan struct{}) {
	for range time.Tick(d) {
		select {
		case <-close:
			return
		default:
		}

		h.check(context.Background())
	}
}

func (h *HealthChecker) Healthy() bool {
	h.healthyMutex.RLock()
	defer h.healthyMutex.RUnlock()
	if h.healthy == nil {
		return false
	}
	return *h.healthy
}

func (h *HealthChecker) check(ctx context.Context) {
	h.healthyMutex.Lock()
	defer h.healthyMutex.Unlock()

	oldHealthy := h.healthy
	newHealthy := isHealthy(ctx, h.grid) &&
		isHealthy(ctx, h.checker) &&
		isHealthy(ctx, h.turnController) &&
		isHealthy(ctx, h.repo) &&
		isHealthy(ctx, h.matchMaker)
	h.healthy = &newHealthy
	if oldHealthy == nil || *oldHealthy != *h.healthy {
		log.Printf("Healthy: %v", *h.healthy)
	}
}

// isHealthy performs a health check against the specified health client.
// true is returned iff the check is successful and the response is SERVING.
func isHealthy(ctx context.Context, client grpc_health_v1.HealthClient) bool {
	res, err := client.Check(context.Background(), &grpc_health_v1.HealthCheckRequest{})
	if err != nil || res.GetStatus() != grpc_health_v1.HealthCheckResponse_SERVING {
		return false
	}
	return true
}
