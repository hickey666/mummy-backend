package environment

import (
	"log"
	"mummy-storage/config"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config  *config.Config
	DbEngin *gorm.DB
	Domain  *DomainEnv
	Svc     *SvcEnv
}

func NewServiceContext(c *config.Config) *ServiceContext {
	sc := &ServiceContext{Config: c}
	sc.DbEngin = initMysql(c)
	sc.Domain = NewDomainEnv(sc.DbEngin)
	return sc
}

type loggerSQL struct {
}

func (l *loggerSQL) Printf(message string, v ...interface{}) {
	logx.Slowf(message, v...)
}

func initMysql(c *config.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second * 2, // 慢 SQL 阈值
			LogLevel:                  logger.Warn,     // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
		},
	)
	db, err := gorm.Open(mysql.Open(c.Mysql), &gorm.Config{
		Logger: newLogger,
	})
	// 如果出错就GameOver了
	if err != nil {
		logx.Errorw("gorm.Open", logx.Field("mySqlConn", c.Mysql), logx.Field("err", err))
		panic("init mysql error")
	}
	return db
}
