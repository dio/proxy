package xdsserver

import (
	"context"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
)

type callbacks struct{}

func (c *callbacks) Report() {}
func (c *callbacks) OnStreamOpen(context.Context, int64, string) error {
	return nil
}
func (cb *callbacks) OnStreamClosed(int64) {}
func (c *callbacks) OnDeltaStreamOpen(context.Context, int64, string) error {
	return nil
}
func (c *callbacks) OnDeltaStreamClosed(int64) {}
func (cb *callbacks) OnStreamRequest(int64, *discovery.DiscoveryRequest) error {
	return nil
}
func (cb *callbacks) OnStreamResponse(context.Context, int64, *discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {
}
func (cb *callbacks) OnStreamDeltaResponse(id int64, req *discovery.DeltaDiscoveryRequest, res *discovery.DeltaDiscoveryResponse) {
}
func (cb *callbacks) OnStreamDeltaRequest(id int64, req *discovery.DeltaDiscoveryRequest) error {
	return nil
}
func (cb *callbacks) OnFetchRequest(_ context.Context, req *discovery.DiscoveryRequest) error {
	return nil
}
func (cb *callbacks) OnFetchResponse(*discovery.DiscoveryRequest, *discovery.DiscoveryResponse) {}
