package models

import (
	"fmt"
	"service-notification/pkg/logging"
	"service-notification/pkg/setting"
	"service-notification/pkg/util"

	//"service-notification/pkg/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var baseDb *gorm.DB
//var baseDb1 *gorm.baseDb
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	IsDeleted int `json:"is_deleted"`
	UpdatedAt int `json:"updated_at"`
	CreatedAt int `json:"created_at"`
}


func Setup() {
	var err error
	baseDb, err = gorm.Open(setting.DbSetting.Type, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DbSetting.User,
		setting.DbSetting.Password,
		setting.DbSetting.Host,
		setting.DbSetting.Name))
	if err != nil {
		logging.Errorf("models.Setup err:%v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DbSetting.TablePrefix + defaultTableName
	}
	baseDb.SingularTable(true)
	baseDb.Callback().Create().Before("gorm:create").Register("update_created_at", updateTimeForCreate)
	baseDb.Callback().Update(). Before("gorm:update").Register("before_update", updateTimeForUpdate)
	baseDb.Callback().Delete().Register("gorm:delete", softDelForDelete)
	//if setting.Env=="dev" {
	baseDb.LogMode(true)
//	}
	baseDb.DB().SetMaxOpenConns(150)
	baseDb.DB().SetMaxIdleConns(100)
	baseDb.SetLogger(logging.DbLogger)
}
func GetDb() *gorm.DB {
	return baseDb
}
func updateTimeForCreate(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	nowTime := time.Now().Unix()
	scope.SetColumn("UpdatedAt", nowTime)
	scope.SetColumn("CreatedAt", nowTime)
}
func updateTimeForUpdate(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	scope.SetColumn("UpdatedAt", time.Now().Unix())
}

func softDelForDelete(scope *gorm.Scope)  {
	if scope.HasError() {
		return
	}
	scope.Raw(fmt.Sprintf(
		"UPDATE %v SET is_deleted=1,updated=%v %v",
		scope.QuotedTableName(),
		time.Now().Unix(),
		util.AddExtraSpaceIfExist(scope.CombinedConditionSql()),
	)).Exec()
}

func CloseDB(){
	defer baseDb.Close()
}
