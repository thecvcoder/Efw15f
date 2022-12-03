package initialize

import (
	"context"
	"fmt"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		global.CONFIG.Mysql.Username,
		global.CONFIG.Mysql.Password,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.Database,
		global.CONFIG.Mysql.Charset,
		global.CONFIG.Mysql.Collation,
		global.CONFIG.Mysql.Query,
	)
	// 隐藏mysql的密码，记录到日志库
	showDsn := fmt.Sprintf("%s:******@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		global.CONFIG.Mysql.Username,
		global.CONFIG.Mysql.Host,
		global.CONFIG.Mysql.Port,
		global.CONFIG.Mysql.Database,
		global.CONFIG.Mysql.Charset,
		global.CONFIG.Mysql.Collation,
		global.CONFIG.Mysql.Query,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name 用于连接mysql
		DefaultStringSize:         191,   // string类型字段默认长度
		SkipInitializeWithVersion: false, // 根据mysql版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction:                   true, // 跳过默认事务
		DisableForeignKeyConstraintWhenMigrating: true, // 约束外键
	})

	if err != nil {
		global.LOG.Panicf("初始化数据库异常: %v", err)
		panic(fmt.Errorf("初始化数据库异常: %v", err))
	}

	// 开启mysql日志，在config.yml文件中开启或关闭
	if global.CONFIG.Mysql.LogMode {
		db.Debug()
	}

	_db = db
	dbAutoMigrate() // 自动迁移表结构
	global.LOG.Info("初始化mysql数据库完成! dsn:", showDsn)

	return _db
}

func dbAutoMigrate() {
	err := _db.AutoMigrate(
		&model.AdminUser{},
	)
	if err != nil {
		global.LOG.Error("初始化数据失败: ", zap.Error(err))
	}
	global.LOG.Info("初始化数据成功!")
}

func NewDbClient(c context.Context) *gorm.DB {
	db := _db
	return db.WithContext(c)
}
