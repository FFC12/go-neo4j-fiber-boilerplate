package database

import (
	"context"
	"fmt"
	"main/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	driver  neo4j.DriverWithContext
	initErr error
)

var Context *context.Context

func InitDB() context.Context {
	ctx := context.Background()
	Context = &ctx

	driver, initErr = neo4j.NewDriverWithContext(
		utils.DB_URI,
		neo4j.BasicAuth(utils.DB_USER, utils.DB_PASSWORD, ""))

	if initErr != nil {
		panic(initErr)
	}

	err := driver.VerifyConnectivity(ctx)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Neo4j Connection established")
	}

	return ctx
}

func ExecuteQuery(ctx context.Context, query string, params map[string]any) (*neo4j.EagerResult, error) {
	result, err := neo4j.ExecuteQuery(ctx, driver,
		query,
		params,
		neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase(utils.DB_NAME),
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func InitConstraints(ctx context.Context) {
	result, err := ExecuteQuery(ctx,
		`
		CREATE CONSTRAINT user_email IF NOT EXISTS
		FOR (user: User) REQUIRE user.email IS UNIQUE;
		CREATE CONSTRAINT user_username IF NOT EXISTS
		FOR (user: User) REQUIRE user.username IS UNIQUE;
	`, map[string]any{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("The query `%v` returned %v records in %+v.\n",
		result.Summary.Query().Text(), len(result.Records),
		result.Summary.ResultAvailableAfter())
}

func CreateSession(ctx context.Context, accessMode neo4j.AccessMode, dbName string) neo4j.SessionWithContext {
	return driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: dbName, AccessMode: accessMode})
}

func CloseSession(ctx context.Context, session neo4j.SessionWithContext) {
	session.Close(ctx)
}

func CloseDriver(ctx context.Context) {
	driver.Close(ctx)
}
