package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/cucumber/godog"
	"github.com/jackc/pgx/v5"
)

type PgxStepHandler struct {
	conn *pgx.Conn
}

func NewPgxStepHandler(conn *pgx.Conn) *PgxStepHandler {
	return &PgxStepHandler{conn: conn}
}

func (h *PgxStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I see in "([^"]*)" table:$`, h.iSeeInTable)
	ctx.Step(`^I see record with ID "([^"]*)" in "([^"]*)" table:$`, h.iSeeWithIDInTable)
}

func (h *PgxStepHandler) iSeeInTable(tableName string, table *godog.Table) error {
	var columnNames []string

	for _, headerCell := range table.Rows[0].Cells {
		columnNames = append(columnNames, headerCell.Value)
	}

	for _, row := range table.Rows[1:] {
		var expectedValues []string

		for _, cell := range row.Cells {
			expectedValues = append(expectedValues, cell.Value)
		}

		rows, err := h.conn.Query(
			context.Background(),
			fmt.Sprintf("SELECT (%s) FROM %s", strings.Join(columnNames, ","), tableName),
		)
		if err != nil {
			return err
		}

		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}

			for i, val := range values {
				if val != expectedValues[i] {
					return fmt.Errorf("got: %s, want: %s", val, expectedValues[i])
				}
			}
		}
	}

	return nil
}

func (h *PgxStepHandler) iSeeWithIDInTable(id string, tableName string, table *godog.Table) error {
	var columnNames []string

	for _, headerCell := range table.Rows[0].Cells {
		columnNames = append(columnNames, headerCell.Value)
	}

	for _, row := range table.Rows[1:] {
		var expectedValues []string

		for _, cell := range row.Cells {
			expectedValues = append(expectedValues, cell.Value)
		}

		rows, err := h.conn.Query(
			context.Background(),
			fmt.Sprintf(`SELECT (%s) FROM %s where id=$1`, strings.Join(columnNames, ","), tableName),
			id,
		)
		if err != nil {
			return err
		}

		for rows.Next() {
			values, err := rows.Values()
			if err != nil {
				return err
			}
			stringSlice := make([]string, 0)

			for _, val := range values {
				switch v := val.(type) {
				case []interface{}:
					for _, i := range v {
						stringSlice = append(stringSlice, i.(string))
					}
				}
			}

			for i, val := range expectedValues {
				if val != stringSlice[i] {
					return fmt.Errorf("got: %s, want: %s", val, expectedValues[i])
				}
			}
		}
	}

	return nil
}
