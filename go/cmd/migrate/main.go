package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"

	"github.com/stellar-map/stellar-map/go/pkg/db"
)

const migrationTemplate = `-- +migrate Up


-- +migrate Down

`

const (
	defaultSourceDir = "data/migrations"
)

func init() {
	migrate.SetTable(db.MigrationsTable)
}

func main() {
	var (
		dbURL     string
		sourceDir string
		limit     int
	)

	rootCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate is a tool for applying database migrations",
		Long:  "migrate is a tool for applying database migrations",
	}
	rootCmd.PersistentFlags().StringVar(&dbURL, "db", "", "postgres database URL")
	rootCmd.PersistentFlags().StringVar(&sourceDir, "source", defaultSourceDir, "directory of migration scripts")

	createCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Generate a new migration SQL script",
		Long:  "Generate a new migration SQL script",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("missing migration name")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			source := migrate.FileMigrationSource{Dir: sourceDir}
			showErr(runCreate(args[0], source))
		},
	}

	upCmd := &cobra.Command{
		Use:   "up [N]",
		Short: "Apply forward migrations",
		Long:  "Apply forward migrations",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}
			n, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return errors.New("error parsing limit N")
			}
			limit = int(n)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			db := connectDb(dbURL)
			source := migrate.FileMigrationSource{Dir: sourceDir}
			showErr(runUp(db, source, limit))
		},
	}

	downCmd := &cobra.Command{
		Use:   "down N",
		Short: "Apply backward migrations",
		Long:  "Apply backward migrations",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("missing limit N")
			}
			n, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return errors.New("error parsing limit N")
			}
			limit = int(n)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			db := connectDb(dbURL)
			source := migrate.FileMigrationSource{Dir: sourceDir}
			showErr(runDown(db, source, limit))
		},
	}

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Show migration status",
		Long:  "Show migration status",
		Run: func(cmd *cobra.Command, args []string) {
			db := connectDb(dbURL)
			source := migrate.FileMigrationSource{Dir: sourceDir}
			showErr(runStatus(db, source))
		},
	}

	rootCmd.AddCommand(
		createCmd,
		upCmd,
		downCmd,
		statusCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func connectDb(url string) *sql.DB {
	if url == "" {
		url = os.Getenv("DB_URL")
	}
	db, err := sql.Open("postgres", url)
	if err != nil {
		showErr(errors.Wrap(err, "error connecting to db"))
	}
	if err := db.Ping(); err != nil {
		showErr(errors.Wrap(err, "error pinging db"))
	}
	return db
}

func runCreate(name string, source migrate.FileMigrationSource) error {
	if name == "" {
		return errors.New("please specify migration name")
	}

	fileName := fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102150405"), strings.TrimSpace(name))
	filePath := path.Join(source.Dir, fileName)

	err := os.MkdirAll(source.Dir, os.ModePerm)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePath, []byte(migrationTemplate), 0644); err != nil {
		return err
	}

	_, _ = fmt.Printf("created migration %s\n", filePath)
	return nil
}

func runUp(db *sql.DB, source migrate.FileMigrationSource, limit int) error {
	_, err := migrate.ExecMax(db, "postgres", source, migrate.Up, limit)
	return err
}

func runDown(db *sql.DB, source migrate.FileMigrationSource, limit int) error {
	_, err := migrate.ExecMax(db, "postgres", source, migrate.Down, limit)
	return err
}

func runStatus(db *sql.DB, source migrate.FileMigrationSource) error {
	migrations, err := source.FindMigrations()
	if err != nil {
		return err
	}

	records, err := migrate.GetMigrationRecords(db, "postgres")
	if err != nil {
		return err
	}

	rows := make(map[string]*statusRow)
	for _, m := range migrations {
		rows[m.Id] = &statusRow{
			ID:       m.Id,
			Migrated: false,
		}
	}

	for _, r := range records {
		if rows[r.Id] == nil {
			return errors.Errorf("could not find migration file: %v", r.Id)
		}
		rows[r.Id].Migrated = true
		rows[r.Id].AppliedAt = r.AppliedAt
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Migration", "Applied"})
	table.SetColWidth(60)

	for _, m := range migrations {
		row := rows[m.Id]
		if row.Migrated {
			table.Append([]string{
				m.Id,
				row.AppliedAt.String(),
			})
		} else {
			table.Append([]string{
				m.Id,
				"",
			})
		}
	}

	table.Render()
	return nil
}

func showErr(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

type statusRow struct {
	ID        string
	Migrated  bool
	AppliedAt time.Time
}
