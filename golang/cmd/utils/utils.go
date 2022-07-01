package utils

import (
	"fmt"
	"github.com/apex/log"
	"github.com/casbin/casbin/v2"
	model2 "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"os"
)

func GetEnforcer() (*casbin.CachedEnforcer, error) {
	modelTest := "[request_definition]\n" +
		"r = sub, dom, obj, act\n" +
		"[policy_definition]\n" +
		"p = sub, dom, obj, act\n" +
		"[role_definition]\n" +
		"g = _, _, _\n" +
		"g2 = _, _\n" +
		"[policy_effect]\n" +
		"e = some(where (p.eft == allow))\n" +
		"[matchers]\n" +
		"m = r.sub.IsAdmin == true || (g(r.sub.Name, p.sub, r.dom) && g2(r.obj, p.obj) && keyMatch(r.dom, p.dom) && keyMatch(r.act, p.act))"
	model, err := model2.NewModelFromString(modelTest)
	if err != nil {
		log.WithError(err).Fatal("Fatal error")
	}

	a, err := GetAdapter()
	if err != nil {
		log.WithError(err).Fatal("Fatal error")
	}

	enforcer, err := casbin.NewCachedEnforcer(model, a)
	if err != nil {
		log.WithError(err).Fatal("Fatal error")
	}
	return enforcer, err
}

func GetAdapter() (persist.Adapter, error) {
	dbName := "jcasbin-sample"
	dbHost := EnvOrDefault("DB_HOST", "localhost")
	dbPort := EnvOrDefault("DB_PORT", "6543")
	dbUser := EnvOrDefault("DB_USER", "db-user")
	dbPassword := EnvOrDefault("DB_PASSWORD", "db-password")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword,
	)

	url := connectionString
	db, _ := gorm.Open(postgres.Open(url), &gorm.Config{})
	return gormadapter.NewAdapterByDB(db)
}

func EnvOrDefault(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		val = def
	}
	return val
}

func RandomItem(items []string) string {
	return items[rand.Intn(len(items))]
}
