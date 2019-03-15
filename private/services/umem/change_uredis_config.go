//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem ChangeURedisConfig

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// ChangeURedisConfigRequest is request schema for ChangeURedisConfig action
type ChangeURedisConfigRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 资源ID
	GroupId *string `required:"true"`

	// 配置文件ID
	ConfigId *string `required:"true"`
}

// ChangeURedisConfigResponse is response schema for ChangeURedisConfig action
type ChangeURedisConfigResponse struct {
	response.CommonBase

	// 返回码
	RetCode int

	// 操作名称
	Action string
}

// NewChangeURedisConfigRequest will create request of ChangeURedisConfig action.
func (c *UMemClient) NewChangeURedisConfigRequest() *ChangeURedisConfigRequest {
	req := &ChangeURedisConfigRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// ChangeURedisConfig - 更换Redis配置文件
func (c *UMemClient) ChangeURedisConfig(req *ChangeURedisConfigRequest) (*ChangeURedisConfigResponse, error) {
	var err error
	var res ChangeURedisConfigResponse

	err = c.client.InvokeAction("ChangeURedisConfig", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
