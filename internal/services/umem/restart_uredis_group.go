//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem RestartURedisGroup

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// RestartURedisGroupRequest is request schema for RestartURedisGroup action
type RestartURedisGroupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 组的ID
	GroupId *string `required:"true"`
}

// RestartURedisGroupResponse is response schema for RestartURedisGroup action
type RestartURedisGroupResponse struct {
	response.CommonBase
}

// NewRestartURedisGroupRequest will create request of RestartURedisGroup action.
func (c *UMemClient) NewRestartURedisGroupRequest() *RestartURedisGroupRequest {
	req := &RestartURedisGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// RestartURedisGroup - 重启主备实例
func (c *UMemClient) RestartURedisGroup(req *RestartURedisGroupRequest) (*RestartURedisGroupResponse, error) {
	var err error
	var res RestartURedisGroupResponse

	err = c.client.InvokeAction("RestartURedisGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}