package spotinst

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spotinst/spotinst-sdk-go/spotinst/util/uritemplates"
)

// A Protocol represents the type of an application protocol.
type Protocol int

const (
	// ProtocolHTTP represents the Hypertext Transfer Protocol (HTTP) protocol.
	ProtocolHTTP Protocol = iota

	// ProtocolHTTPS represents the Hypertext Transfer Protocol (HTTP) within
	// a connection encrypted by Transport Layer Security, or its predecessor,
	// Secure Sockets Layer.
	ProtocolHTTPS

	// ProtocolHTTP2 represents the Hypertext Transfer Protocol (HTTP) protocol
	// version 2.
	ProtocolHTTP2
)

var protocolName = map[Protocol]string{
	ProtocolHTTP:  "HTTP",
	ProtocolHTTPS: "HTTPS",
	ProtocolHTTP2: "HTTP2",
}

func (p Protocol) String() string {
	return protocolName[p]
}

// A ReadinessStatus represents the readiness status of a target.
type ReadinessStatus int

const (
	// StatusReady represents a ready state.
	StatusReady ReadinessStatus = iota

	// StatusMaintenance represents a maintenance state.
	StatusMaintenance
)

var readinessStatusName = map[ReadinessStatus]string{
	StatusReady:       "READY",
	StatusMaintenance: "MAINTENANCE",
}

func (s ReadinessStatus) String() string {
	return readinessStatusName[s]
}

// A HealthinessStatus represents the healthiness status of a target.
type HealthinessStatus int

const (
	// StatusUnknown represents an unknown state.
	StatusUnknown HealthinessStatus = iota

	// StatusHealthy represents a healthy state.
	StatusHealthy

	// StatusUnhealthy represents an unhealthy state.
	StatusUnhealthy
)

var healthinessStatusName = map[HealthinessStatus]string{
	StatusUnknown:   "UNKNOWN",
	StatusHealthy:   "HEALTHY",
	StatusUnhealthy: "UNHEALTHY",
}

func (s HealthinessStatus) String() string {
	return healthinessStatusName[s]
}

// A Strategy represents the load balancing methods used to determine which
// application server to send a request to.
type Strategy int

const (
	// StrategyRandom represents a random load balancing method where
	// a request is passed to the server with the least number of
	// active connections.
	StrategyRandom Strategy = iota

	// StrategyRoundRobin represents a random load balancing method where
	// a request is passed to the server in round-robin fashion.
	StrategyRoundRobin

	// StrategyLeastConn represents a random load balancing method where
	// a request is passed to the server with the least number of
	// active connections.
	StrategyLeastConn

	// StrategyIPHash represents a IP hash load balancing method where
	// a request is passed to the server based on the result of hashing
	// the request IP address.
	StrategyIPHash
)

var strategyName = map[Strategy]string{
	StrategyRandom:     "RANDOM",
	StrategyRoundRobin: "ROUNDROBIN",
	StrategyLeastConn:  "LEASTCONN",
	StrategyIPHash:     "IPHASH",
}

func (s Strategy) String() string {
	return strategyName[s]
}

