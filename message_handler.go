package sqsconsumer

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"golang.org/x/net/context"
)

// MessageHandlerFunc is the interface that users of this library should implement. It will be called once per message and should return an error if there was a problem processing the message. Note that Consumer ignores the error, but it is necessary for some middleware to know whether handling was successful or not.
type MessageHandlerFunc func(ctx context.Context, msg *sqs.Message) error
