package models
/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: models.go
 => Version: X.X.X.X
 => Author: Jonathan McAllister
 => Purpose: 
    Model top level. Handles database constructs of created_on
    created_by and date time etc. wthin the database

*/


// Imports
// ----------------------------------
import (
	
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"microserviceframework/config"
	"time"
)


// Global Variables
// ----------------------------------
var db *gorm.DB



// Model Struct [inherited]
// ----------------------------------
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}


// FUNC: Setup()
// ----------------------------------
func Setup() {
	var err error
	db, err = gorm.Open(config.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}


// FUNC: CloseDB
// ----------------------------------
func CloseDB() {

	defer db.Close()

}


// FUNC: UpdateTimeStampForCreateCallback
// ----------------------------------
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// FUNC: updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
// ----------------------------------
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// FUNC: deleteCallback will set `DeletedOn` where deleting
// ----------------------------------
func deleteCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}

}

// FUNC: addExtraSpaceIfExist adds a separator
// ----------------------------------
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
