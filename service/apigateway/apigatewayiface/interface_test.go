// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package apigatewayiface_test

import (
	"testing"

	"github.com/upstartmobile/aws-sdk-go/service/apigateway"
	"github.com/upstartmobile/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*apigatewayiface.APIGatewayAPI)(nil), apigateway.New(nil))
}
