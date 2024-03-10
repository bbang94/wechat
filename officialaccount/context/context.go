package context

import (
	"github.com/bbang94/wechat/v2/credential"
	"github.com/bbang94/wechat/v2/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
