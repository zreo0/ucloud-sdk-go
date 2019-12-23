// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// LoginByPasswordRequest is request schema for LoginByPassword action
type LoginByPasswordRequest struct {
	request.CommonBase

	// 密码
	Password *string `required:"true"`

	// 邮箱
	UserEmail *string `required:"true"`
}

// LoginByPasswordResponse is response schema for LoginByPassword action
type LoginByPasswordResponse struct {
	response.CommonBase

	//
	Message string
}

// NewLoginByPasswordRequest will create request of LoginByPassword action.
func (c *UCloudStackClient) NewLoginByPasswordRequest() *LoginByPasswordRequest {
	req := &LoginByPasswordRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// LoginByPassword - 登录账户
func (c *UCloudStackClient) LoginByPassword(req *LoginByPasswordRequest) (*LoginByPasswordResponse, error) {
	var err error
	var res LoginByPasswordResponse

	reqCopier := *req

	err = c.Client.InvokeAction("LoginByPassword", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}