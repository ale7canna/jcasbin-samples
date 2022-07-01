package utils

import (
	"bufio"
	"fmt"
	"github.com/apex/log"
	"github.com/casbin/casbin/v2"
	"github.com/spf13/cobra"
	"golang/cmd/utils"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type CommandsManager struct {
	Enforcer *casbin.CachedEnforcer
}

func NewManager() CommandsManager {
	start := time.Now().UnixMilli()
	enforcer, err := utils.GetEnforcer()
	log.WithField("timeSpent", time.Now().UnixMilli()-start).Info("enforcer init took {timeSpent}")

	if err != nil {
		panic(err)
	}
	return CommandsManager{
		Enforcer: enforcer,
	}
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

func (m CommandsManager) CheckPolicy() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		m.checkPolicy(args)
		return nil
	}

	cmd := &cobra.Command{
		Use:  "check-policy",
		RunE: run,
	}
	return cmd
}

func (m CommandsManager) Benchmark() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		m.benchmark(args)
		return nil
	}

	cmd := &cobra.Command{
		Use:  "benchmark",
		RunE: run,
	}
	return cmd
}

func (m CommandsManager) Interactive() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		m.interactive(args)
		return nil
	}

	cmd := &cobra.Command{
		Use:  "interactive",
		RunE: run,
	}
	return cmd
}

func (m CommandsManager) setupDB() {
	enforcer := m.Enforcer

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

func (m CommandsManager) checkPolicy(policy []string) {
	enforcer := m.Enforcer

	subject := policy[0]
	domain := policy[1]
	obj := policy[2]
	action := policy[3]
	sub := CustomSubject{Name: subject, IsAdmin: false} // the user that wants to access a resource.
	result, err := enforcer.Enforce(sub, domain, obj, action)
	if err != nil {
		log.WithError(err).Fatal("Error")
	}
	log.WithField("policy", policy).WithField("result", result).
		Info("Checking policy {policy}. Result: {result}")
}

func (m CommandsManager) benchmark(args []string) {
	enforcer := m.Enforcer

	names := utils.GetNames()
	actions := []string{"read", "edit", "consume", "share"}
	objects := utils.GetResources()
	domains := utils.GetDomains()
	nPolicies := 100
	if len(args) > 0 {
		nPolicies, _ = strconv.Atoi(args[0])
	}

	start := time.Now().UnixMilli()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= nPolicies; i++ {
		name := utils.RandomItem(names)
		domain := utils.RandomItem(domains)
		obj := utils.RandomItem(objects)
		act := utils.RandomItem(actions)
		isAdmin := false
		sub := CustomSubject{Name: name, IsAdmin: isAdmin}
		_, err := enforcer.Enforce(sub, domain, obj, act)
		if err != nil {
			log.WithError(err).Fatal("Error")
		}
	}
	log.WithField("nPolicies", nPolicies).
		WithField("timeSpent", time.Now().UnixMilli()-start).
		Info("Computing {nPolicies} policies took {timeSpent} ms")
}

func (m CommandsManager) interactive(args []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, _ := reader.ReadString('\n')
	for input != "exit" {
		i := strings.Split(strings.TrimSuffix(input, "\n"), " ")
		command := i[0]
		args = i[1:]

		switch command {
		case "check":
			m.checkPolicy(args)
		case "setup-db":
			m.setupDB()
		case "benchmark":
			m.benchmark(args)
		default:
			break
		}
		fmt.Print("Enter command: ")
		input, _ = reader.ReadString('\n')
	}
}
