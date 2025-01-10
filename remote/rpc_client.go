package remote

import (
	"reflect"

	"github.com/vgough/sequin"
	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
	"github.com/vgough/sequin/gen/sequin/v1/sequinv1connect"
	"github.com/vgough/sequin/registry"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"connectrpc.com/connect"
)

type Client struct {
	client sequinv1connect.SequinServiceClient
}

var _ sequin.Runtime = &Client{}

func NewClient(client sequinv1connect.SequinServiceClient) *Client {
	return &Client{client: client}
}

func (c *Client) Exec(ep *registry.Endpoint, args []reflect.Value) []reflect.Value {
	ctx := ep.GetContext(args)
	req := &sequinv1.ExecRequest{}
	stream, err := c.client.Exec(ctx, connect.NewRequest(req))
	if err != nil {
		return ep.MakeError(err)
	}

	for stream.Receive() {
		msg := stream.Msg()
		if msg.Done {
			return c.parseResult(ep, msg.GetResult())
		}
	}
	return ep.MakeError(stream.Err())
}

func (c *Client) parseResult(ep *registry.Endpoint, result any) []reflect.Value {
	switch d := result.(type) {
	case *longrunningpb.Operation_Response:
		var out sequinv1.OperationResult
		err := anypb.UnmarshalTo(d.Response, &out, proto.UnmarshalOptions{})
		if err != nil {
			return ep.MakeError(err)
		}
		return c.makeResult(ep, &out)
	case *longrunningpb.Operation_Error:
		s := status.FromProto(d.Error)
		return ep.MakeError(s.Err())
	}
	return nil
}

func (c *Client) makeResult(ep *registry.Endpoint, out *sequinv1.OperationResult) []reflect.Value {
	// TODO: codec
	return nil
}
