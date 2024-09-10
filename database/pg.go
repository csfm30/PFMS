package database

import (
	"context"
	"fmt"

	"time"

	"pfms/logs"
	modelsPg "pfms/models/pg"

	"gorm.io/driver/postgres"
	"gorm.io/plugin/dbresolver"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc() // ตัว fc จะทำการแสดงผลของตัวคำสั่งใน GORM ออกมาเป็น Query
	fmt.Printf("%v\n--------------------------------------\n", sql)
}

func InitDatabase() {
	logs.Info("Init Database")
	fmt.Println("host", viper.GetString("pg.host"))
	fmt.Println("host", viper.GetString("pg.host2"))
	fmt.Println("host", viper.GetString("pg.host3"))
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok", viper.GetString("pg.host"), viper.GetString("pg.username"), viper.GetString("pg.password"), viper.GetString("pg.name"), viper.GetString("pg.port"))
	dsn2 := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok", viper.GetString("pg.host2"), viper.GetString("pg.username"), viper.GetString("pg.password"), viper.GetString("pg.name"), viper.GetString("pg.port"))
	dsn3 := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok", viper.GetString("pg.host3"), viper.GetString("pg.username"), viper.GetString("pg.password"), viper.GetString("pg.name"), viper.GetString("pg.port"))

	//DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Silent),
	//	// Logger: &SqlLogger{},
	//})

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		//Logger: &SqlLogger{},
	})
	DBConn.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources:  []gorm.Dialector{postgres.Open(dsn)},
		Replicas: []gorm.Dialector{postgres.Open(dsn2), postgres.Open(dsn3)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
		// print sources/replicas mode in logger
		TraceResolverMode: true,
	}))

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	fmt.Println("Database connection successfully")

	if err := DBConn.AutoMigrate(
		&modelsPg.Account{},
		&modelsPg.User{},
		&modelsPg.IncomeSource{},
		&modelsPg.ExpenseCategory{},
		&modelsPg.Transaction{},
		&modelsPg.Saving{},
	); err != nil {
		log.Error(err)
	}
}
