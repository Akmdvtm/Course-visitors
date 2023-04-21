package user

import (
	"github.com/Akezhan1/lecvisitor/internal/util/test"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	// global test user storage
	// will be used in tests
	testStg *storage
)

// TestMain runs before all tests
// TestMain doing connection to database in docker postgres container and init global user storage
func TestMain(m *testing.M) {

	// get currently directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to resolve working directory: %v\n", err)
	}

	// create docker postgres container with sql migrations and test data,
	// and return connection to container
	testDB, shutDown := test.MustConnectToPG(&test.Config{
		Schemas:  test.GetDeploymentSchemas(),
		TestData: []string{filepath.Join(wd, "data_test.sql")},
	})

	// init global user storage for tests
	testStg = NewStorage(testDB)

	// run all tests
	exitCode := m.Run()

	// shutdown postgres docker container
	shutDown()

	// exit from TestMain func
	os.Exit(exitCode)
}