// BalancerService is an interface for interfacing with the balancer
// targets of the Spotinst API.
type BalancerService interface {
	ListBalancers(*ListBalancersInput) (*ListBalancersOutput, error)
	CreateBalancer(*CreateBalancerInput) (*CreateBalancerOutput, error)
	ReadBalancer(*ReadBalancerInput) (*ReadBalancerOutput, error)
	UpdateBalancer(*UpdateBalancerInput) (*UpdateBalancerOutput, error)
	DeleteBalancer(*DeleteBalancerInput) (*DeleteBalancerOutput, error)

	ListListeners(*ListListenersInput) (*ListListenersOutput, error)
	CreateListener(*CreateListenerInput) (*CreateListenerOutput, error)
	ReadListener(*ReadListenerInput) (*ReadListenerOutput, error)
	UpdateListener(*UpdateListenerInput) (*UpdateListenerOutput, error)
	DeleteListener(*DeleteListenerInput) (*DeleteListenerOutput, error)

	ListRoutingRules(*ListRoutingRulesInput) (*ListRoutingRulesOutput, error)
	CreateRoutingRule(*CreateRoutingRuleInput) (*CreateRoutingRuleOutput, error)
	ReadRoutingRule(*ReadRoutingRuleInput) (*ReadRoutingRuleOutput, error)
	UpdateRoutingRule(*UpdateRoutingRuleInput) (*UpdateRoutingRuleOutput, error)
	DeleteRoutingRule(*DeleteRoutingRuleInput) (*DeleteRoutingRuleOutput, error)

	ListMiddlewares(*ListMiddlewaresInput) (*ListMiddlewaresOutput, error)
	CreateMiddleware(*CreateMiddlewareInput) (*CreateMiddlewareOutput, error)
	ReadMiddleware(*ReadMiddlewareInput) (*ReadMiddlewareOutput, error)
	UpdateMiddleware(*UpdateMiddlewareInput) (*UpdateMiddlewareOutput, error)
	DeleteMiddleware(*DeleteMiddlewareInput) (*DeleteMiddlewareOutput, error)

	ListTargetSets(*ListTargetSetsInput) (*ListTargetSetsOutput, error)
	CreateTargetSet(*CreateTargetSetInput) (*CreateTargetSetOutput, error)
	ReadTargetSet(*ReadTargetSetInput) (*ReadTargetSetOutput, error)
	UpdateTargetSet(*UpdateTargetSetInput) (*UpdateTargetSetOutput, error)
	DeleteTargetSet(*DeleteTargetSetInput) (*DeleteTargetSetOutput, error)

	ListTargets(*ListTargetsInput) (*ListTargetsOutput, error)
	CreateTarget(*CreateTargetInput) (*CreateTargetOutput, error)
	ReadTarget(*ReadTargetInput) (*ReadTargetOutput, error)
	UpdateTarget(*UpdateTargetInput) (*UpdateTargetOutput, error)
	DeleteTarget(*DeleteTargetInput) (*DeleteTargetOutput, error)

	ListRuntimes(*ListRuntimesInput) (*ListRuntimesOutput, error)
	ReadRuntime(*ReadRuntimeInput) (*ReadRuntimeOutput, error)
}

// BalancerServiceOp handles communication with the balancer related methods
// of the Spotinst API.
type BalancerServiceOp struct {
	client *Client
}

var _ BalancerService = &BalancerServiceOp{}

