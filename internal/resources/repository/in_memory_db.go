package repository

import (
	"sync"

	"yalo-api/internal/resources/models"
)

type UserInteractionDB struct {
	store map[string]models.UserInteraction
	mu    sync.RWMutex
}

func NewUserInteractionDB() *UserInteractionDB {
	return &UserInteractionDB{
		store: make(map[string]models.UserInteraction),
	}
}

func (ts *UserInteractionDB) Set(key string, userInteraction models.UserInteraction) bool {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	ts.store[key] = userInteraction

	return true
}

func (ts *UserInteractionDB) Get(key string) (models.UserInteraction, bool) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	value, exists := ts.store[key]

	return value, exists
}

func (ts *UserInteractionDB) Delete(key string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.store, key)
}

func (ts *UserInteractionDB) GetAll() map[string]models.UserInteraction {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	copyMap := make(map[string]models.UserInteraction)

	for key, value := range ts.store {
		copyMap[key] = value
	}

	return copyMap
}

func (ts *UserInteractionDB) GetAllByUserID(userID string) map[string]models.UserInteraction {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	copyMap := make(map[string]models.UserInteraction)

	for key, value := range ts.store {
		if value.UserID == userID {
			copyMap[key] = value
		}
	}

	return copyMap
}
