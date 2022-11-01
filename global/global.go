package global

import (
	"github.com/lqf1215/kitex-demo/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.ClientConfig
)
