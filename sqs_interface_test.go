package sqsconsumer_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
	"github.com/tprovoost/sqsconsumer"
)

func TestSQSInterfaceImplementsSQSAPI(t *testing.T) {
	assert.Implements(t, (*sqsconsumer.SQSAPI)(nil), sqs.New(session.New()))
}
