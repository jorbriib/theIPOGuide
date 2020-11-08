package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golossus/routing"
	"github.com/jorbriib/theIPOGuide/backend/src/application"
	"github.com/jorbriib/theIPOGuide/backend/src/infrastructure"
	"github.com/jorbriib/theIPOGuide/backend/src/ui/public/api"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := routing.NewRouter()

	db := getDB()

	corsOrigin := os.Getenv("CORS_ORIGIN")
	throttleLimit := os.Getenv("THROTTLE_LIMIT")
	throttleLimitFloat, err := strconv.ParseFloat(throttleLimit, 64)
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

	ipoRepository := infrastructure.NewMySQLIpoRepository(db)
	marketRepository := infrastructure.NewMySQLMarketRepository(db)
	companyRepository := infrastructure.NewMySQLCompanyRepository(db)
	countryRepository := infrastructure.NewMySQLCountryRepository(db)
	sectorRepository := infrastructure.NewMySQLSectorRepository(db)

	emailConfig := emailConfig()
	smtpEmailService := infrastructure.NewSmtpEmailService(emailConfig, smtp.SendMail)

	getIposService := application.NewGetIposService(ipoRepository, marketRepository, companyRepository, countryRepository, sectorRepository)
	getIposController := api.NewGetIposController(getIposService)
	_ = r.Get("/v1/ipos", getIposController.Run)

	searchByTextService := application.NewSearchByTextService(ipoRepository, marketRepository, companyRepository, countryRepository, sectorRepository)
	searchByTextController := api.NewSearchByTextController(searchByTextService)
	_ = r.Get("/v1/ipos/search", searchByTextController.Run)

	getRelatedIposService := application.NewGetRelatedIposService(marketRepository, countryRepository, sectorRepository)
	getRelatedIposController := api.NewGetRelatedIposController(getRelatedIposService)
	_ = r.Get("/v1/ipos/relations", getRelatedIposController.Run)

	getIpoService := application.NewGetIpoService(ipoRepository, marketRepository, companyRepository)
	getIpoController := api.NewGetIpoController(getIpoService)
	_ = r.Get("/v1/ipos/{alias}", getIpoController.Run)

	getSimilarIposService := application.NewGetSimilarIposService(ipoRepository, marketRepository, companyRepository)
	getSimilarIpoController := api.NewGetSimilarIposController(getSimilarIposService)
	_ = r.Get("/v1/ipos/{alias}/similar", getSimilarIpoController.Run)

	reportService := application.NewSendReportService(smtpEmailService)
	reportController := api.NewSendReportController(reportService)
	_ = r.Post(
		"/v1/report",
		api.VerifyRecaptcha(recaptchaSiteUrl, recaptchaSecret, reportController.Run),
	)

	contactFormService := application.NewSendContactFormService(smtpEmailService)
	contactController := api.NewSendContactFormController(contactFormService)
	_ = r.Post(
		"/v1/contact",
		api.VerifyRecaptcha(recaptchaSiteUrl, recaptchaSecret, contactController.Run),
	)

	_ = r.Get("/v1/{notFound}", notFound)
	_ = r.Get("/health-check", healthCheck)
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

func healthCheck(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
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

func emailConfig() infrastructure.EmailConfig {
	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailTo := os.Getenv("EMAIL_TO")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPort := os.Getenv("EMAIL_PORT")
	emailPortInt, err := strconv.Atoi(emailPort)
	if err != nil {
		log.Fatal(err)
	}

	return infrastructure.NewEmailConfig(emailHost, emailPortInt, emailFrom, emailPassword, emailFrom, emailTo)
}
