package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	testDB *db
)

var migrationsDir = func() string {
	goPath := os.Getenv("GOPATH")
	return filepath.Join(goPath, fmt.Sprintf("src/github.com/stellar-map/stellar-map/data/migrations"))
}()

var dbURL = func() string {
	if env := os.Getenv("DB_URL_TEST"); env != "" {
		return env
	}
	return "postgres://postgres:@127.0.0.1:5432/stellarmap_test?sslmode=disable"
}()

var _ = BeforeSuite(func() {
	sqlDB, err := sql.Open("postgres", dbURL)
	Expect(err).NotTo(HaveOccurred())

	_, err = sqlDB.Exec("DROP SCHEMA public CASCADE")
	Expect(err).NotTo(HaveOccurred())
	_, err = sqlDB.Exec("CREATE SCHEMA public")
	Expect(err).NotTo(HaveOccurred())

	source := migrate.FileMigrationSource{Dir: migrationsDir}
	_, err = migrate.Exec(sqlDB, "postgres", source, migrate.Up)
	Expect(err).NotTo(HaveOccurred())

	repo, err := New(dbURL)
	Expect(err).NotTo(HaveOccurred())

	var ok bool
	testDB, ok = repo.(*db)
	Expect(ok).To(BeTrue())
})

var _ = AfterSuite(func() {
	err := testDB.Close()
	Expect(err).NotTo(HaveOccurred())
})

func TestDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Db Suite")
}
