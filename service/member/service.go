package member

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
	"wallet/internal/config"
	"wallet/service/wallet"
	"wallet/storage/member"
)

type Service struct {
	member member.Storage
	wallet *wallet.Service
	rdb    *redis.Client

	mu   sync.Mutex
	inTx bool
}

func New(
	member member.Storage,
	wallet *wallet.Service,
	rdb *redis.Client,
) *Service {
	return &Service{
		member: member,
		wallet: wallet,
		rdb:    rdb,
	}
}

func (s *Service) withTX(tx *sql.Tx) (*Service, error) {
	var err error
	service := *s
	service.member, err = s.member.WithTX(tx)
	if err != nil {
		return nil, err
	}
	service.wallet, err = s.wallet.WithTX(tx)
	if err != nil {
		return nil, err
	}
	service.inTx = true
	return &service, nil
}

func (s *Service) ToDBModel(u *DTO) *member.Member {
	return &member.Member{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}

}

func (s *Service) FromDBModel(u *member.Member) *DTO {
	return &DTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (s *Service) FromCreateRequest(r *CreateRequest) *member.Member {
	return &member.Member{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Phone:     r.Phone,
	}
}

func (s *Service) RetrieveFromRedis(key string) ([]*DTO, error) {
	ms, err := s.rdb.Get(context.Background(), config.RDBPrefix()+key).Result()
	if err != nil {
		return nil, err
	}
	var members []*DTO
	err = json.Unmarshal([]byte(ms), &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s *Service) UpdateOrInsertInRedis(key string, ms []*DTO, exp time.Duration) error {
	m, err := json.Marshal(ms)
	if err != nil {
		return err
	}
	err = s.rdb.Set(context.Background(), config.RDBPrefix()+key, m, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveWithKey(k string) {
	s.rdb.Del(context.Background(), config.RDBPrefix()+k)
}

func (ms *DTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(ms)
}