type Balancer struct {
	ID              *string    `json:"id,omitempty"`
	AccountID       *string    `json:"accountId,omitempty"`
	Name            *string    `json:"name,omitempty"`
	DNSRRType       *string    `json:"dnsRrType,omitempty"`
	DNSRRName       *string    `json:"dnsRrName,omitempty"`
	DNSCNAMEAliases []string   `json:"dnsCnameAliases,omitempty"`
	Timeouts        *Timeouts  `json:"timeouts,omitempty"`
	Tags            []*Tag     `json:"tags,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `json:"updatedAt,omitempty"`
}

type Timeouts struct {
	Idle     *int `json:"idle"`
	Draining *int `json:"draining"`
}

type ListBalancersInput struct {
	AccountID    *string `json:"accountId,omitempty"`
	DeploymentID *string `json:"deploymentId,omitempty"`
}

type ListBalancersOutput struct {
	Balancers []*Balancer `json:"balancers,omitempty"`
}

type CreateBalancerInput struct {
	AccountID *string   `json:"accountId,omitempty"`
	Balancer  *Balancer `json:"balancer,omitempty"`
}

type CreateBalancerOutput struct {
	Balancer *Balancer `json:"balancer,omitempty"`
}

type ReadBalancerInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	BalancerID *string `json:"balancerId,omitempty"`
}

type ReadBalancerOutput struct {
	Balancer *Balancer `json:"balancer,omitempty"`
}

type UpdateBalancerInput struct {
	AccountID *string   `json:"accountId,omitempty"`
	Balancer  *Balancer `json:"balancer,omitempty"`
}

type UpdateBalancerOutput struct{}

type DeleteBalancerInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	BalancerID *string `json:"balancerId,omitempty"`
}

type DeleteBalancerOutput struct{}

func balancerFromJSON(in []byte) (*Balancer, error) {
	b := new(Balancer)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func balancersFromJSON(in []byte) ([]*Balancer, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Balancer, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rb := range rw.Response.Items {
		b, err := balancerFromJSON(rb)
		if err != nil {
			return nil, err
		}
		out[i] = b
	}
	return out, nil
}

func balancersFromHttpResponse(resp *http.Response) ([]*Balancer, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return balancersFromJSON(body)
}

func (b *BalancerServiceOp) ListBalancers(input *ListBalancersInput) (*ListBalancersOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/balancer")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.DeploymentID != nil {
		r.params.Set("deploymentId", StringValue(input.DeploymentID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := balancersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListBalancersOutput{
		Balancers: bs,
	}, nil
}

func (b *BalancerServiceOp) CreateBalancer(input *CreateBalancerInput) (*CreateBalancerOutput, error) {
	r := b.client.newRequest("POST", "/loadBalancer/balancer")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := balancersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateBalancerOutput{
		Balancer: bs[0],
	}, nil
}

func (b *BalancerServiceOp) ReadBalancer(input *ReadBalancerInput) (*ReadBalancerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/balancer/{balancerId}", map[string]string{
		"balancerId": StringValue(input.BalancerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := balancersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadBalancerOutput{
		Balancer: bs[0],
	}, nil
}

func (b *BalancerServiceOp) UpdateBalancer(input *UpdateBalancerInput) (*UpdateBalancerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/balancer/{balancerId}", map[string]string{
		"balancerId": StringValue(input.Balancer.ID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateBalancerOutput{}, nil
}

func (b *BalancerServiceOp) DeleteBalancer(input *DeleteBalancerInput) (*DeleteBalancerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/balancer/{balancerId}", map[string]string{
		"balancerId": StringValue(input.BalancerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteBalancerOutput{}, nil
}

type Listener struct {
	ID         *string    `json:"id,omitempty"`
	BalancerID *string    `json:"balancerId,omitempty"`
	Protocol   *string    `json:"protocol,omitempty"`
	Port       *int       `json:"port,omitempty"`
	TLSConfig  *TLSConfig `json:"tlsConfig,omitempty"`
	Tags       []*Tag     `json:"tags,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

type ListListenersInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListListenersOutput struct {
	Listeners []*Listener `json:"listeners,omitempty"`
}

type CreateListenerInput struct {
	AccountID *string   `json:"accountId,omitempty"`
	Listener  *Listener `json:"listener,omitempty"`
}

type CreateListenerOutput struct {
	Listener *Listener `json:"listener,omitempty"`
}

type ReadListenerInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	ListenerID *string `json:"listenerId,omitempty"`
}

type ReadListenerOutput struct {
	Listener *Listener `json:"listener,omitempty"`
}

type UpdateListenerInput struct {
	AccountID *string   `json:"accountId,omitempty"`
	Listener  *Listener `json:"listener,omitempty"`
}

type UpdateListenerOutput struct{}

type DeleteListenerInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	ListenerID *string `json:"listenerId,omitempty"`
}

type DeleteListenerOutput struct{}

func listenerFromJSON(in []byte) (*Listener, error) {
	b := new(Listener)
	if err := json.Unmarshal(in, b); err != nil {
		return nil, err
	}
	return b, nil
}

