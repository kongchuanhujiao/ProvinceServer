package account

import (
	"errors"

	"github.com/kongchuanhujiao/server/internal/app/client"
	"github.com/kongchuanhujiao/server/internal/app/client/message"
	"github.com/kongchuanhujiao/server/internal/app/datahub/pkg/account"

	"go.uber.org/zap"
)

// sendCode 发送验证码
func sendCode(id string) (err error) {

	a, err := account.SelectAccount(id, 0)
	if err != nil {
		zap.L().Error("发送验证码失败", zap.Error(err))
		return
	}
	if len(a) == 0 {
		return errors.New("账号不存在")
	}

	client.GetClient().SendMessage(
		message.NewTextMessage("您的验证码是：" + account.GenerateCode(id) + "，请勿泄露给他人。有效期5分钟").
			SetTarget(&message.Target{ID: a[0].QQ}),
	)

	return
}
