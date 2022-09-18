package db

import (
	"github.com/timshannon/bolthold"
	"os"
	"regexp"
)

var store *bolthold.Store

func Store() *bolthold.Store {
	return store
}

func Open(cfg Options) (err error) {
	store, err = bolthold.Open(cfg.Path, os.ModePerm, nil)
	return
}

func Close() error {
	return store.Close()
}

func Search[T any](fields []string, keyword string, skip, limit int) ([]T, error) {
	query := &bolthold.Query{}
	if keyword != "" {
		re := regexp.MustCompile(keyword)
		for _, f := range fields {
			query.And(f).RegExp(re)
		}
	}
	var result []T
	err := store.Find(&result, query.Skip(skip).Limit(limit))
	if err == bolthold.ErrNotFound {
		result = make([]T, 0)
		err = nil
	}
	return result, err
}

func List[T any](skip, limit int) ([]T, error) {
	query := &bolthold.Query{}
	var result []T
	err := store.Find(&result, query.Skip(skip).Limit(limit))
	return result, err
}
