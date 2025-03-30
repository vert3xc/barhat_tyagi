package utils

import (
	"database/sql"
	"fmt"
	"os"
        "log"
        "reflect"
        "strings"

	_ "github.com/lib/pq"
)

func ConnectToDb() (*sql.DB, error) {
	pwd := os.Getenv("POSTGRES_PASSWORD")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	var connStr string
	connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return db, nil
}

func getDataAsync[T any](db *sql.DB, query string, args ...any) (<-chan T, error) {
    chanResult := make(chan T, 1)
    go func() {
        defer close(chanResult)
        rows, err := db.Query(query, args...)
        if err != nil{
            log.Printf("Error with database %s", err.Error())
            return chanResult, err
        }
        columns, _ := rows.Columns()
        columnTypes, _ := rows.ColumnTypes()
        for rows.Next() {
            var result T
            resultVal := reflect.ValueOf(&result).Elem()
            pointers := preparePointers(resultVal, columns, columnTypes)
            if err := rows.Scan(pointers...); err != nil{
                return chanResult, err
            }
            chanResult <- result
        }
    }()
    return chanResult, nil
}

func preparePointers(
    v reflect.Value,
    columns []string,
    columnTypes []*sql.ColumnType,
) []any {
    pointers := make([]interface{}, len(columns))
    typ := v.Type()

    for i, colName := range columns {
        colName = strings.ToLower(strings.TrimSpace(colName))

        for j := 0; j < typ.NumField(); j++ {
            field := typ.Field(j)
            fieldName := strings.ToLower(field.Name)

            tag := field.Tag.Get("sql")
            if tag == "" {
                tag = field.Tag.Get("db")
            }
            if tag != "" {
                tag = strings.Split(tag, ",")[0]
                tag = strings.ToLower(tag)
            }

            if tag == colName || (tag == "" && fieldName == colName) {
                pointers[i] = v.Field(j).Addr().Interface()
                break
            }
        }
    }

    return pointers
}
