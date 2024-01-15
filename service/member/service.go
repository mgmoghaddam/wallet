package member

import (
	"context"
	"database/sql"
	"encoding/json"
	"sync"
	"time"
	"wallet/db"
	"wallet/internal/config"
	"wallet/service/wallet"
	"wallet/storage/member"
)

const keyRdb = ":MEMBERS:"

type UseCase interface {
	Create(r *CreateRequest) (*DTO, error)
	GetById(id int64) (*DTO, error)
	Update(r *DTO) (*DTO, error)
	GetByPhone(phone string) (*DTO, error)
	GetMembersByGiftCode(gift string, limit, offset int) ([]*DTO, error)
	WithTX(tx *sql.Tx) (*Service, error)
}

type Service struct {
	member member.Repository
	wallet wallet.UseCase
	rdb    db.RedisClient

	mu   sync.Mutex
	inTx bool
}

func New(
	member member.Repository,
	wallet wallet.UseCase,
	rdb db.RedisClient,
) *Service {
	return &Service{
		member: member,
		wallet: wallet,
		rdb:    rdb,
	}
}

func (s *Service) WithTX(tx *sql.Tx) (*Service, error) {
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
	ms, err := s.rdb.Get(context.Background(), config.RDBPrefix()+key)
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
	err = s.rdb.Set(context.Background(), config.RDBPrefix()+key, m, exp)
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
