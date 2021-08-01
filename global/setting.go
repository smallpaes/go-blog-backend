package global

import (
	"github.com/smallpaes/go-blog-backend/pkg/setting"
)

// used to save setting config globally
// can be access al over the app
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