func listenersFromJSON(in []byte) ([]*Listener, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Listener, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rl := range rw.Response.Items {
		l, err := listenerFromJSON(rl)
		if err != nil {
			return nil, err
		}
		out[i] = l
	}
	return out, nil
}

func listenersFromHttpResponse(resp *http.Response) ([]*Listener, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return listenersFromJSON(body)
}

func (b *BalancerServiceOp) ListListeners(input *ListListenersInput) (*ListListenersOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/listener")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls, err := listenersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListListenersOutput{
		Listeners: ls,
	}, nil
}

func (b *BalancerServiceOp) CreateListener(input *CreateListenerInput) (*CreateListenerOutput, error) {
	r := b.client.newRequest("POST", "/loadBalancer/listener")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls, err := listenersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateListenerOutput{
		Listener: ls[0],
	}, nil
}

func (b *BalancerServiceOp) ReadListener(input *ReadListenerInput) (*ReadListenerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/listener/{listenerId}", map[string]string{
		"listenerId": StringValue(input.ListenerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ls, err := listenersFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadListenerOutput{
		Listener: ls[0],
	}, nil
}

func (b *BalancerServiceOp) UpdateListener(input *UpdateListenerInput) (*UpdateListenerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/listener/{listenerId}", map[string]string{
		"listenerId": StringValue(input.Listener.ID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateListenerOutput{}, nil
}

func (b *BalancerServiceOp) DeleteListener(input *DeleteListenerInput) (*DeleteListenerOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/listener/{listenerId}", map[string]string{
		"listenerId": StringValue(input.ListenerID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteListenerOutput{}, nil
}

type RoutingRule struct {
	ID            *string    `json:"id,omitempty"`
	BalancerID    *string    `json:"balancerId,omitempty"`
	ListenerID    *string    `json:"listenerId,omitempty"`
	MiddlewareIDs []string   `json:"middlewareIds,omitempty"`
	TargetSetIDs  []string   `json:"targetSetIds,omitempty"`
	Priority      *int       `json:"priority,omitempty"`
	Strategy      *string    `json:"strategy,omitempty"`
	Route         *string    `json:"route,omitempty"`
	Tags          []*Tag     `json:"tags,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

type ListRoutingRulesInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListRoutingRulesOutput struct {
	RoutingRules []*RoutingRule `json:"routingRules,omitempty"`
}

type CreateRoutingRuleInput struct {
	AccountID   *string      `json:"accountId,omitempty"`
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type CreateRoutingRuleOutput struct {
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type ReadRoutingRuleInput struct {
	AccountID     *string `json:"accountId,omitempty"`
	RoutingRuleID *string `json:"routingRuleId,omitempty"`
}

type ReadRoutingRuleOutput struct {
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type UpdateRoutingRuleInput struct {
	AccountID   *string      `json:"accountId,omitempty"`
	RoutingRule *RoutingRule `json:"routingRule,omitempty"`
}

type UpdateRoutingRuleOutput struct{}

type DeleteRoutingRuleInput struct {
	AccountID     *string `json:"accountId,omitempty"`
	RoutingRuleID *string `json:"routingRuleId,omitempty"`
}

type DeleteRoutingRuleOutput struct{}

func routingRuleFromJSON(in []byte) (*RoutingRule, error) {
	rr := new(RoutingRule)
	if err := json.Unmarshal(in, rr); err != nil {
		return nil, err
	}
	return rr, nil
}

func routingRulesFromJSON(in []byte) ([]*RoutingRule, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*RoutingRule, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rr := range rw.Response.Items {
		r, err := routingRuleFromJSON(rr)
		if err != nil {
			return nil, err
		}
		out[i] = r
	}
	return out, nil
}

func routingRulesFromHttpResponse(resp *http.Response) ([]*RoutingRule, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return routingRulesFromJSON(body)
}

func (b *BalancerServiceOp) ListRoutingRules(input *ListRoutingRulesInput) (*ListRoutingRulesOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/routingRule")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rr, err := routingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListRoutingRulesOutput{
		RoutingRules: rr,
	}, nil
}

func (b *BalancerServiceOp) CreateRoutingRule(input *CreateRoutingRuleInput) (*CreateRoutingRuleOutput, error) {
	r := b.client.newRequest("POST", "/loadBalancer/routingRule")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rr, err := routingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateRoutingRuleOutput{
		RoutingRule: rr[0],
	}, nil
}

func (b *BalancerServiceOp) ReadRoutingRule(input *ReadRoutingRuleInput) (*ReadRoutingRuleOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/routingRule/{routingRuleId}", map[string]string{
		"routingRuleId": StringValue(input.RoutingRuleID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rr, err := routingRulesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadRoutingRuleOutput{
		RoutingRule: rr[0],
	}, nil
}

func (b *BalancerServiceOp) UpdateRoutingRule(input *UpdateRoutingRuleInput) (*UpdateRoutingRuleOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/routingRule/{routingRuleId}", map[string]string{
		"routingRuleId": StringValue(input.RoutingRule.ID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateRoutingRuleOutput{}, nil
}

func (b *BalancerServiceOp) DeleteRoutingRule(input *DeleteRoutingRuleInput) (*DeleteRoutingRuleOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/routingRule/{routingRuleId}", map[string]string{
		"routingRuleId": StringValue(input.RoutingRuleID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteRoutingRuleOutput{}, nil
}

type Middleware struct {
	ID         *string         `json:"id,omitempty"`
	AccountID  *string         `json:"accountId,omitempty"`
	BalancerID *string         `json:"balancerId,omitempty"`
	Type       *string         `json:"type,omitempty"`
	Priority   *int            `json:"priority,omitempty"`
	Spec       json.RawMessage `json:"spec,omitempty"`
	Tags       []*Tag          `json:"tags,omitempty"`
	CreatedAt  *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time      `json:"updatedAt,omitempty"`
}

type ListMiddlewaresInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListMiddlewaresOutput struct {
	Middlewares []*Middleware `json:"middlewares,omitempty"`
}

type CreateMiddlewareInput struct {
	AccountID  *string     `json:"accountId,omitempty"`
	Middleware *Middleware `json:"middleware,omitempty"`
}

type CreateMiddlewareOutput struct {
	Middleware *Middleware `json:"middleware,omitempty"`
}

type ReadMiddlewareInput struct {
	AccountID    *string `json:"accountId,omitempty"`
	MiddlewareID *string `json:"middlewareId,omitempty"`
}

type ReadMiddlewareOutput struct {
	Middleware *Middleware `json:"middleware,omitempty"`
}

type UpdateMiddlewareInput struct {
	AccountID  *string     `json:"accountId,omitempty"`
	Middleware *Middleware `json:"middleware,omitempty"`
}

type UpdateMiddlewareOutput struct{}

type DeleteMiddlewareInput struct {
	AccountID    *string `json:"accountId,omitempty"`
	MiddlewareID *string `json:"middlewareId,omitempty"`
}

type DeleteMiddlewareOutput struct{}

func middlewareFromJSON(in []byte) (*Middleware, error) {
	m := new(Middleware)
	if err := json.Unmarshal(in, m); err != nil {
		return nil, err
	}
	return m, nil
}

func middlewaresFromJSON(in []byte) ([]*Middleware, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Middleware, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rm := range rw.Response.Items {
		m, err := middlewareFromJSON(rm)
		if err != nil {
			return nil, err
		}
		out[i] = m
	}
	return out, nil
}

func middlewaresFromHttpResponse(resp *http.Response) ([]*Middleware, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return middlewaresFromJSON(body)
}

func (b *BalancerServiceOp) ListMiddlewares(input *ListMiddlewaresInput) (*ListMiddlewaresOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/middleware")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ms, err := middlewaresFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListMiddlewaresOutput{
		Middlewares: ms,
	}, nil
}

func (b *BalancerServiceOp) CreateMiddleware(input *CreateMiddlewareInput) (*CreateMiddlewareOutput, error) {
	r := b.client.newRequest("POST", "/loadBalancer/middleware")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ms, err := middlewaresFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateMiddlewareOutput{
		Middleware: ms[0],
	}, nil
}

func (b *BalancerServiceOp) ReadMiddleware(input *ReadMiddlewareInput) (*ReadMiddlewareOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/middleware/{middlewareId}", map[string]string{
		"middlewareId": StringValue(input.MiddlewareID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ms, err := middlewaresFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadMiddlewareOutput{
		Middleware: ms[0],
	}, nil
}

func (b *BalancerServiceOp) UpdateMiddleware(input *UpdateMiddlewareInput) (*UpdateMiddlewareOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/middleware/{middlewareId}", map[string]string{
		"middlewareId": StringValue(input.Middleware.ID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateMiddlewareOutput{}, nil
}

func (b *BalancerServiceOp) DeleteMiddleware(input *DeleteMiddlewareInput) (*DeleteMiddlewareOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/middleware/{middlewareId}", map[string]string{
		"middlewareId": StringValue(input.MiddlewareID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteMiddlewareOutput{}, nil
}

type TargetSet struct {
	ID           *string      `json:"id,omitempty"`
	BalancerID   *string      `json:"balancerId,omitempty"`
	DeploymentID *string      `json:"deploymentId,omitempty"`
	Name         *string      `json:"name,omitempty"`
	Protocol     *string      `json:"protocol,omitempty"`
	Port         *int         `json:"port,omitempty"`
	Weight       *int         `json:"weight,omitempty"`
	HealthCheck  *HealthCheck `json:"healthCheck,omitempty"`
	Tags         []*Tag       `json:"tags,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time   `json:"updatedAt,omitempty"`
}

type HealthCheck struct {
	Path                    *string `json:"path,omitempty"`
	Port                    *int    `json:"port,omitempty"`
	Protocol                *string `json:"protocol,omitempty"`
	Timeout                 *int    `json:"timeout,omitempty"`
	Interval                *int    `json:"interval,omitempty"`
	HealthyThresholdCount   *int    `json:"healthyThresholdCount,omitempty"`
	UnhealthyThresholdCount *int    `json:"unhealthyThresholdCount,omitempty"`
}

type ListTargetSetsInput struct {
	AccountID  *string `json:"accountId,omitempty"`
	BalancerID *string `json:"balancerId,omitempty"`
}

type ListTargetSetsOutput struct {
	TargetSets []*TargetSet `json:"targetSets,omitempty"`
}

type CreateTargetSetInput struct {
	AccountID *string    `json:"accountId,omitempty"`
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type CreateTargetSetOutput struct {
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type ReadTargetSetInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
}

type ReadTargetSetOutput struct {
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type UpdateTargetSetInput struct {
	AccountID *string    `json:"accountId,omitempty"`
	TargetSet *TargetSet `json:"targetSet,omitempty"`
}

type UpdateTargetSetOutput struct{}

type DeleteTargetSetInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
}

type DeleteTargetSetOutput struct{}

func targetSetFromJSON(in []byte) (*TargetSet, error) {
	ts := new(TargetSet)
	if err := json.Unmarshal(in, ts); err != nil {
		return nil, err
	}
	return ts, nil
}

func targetSetsFromJSON(in []byte) ([]*TargetSet, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*TargetSet, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rts := range rw.Response.Items {
		ts, err := targetSetFromJSON(rts)
		if err != nil {
			return nil, err
		}
		out[i] = ts
	}
	return out, nil
}

func targetSetsFromHttpResponse(resp *http.Response) ([]*TargetSet, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return targetSetsFromJSON(body)
}

func (b *BalancerServiceOp) ListTargetSets(input *ListTargetSetsInput) (*ListTargetSetsOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/targetSet")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetSetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListTargetSetsOutput{
		TargetSets: ts,
	}, nil
}

func (b *BalancerServiceOp) CreateTargetSet(input *CreateTargetSetInput) (*CreateTargetSetOutput, error) {
	r := b.client.newRequest("POST", "/loadBalancer/targetSet")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetSetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateTargetSetOutput{
		TargetSet: ts[0],
	}, nil
}

func (b *BalancerServiceOp) ReadTargetSet(input *ReadTargetSetInput) (*ReadTargetSetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/targetSet/{targetSetId}", map[string]string{
		"targetSetId": StringValue(input.TargetSetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetSetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadTargetSetOutput{
		TargetSet: ts[0],
	}, nil
}

func (b *BalancerServiceOp) UpdateTargetSet(input *UpdateTargetSetInput) (*UpdateTargetSetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/targetSet/{targetSetId}", map[string]string{
		"targetSetId": StringValue(input.TargetSet.ID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateTargetSetOutput{}, nil
}

func (b *BalancerServiceOp) DeleteTargetSet(input *DeleteTargetSetInput) (*DeleteTargetSetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/targetSet/{targetSetId}", map[string]string{
		"targetSetId": StringValue(input.TargetSetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteTargetSetOutput{}, nil
}

type Target struct {
	ID          *string    `json:"id,omitempty"`
	AccountID   *string    `json:"accountId,omitempty"`
	BalancerID  *string    `json:"balancerId,omitempty"`
	TargetSetID *string    `json:"targetSetId,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Host        *string    `json:"host,omitempty"`
	Port        *int       `json:"port,omitempty"`
	Weight      *int       `json:"weight,omitempty"`
	Status      *Status    `json:"status,omitempty"`
	Tags        []*Tag     `json:"tags,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type Status struct {
	Readiness   *string `json:"readiness"`
	Healthiness *string `json:"healthiness"`
}

type ListTargetsInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	BalancerID  *string `json:"balancerId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
}

type ListTargetsOutput struct {
	Targets []*Target `json:"targets,omitempty"`
}

type CreateTargetInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
	Target      *Target `json:"target,omitempty"`
}

type CreateTargetOutput struct {
	Target *Target `json:"target,omitempty"`
}

type ReadTargetInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
	TargetID    *string `json:"targetId,omitempty"`
}

type ReadTargetOutput struct {
	Target *Target `json:"target,omitempty"`
}

type UpdateTargetInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
	Target      *Target `json:"target,omitempty"`
}

type UpdateTargetOutput struct{}

type DeleteTargetInput struct {
	AccountID   *string `json:"accountId,omitempty"`
	TargetSetID *string `json:"targetSetId,omitempty"`
	TargetID    *string `json:"targetId,omitempty"`
}

type DeleteTargetOutput struct{}

func targetFromJSON(in []byte) (*Target, error) {
	t := new(Target)
	if err := json.Unmarshal(in, t); err != nil {
		return nil, err
	}
	return t, nil
}

func targetsFromJSON(in []byte) ([]*Target, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Target, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rt := range rw.Response.Items {
		t, err := targetFromJSON(rt)
		if err != nil {
			return nil, err
		}
		out[i] = t
	}
	return out, nil
}

func targetsFromHttpResponse(resp *http.Response) ([]*Target, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return targetsFromJSON(body)
}

func (b *BalancerServiceOp) ListTargets(input *ListTargetsInput) (*ListTargetsOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/target")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.BalancerID != nil {
		r.params.Set("balancerId", StringValue(input.BalancerID))
	}

	if input.TargetSetID != nil {
		r.params.Set("targetSetId", StringValue(input.TargetSetID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListTargetsOutput{
		Targets: ts,
	}, nil
}

func (b *BalancerServiceOp) CreateTarget(input *CreateTargetInput) (*CreateTargetOutput, error) {
	r := b.client.newRequest("POST", "/loadBalancer/target")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &CreateTargetOutput{
		Target: ts[0],
	}, nil
}

func (b *BalancerServiceOp) ReadTarget(input *ReadTargetInput) (*ReadTargetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/target/{targetId}", map[string]string{
		"targetId": StringValue(input.TargetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ts, err := targetsFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadTargetOutput{
		Target: ts[0],
	}, nil
}

func (b *BalancerServiceOp) UpdateTarget(input *UpdateTargetInput) (*UpdateTargetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/target/{targetId}", map[string]string{
		"targetId": StringValue(input.Target.ID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("PUT", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &UpdateTargetOutput{}, nil
}

func (b *BalancerServiceOp) DeleteTarget(input *DeleteTargetInput) (*DeleteTargetOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/target/{targetId}", map[string]string{
		"targetId": StringValue(input.TargetID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("DELETE", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &DeleteTargetOutput{}, nil
}

type Runtime struct {
	ID             *string    `json:"id,omitempty"`
	AccountID      *string    `json:"accountId,omitempty"`
	DeploymentID   *string    `json:"deploymentId,omitempty"`
	IPAddr         *string    `json:"ip,omitempty"`
	Version        *string    `json:"version,omitempty"`
	Status         *Status    `json:"status,omitempty"`
	LastReportedAt *time.Time `json:"lastReported,omitempty"`
	Leader         *bool      `json:"isLeader,omitempty"`
	Tags           []*Tag     `json:"tags,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

type ListRuntimesInput struct {
	AccountID    *string `json:"accountId,omitempty"`
	DeploymentID *string `json:"deploymentId,omitempty"`
}

type ListRuntimesOutput struct {
	Runtimes []*Runtime `json:"runtimes,omitempty"`
}

type ReadRuntimeInput struct {
	AccountID *string `json:"accountId,omitempty"`
	RuntimeID *string `json:"runtimeId,omitempty"`
}

type ReadRuntimeOutput struct {
	Runtime *Runtime `json:"runtime,omitempty"`
}

func runtimeFromJSON(in []byte) (*Runtime, error) {
	rt := new(Runtime)
	if err := json.Unmarshal(in, rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func runtimesFromJSON(in []byte) ([]*Runtime, error) {
	var rw responseWrapper
	if err := json.Unmarshal(in, &rw); err != nil {
		return nil, err
	}
	out := make([]*Runtime, len(rw.Response.Items))
	if len(out) == 0 {
		return out, nil
	}
	for i, rrt := range rw.Response.Items {
		rt, err := runtimeFromJSON(rrt)
		if err != nil {
			return nil, err
		}
		out[i] = rt
	}
	return out, nil
}

func runtimesFromHttpResponse(resp *http.Response) ([]*Runtime, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return runtimesFromJSON(body)
}

func (b *BalancerServiceOp) ListRuntimes(input *ListRuntimesInput) (*ListRuntimesOutput, error) {
	r := b.client.newRequest("GET", "/loadBalancer/runtime")
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	if input.DeploymentID != nil {
		r.params.Set("deploymentId", StringValue(input.DeploymentID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rts, err := runtimesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ListRuntimesOutput{
		Runtimes: rts,
	}, nil
}

func (b *BalancerServiceOp) ReadRuntime(input *ReadRuntimeInput) (*ReadRuntimeOutput, error) {
	path, err := uritemplates.Expand("/loadBalancer/runtime/{runtimeId}", map[string]string{
		"runtimeId": StringValue(input.RuntimeID),
	})
	if err != nil {
		return nil, err
	}

	r := b.client.newRequest("GET", path)
	r.obj = input

	if input.AccountID != nil {
		r.params.Set("accountId", StringValue(input.AccountID))
	}

	_, resp, err := requireOK(b.client.doRequest(r))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rt, err := runtimesFromHttpResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ReadRuntimeOutput{
		Runtime: rt[0],
	}, nil
}
