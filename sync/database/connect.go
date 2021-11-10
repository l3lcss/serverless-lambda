package database

import (
	"fmt"
	"myservice/sync/appconf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectMySQL func
func ConnectMySQL(mySQL *appconf.MySQL) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&&parseTime=true",
		mySQL.Username,
		mySQL.Password,
		mySQL.Host,
		mySQL.Port,
		mySQL.Database,
	)

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	return dbConn
}

// ConnectMSSQL func
func ConnectMSSQL(MSSQL *appconf.MSSQL) *gorm.DB {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		MSSQL.Username,
		MSSQL.Password,
		MSSQL.Host,
		MSSQL.Port,
		MSSQL.Database,
	)

	dbConn, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	return dbConn
}
