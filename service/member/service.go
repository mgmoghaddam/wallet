package member

import (
	"database/sql"
	"sync"
	"wallet/service/wallet"
	"wallet/storage/member"
)

type Service struct {
	member member.Storage
	wallet *wallet.Service

	mu   sync.Mutex
	inTx bool
}

func New(
	member member.Storage,
	wallet *wallet.Service,
) *Service {
	return &Service{
		member: member,
		wallet: wallet,
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
