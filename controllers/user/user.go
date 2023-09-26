package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"main/database"
	"main/models"
	"main/utils"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func UserSignUp(ctx context.Context, user models.UserSignUp) bool {
	session := database.CreateSession(ctx, neo4j.AccessModeWrite, utils.DB_NAME)

	defer database.CloseSession(ctx, session)

	exist, err := database.SessionExecuteOne(ctx, session, database.SessionWrite,
		func(tx neo4j.ManagedTransaction) (any, error) {
			if ManagedCheckIfUserExist(ctx, tx, user.Username, user.Email) {
				return false, nil
			}

			createUserQuery := `
				CREATE 
				(:User {name: $name, lastName: $lastName, password: $password, email: $email, username: $username})
			`

			hashed, err := utils.HashPassword(user.Password)

			if err != nil {
				return nil, err
			}

			createUserParams := map[string]any{
				"name":     user.Name,
				"lastName": user.LastName,
				"password": hashed,
				"email":    user.Email,
				"username": user.Username,
			}

			result, err := tx.Run(ctx, createUserQuery, createUserParams)

			if err != nil {
				return nil, err
			}

			summary, err := result.Consume(ctx)

			if err != nil {
				return nil, err
			}

			utils.InfoLogger.Println(summary.Counters().NodesCreated())

			return true, nil

		})

	if err != nil {
		utils.ErrorLogger.Panic(err)
	}

	return exist.(bool)
}

func UserLogin(ctx context.Context, user models.UserLogin) (map[string]any, error) {
	session := database.CreateSession(ctx, neo4j.AccessModeWrite, utils.DB_NAME)

	defer database.CloseSession(ctx, session)

	data, err := database.SessionExecuteOne(ctx, session, database.SessionWrite,
		func(tx neo4j.ManagedTransaction) (any, error) {
			matchUserQuery := `
				MATCH (user:User)
				WHERE user.username = $username 
				RETURN ID(user) AS nodeID, user
			`

			matchUserParams := map[string]any{
				"username": user.Username,
			}

			result, err := tx.Run(ctx, matchUserQuery, matchUserParams)
			if err != nil {
				return nil, err
			}

			record, err := result.Single(ctx)

			if err != nil {
				return nil, err
			}

			nodeID, exists := record.Get("nodeID")
			if !exists {
				return nil, errors.New("ID does not exist at node?")
			}

			jsonBytes, err := json.Marshal(record.AsMap())
			if err != nil {
				return nil, err
			}

			var userResponse database.DBUserResponse

			if err := json.Unmarshal(jsonBytes, &userResponse); err != nil {
				return nil, err
			}

			props := userResponse.User.Props

			// Check if password correct or not
			if !utils.CheckPasswordHash(user.Password, props["password"].(string)) {
				utils.InfoLogger.Println("Password mismatched for: ", props["username"].(string))

				return nil, errors.New("invalid password")
			}

			return map[string]any{
				"id":       nodeID.(int64),
				"username": props["username"],
				"email":    props["email"],
			}, nil

		})

	if err != nil {
		return nil, err
	} else {
		return data.(map[string]any), nil
	}
}

// These properties shows also what properties has constraints...
// They have to be unique and DB cannot have more than one.
// That's why we're checking that they exist or not on DB
// *Managed* implies, this is a transaction which is managed by us
// If you don't know what's going on here, see documentation:
// https://neo4j.com/docs/go-manual/current/transactions/
func ManagedCheckIfUserExist(ctx context.Context, tx neo4j.ManagedTransaction, username string, email string) bool {
	result, err := tx.Run(ctx,
		`
		MATCH (user:User)
		WHERE user.username = $username OR user.email = $email
		RETURN user
	`,
		map[string]any{
			"username": username,
			"email":    email,
		})

	if err != nil {
		utils.ErrorLogger.Panic(err)
	}

	return result.Next(ctx)
}
