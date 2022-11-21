package middleware

import (
	"os"
	"testing"

	"github.com/hymnis/fikalistan/config"
	"github.com/hymnis/fikalistan/ent"
	"github.com/hymnis/fikalistan/pkg/services"
	"github.com/hymnis/fikalistan/pkg/tests"
)

var (
	c   *services.Container
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = services.NewContainer()

	// Create a user
	var err error
	if usr, err = tests.CreateUser(c.ORM); err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err = c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
