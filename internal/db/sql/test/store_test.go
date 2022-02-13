package db_test

import (
	"context"
	"github.com/stretchr/testify/require"
	db "highscores-core/internal/db/sql"
	"testing"
	"time"
)

func TestCreateCharacterTx(t *testing.T) {
	store := db.NewStore(testDB)

	characterParams := db.CreateCharacterParams{
		Name:      "me a pickle",
		LastCheck: time.Now(),
		CreatedAt: time.Now(),
	}
	character, err := store.CharacterCreationTx(context.Background(), characterParams)
	require.NoError(t, err)
	require.NotEmpty(t, character)
}

func TestUpdateCharacterTx(t *testing.T) {
	store := db.NewStore(testDB)
	now := time.Now()
	updateParams := db.UpdateCharacterParams{
		ID:        2,
		LastCheck: now,
	}

	err := store.CharacterUpdateTx(context.Background(), updateParams)
	require.NoError(t, err)
}
