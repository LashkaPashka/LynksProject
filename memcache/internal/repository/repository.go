package repository

import (
	"memCache/internal/db"
	"memCache/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type MemcacheRepository struct {
	rDb *db.Db
}

func New() (*MemcacheRepository, error){
	repo := &MemcacheRepository{}

	db, err := db.New()
	if err != nil{
		return nil, err
	}

	repo.rDb = db
	return repo, nil
}

func (m *MemcacheRepository) GetInMemory(hash string) (*model.LinkInMemory, error) {
	var data *model.LinkInMemory
	cmd := m.rDb.RedisDb.Get(context.Background(), "links:" + hash)

	if err := json.Unmarshal([]byte(cmd.Val()), &data); err != nil {
		return nil, err
	}
	
	fmt.Println(data)
	return data, nil
}


func (m *MemcacheRepository) SaveInMemory(data model.LinkInMemory) error {
	hash := data.Short_url[len(data.Short_url)-10:]
	
	key := "links:" + hash

	val, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err = m.rDb.RedisDb.Set(context.Background(), key, string(val), time.Minute*10).Err(); err != nil {
		return err
	}

	return nil
}