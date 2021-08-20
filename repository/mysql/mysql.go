package mysql

import (
	"fmt"

	"github.com/macoli/redisplat/models"
	"go.uber.org/zap"

	"github.com/macoli/redisplat/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// connet msyql,也可以使用 MustConnet,连接不成功就 panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect mysql failed, err:%v\n", err)
		return
	}
	// 设置连接参数
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	return
}

func Close() {
	_ = db.Close()
}

// GetRedisAppInfo 从 CacheCloud 的 mysql 数据库中查询 redis app 信息
func GetRedisAppInfo(ip string, port int) (data *models.RespRedisApp, err error) {
	data = new(models.RespRedisApp)
	sqlStr := `select i.status as instance_status, 
a.app_id, a.name as app_name, a.status as app_status, a.type as app_type 
from instance_info i 
left join app_desc a 
on i.app_id=a.app_id where ip=? and port=?`

	err = db.Get(data, sqlStr, ip, port)
	if err != nil {
		zap.L().Error("db.Get(data, sqlStr, ip, port) failed", zap.Error(err))
		return
	}

	zap.L().Debug("mysql.GetRedisAppInfo(ipStr, portInt)",
		zap.Int("ret.ID", data.ID),
		zap.String("ret.Name", data.Name))

	return
}
