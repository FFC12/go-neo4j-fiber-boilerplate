package database

import (
	"context"
	"errors"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SessionMode int

const (
	SessionWrite = iota
	SessionRead
)

func SessionExecuteMany(
	ctx context.Context,
	session neo4j.SessionWithContext,
	mode SessionMode,
	f func(neo4j.ManagedTransaction) (any, error)) ([]*neo4j.Record, error) {

	if mode == SessionWrite {
		records, err := session.ExecuteWrite(ctx, f)

		if err != nil {
			return nil, err
		}

		return records.([]*neo4j.Record), err
	}

	if mode == SessionRead {
		records, err := session.ExecuteRead(ctx, f)

		if err != nil {
			return nil, err
		}

		return records.([]*neo4j.Record), err
	}

	return nil, errors.New("While executing a session an error occurred!")
}

func SessionExecuteOne(
	ctx context.Context,
	session neo4j.SessionWithContext,
	mode SessionMode,
	f func(neo4j.ManagedTransaction) (any, error)) (any, error) {

	if mode == SessionWrite {
		records, err := session.ExecuteWrite(ctx, f)

		if err != nil {
			return nil, err
		}

		return records, err
	}

	if mode == SessionRead {
		records, err := session.ExecuteRead(ctx, f)

		if err != nil {
			return nil, err
		}

		return records, err
	}

	return nil, errors.New("While executing a session an error occurred!")
}

