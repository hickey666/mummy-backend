package config

import (
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/rest"
)

const ProjectName = "share-profit"

var Global *Config

type Config struct {
	Mode    string `json:",default=pro,options=dev|test|rt|pre|pro"`
	LogConf logx.LogConf
	Mysql   string `json:",optional"`
	API     rest.RestConf
}
