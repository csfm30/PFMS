package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"pfms/database"
	"pfms/logs"
	"pfms/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/spf13/viper"
)

func main() {

	initConfig()
	initTimeZone()

	if err := database.InitCaching(); err != nil {
		logs.Error("Init caching fail:" + err.Error())
	}
	database.InitDatabase()

	logs.Info("This is ENV " + os.Getenv("ENV"))

	engine := html.New("./views", ".html")
	engine.Reload(true)    // Optional. Default: false
	engine.Debug(false)    // Optional. Default: false
	engine.Layout("embed") // Optional. Default: "embed"
	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		BodyLimit: 300 * 1024 * 1024,
		Views:     engine,
	})
	app.Use(cors.New())

	if os.Getenv("ENV") == "uat" || os.Getenv("ENV") == "dev" {
		app.Use(logger.New(logger.Config{
			Format:     "${blue}${time} ${yellow}${status} - ${red}${latency} ${cyan}${method} ${path} ${green} ${ip} ${ua} ${reset}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   "Asia/Bangkok",
			Output:     os.Stdout,
		}))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("PFMS API " + os.Getenv("ENV") + " " + getVersionNumber())
	})

	port := viper.GetString("app.port")
	if os.Getenv("ENV") == "dev" || os.Getenv("ENV") == "prolocal" {
		port = viper.GetString("app.port_dev")
	}
	router.SetUpRouter(app)

	if err := app.Listen(port); err != nil {
		fmt.Println("error start server ->", err)
	}

}

func initConfig() {
	logs.Info("Init Config")
	switch os.Getenv("ENV") {
	case "":
		os.Setenv("ENV", "dev")
		viper.SetConfigName("config_dev")
	default:
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func getVersionNumber() string {
	version := "1.0.0"
	inFile, err := os.Open("./Makefile")
	if err != nil {
		logs.Error(err.Error() + `: ` + err.Error())
		return version
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		lineVersion := scanner.Text()
		if strings.TrimSpace(lineVersion) != "" {
			listFirstLine := strings.Split(lineVersion, " ")
			version = listFirstLine[len(listFirstLine)-1]
			break
		} else {
			break
		}
	}

	return version
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
