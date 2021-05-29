package monitoring_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
)

func TestHTTP(t *testing.T) {
	var callCount = 0
	handlerFunc := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		callCount++
	})
	handlerOut := monitoring.WrapHTTP(handlerFunc, "test")
	if handlerOut == nil {
		t.Error("expected non-nil http handler")
	}

	// Verify that the output is unchanged by checking call behavior
	handlerFunc(nil, nil)
	handlerOut.ServeHTTP(nil, nil)
	if callCount != 2 {
		t.Errorf("expected two calls to handler, got: %v", callCount)
	}
}

func TestHTTPTransport(t *testing.T) {
	var callCount = 0
	transportFunc := func(*http.Request) (*http.Response, error) {
		callCount++
		return nil, nil
	}
	transport := &rountTripImpl{
		roundTripFunc: transportFunc,
	}
	transportOut := monitoring.WrapHTTPTransport(transport)
	if transportOut == nil {
		t.Error("expected non-nil http transport")
	}

	// Verify that the output is unchanged by checking call behavior
	_, _ = transport.RoundTrip(nil)
	_, _ = transportOut.RoundTrip(nil)
	if callCount != 2 {
		t.Errorf("expected two calls to transport, got: %v", callCount)
	}
}

var _ http.RoundTripper = &rountTripImpl{}

type rountTripImpl struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func (r *rountTripImpl) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.roundTripFunc(req)
}

func TestGrpc(t *testing.T) {
	clientInterceptor := monitoring.Default.GRPCUnaryClientInterceptor()
	if clientInterceptor != nil {
		t.Errorf("expected nil client interceptor, got: %v", clientInterceptor)
	}

	unaryInterceptor := monitoring.Default.GRPCUnaryServerInterceptor()
	if unaryInterceptor != nil {
		t.Errorf("expected nil unary  interceptor, got: %v", unaryInterceptor)
	}
}

func TestNullSpan(t *testing.T) {
	ctx := context.Background()
	ctxOut, span := monitoring.StartSpan(ctx, "test")

	if ctxOut != ctx {
		t.Error("expected no change to context")
	}

	err := span.Finish()
	if err != nil {
		t.Error(err)
	}
}

func TestAddField(t *testing.T) {
	// Adding fields should do nothing
	// Verify these calls do not panic
	monitoring.AddFieldToTrace(context.Background(), "test", "test")
	monitoring.AddFieldToSpan(context.Background(), "test", "test")

}

func TestNullClose(t *testing.T) {
	err := monitoring.Close()
	if err != nil {
		t.Error(err)
	}
}
