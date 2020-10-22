package sqsconsumer_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"

	sqsconsumer "github.com/tprovoost/sqsconsumer_light"
)

func TestSQSInterfaceImplementsSQSAPI(t *testing.T) {
	assert.Implements(t, (*sqsconsumer.SQSAPI)(nil), sqs.New(session.New()))
}
