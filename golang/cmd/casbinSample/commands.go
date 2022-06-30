package commands

import (
	"fmt"
	"github.com/apex/log"
	"github.com/casbin/casbin/v2"
	model2 "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type CommandsManager struct {
}

func NewManager() CommandsManager {
	return CommandsManager{}
}

func (m CommandsManager) SetupDB() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		m.setupDB()
		return nil
	}

	cmd := &cobra.Command{
		Use:  "setup-db",
		RunE: run,
	}
	return cmd
}

func (m CommandsManager) setupDB() {
	enforcer, _ := m.getEnforcer()

	userRoles := [][]string{
		{"alice", "admin", "domain1"},
		{"bob", "admin", "domain2"},
		{"ale", "admin", "domain1"},
	}
	_, _ = enforcer.AddNamedGroupingPolicies("g", userRoles)

	resourceRoles := [][]string{
		{"content", "root"},
		{"course", "content"},
		{"exam", "content"},
		{"course1", "course"},
		{"course2", "course"},
		{"exam1", "exam"},
		{"exam2", "exam"},
	}
	_, _ = enforcer.AddNamedGroupingPolicies("g2", resourceRoles)

	policies := [][]string{
		{"admin", "*", "data1", "read"},
		{"admin", "*", "data1", "write"},
		{"admin", "domain2", "data2", "read"},
		{"admin", "domain2", "data2", "write"},
		{"admin", "*", "content", "*"},
	}
	_, _ = enforcer.AddPolicies(policies)
	_ = enforcer.SavePolicy()
}

func (m CommandsManager) getEnforcer() (*casbin.SyncedEnforcer, error) {
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
		"m = r.sub.isAdmin == true || (g(r.sub.name, p.sub, r.dom) && g2(r.obj, p.obj) && keyMatch(r.dom, p.dom) && keyMatch(r.act, p.act))"
	model, err := model2.NewModelFromString(modelTest)
	if err != nil {
		log.WithError(err).Fatal("Fatal error")
	}

	a, err := m.getAdapter()
	if err != nil {
		log.WithError(err).Fatal("Fatal error")
	}

	enforcer, err := casbin.NewSyncedEnforcer(model, a)
	if err != nil {
		log.WithError(err).Fatal("Fatal error")
	}
	return enforcer, err
}

func (m CommandsManager) getAdapter() (persist.Adapter, error) {
	dbName := "jcasbin-sample"
	dbHost := envOrDefault("DB_HOST", "localhost")
	dbPort := envOrDefault("DB_PORT", "6543")
	dbUser := envOrDefault("DB_USER", "db-user")
	dbPassword := envOrDefault("DB_PASSWORD", "db-password")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword,
	)

	url := connectionString
	db, _ := gorm.Open(postgres.Open(url), &gorm.Config{})
	return gormadapter.NewAdapterByDB(db)
}

func envOrDefault(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		val = def
	}
	return val
}
