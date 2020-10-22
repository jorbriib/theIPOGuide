package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/src/ipo/application"
	"github.com/jorbriib/theIPOGuide/src/ipo/infrastructure"
	"github.com/jorbriib/theIPOGuide/src/ipo/ui/public/api"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := routing.NewRouter()

	db := getDB()

	corsOrigin := os.Getenv("CORS_ORIGIN")
	throttleLimit := os.Getenv("THROTTLE_LIMIT")
	throttleLimitFloat, err := strconv.ParseFloat(throttleLimit,64)
	if err != nil {
		log.Fatal(err)
	}
	throttleBucket := os.Getenv("THROTTLE_BUCKET")
	throttleBucketInt, err := strconv.Atoi(throttleBucket)
	if err != nil {
		log.Fatal(err)
	}

	recaptchaSiteUrl := os.Getenv("RECAPTCHA_SITE_URL")
	recaptchaSecret := os.Getenv("RECAPTCHA_SECRET")

	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailTo := os.Getenv("EMAIL_TO")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPort := os.Getenv("EMAIL_PORT")
	emailPortInt, err := strconv.Atoi(emailPort)
	if err != nil {
		log.Fatal(err)
	}

	ipoRepository := infrastructure.NewMySQLIpoRepository(db)
	marketRepository := infrastructure.NewMySQLMarketRepository(db)
	companyRepository := infrastructure.NewMySQLCompanyRepository(db)

	service := application.NewService(ipoRepository, marketRepository, companyRepository)
	controller := api.NewController(service)

	_ = r.Get("/v1/ipos", controller.GetIpos)
	_ = r.Get("/v1/ipos/{alias}", controller.GetIpo)

	smtpEmailService :=  infrastructure.NewSmtpEmailService(emailHost, emailFrom, emailTo, emailPassword, emailPortInt)
	reportService := application.NewReportService(smtpEmailService)
	reportController := api.NewReportController(reportService)
	_ = r.Post(
		"/v1/report",
		api.VerifyRecaptcha(recaptchaSiteUrl, recaptchaSecret, reportController.SendReport),
	)

	_ = r.Get("/v1/{notFound}", notFound)
	_ = r.Get("/{notFound}", notFound)

	log.Println("Server listening")
	errServer := http.ListenAndServe(":80",
		api.ContentTypeMiddleware(
			"application/json; charset=UTF-8",
			api.EnableCorsMiddleware(
				corsOrigin,
				api.ThrottleMiddleware(
					throttleLimitFloat,
					throttleBucketInt,
					&r,
				),
			),
		),
	)
	if errServer != nil {
		log.Fatal(errServer)
	}
}

func notFound(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
}


func getDB() *sql.DB {
	mysqlAddr := os.Getenv("MYSQL_ADDR")
	mysqlDBName := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlAddr, mysqlDBName)

	db, errorDB := sql.Open("mysql", conn)
	if errorDB != nil {
		log.Fatal(errorDB)
	}
	return db
}
