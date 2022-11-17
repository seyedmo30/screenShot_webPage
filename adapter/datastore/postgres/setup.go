package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var err error
var conn *pgx.Conn

type Postgres struct {
}

func init() {
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:salam@localhost:5432/screenshot")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		panic("Failed to connect to database!")
	}

}
