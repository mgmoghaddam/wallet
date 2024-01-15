package member

import (
	"github.com/rs/zerolog/log"
	"time"
)

// creates a new member.
func (s *Service) Create(r *CreateRequest) (*DTO, error) {
	memberRecord := s.FromCreateRequest(r)

	err := s.member.Create(memberRecord)
	if err != nil {
		return nil, err
	}

	return s.FromDBModel(memberRecord), nil
}

// gets a member by id.
func (s *Service) GetById(id int64) (*DTO, error) {
	member, err := s.member.GetById(id)
	if err != nil {
		return nil, err
	}
	return s.FromDBModel(member), nil
}

// updates a member by id.
func (s *Service) Update(r *DTO) (*DTO, error) {
	memberRecord := s.ToDBModel(r)

	err := s.member.Update(memberRecord)
	if err != nil {
		return nil, err
	}

	return s.FromDBModel(memberRecord), nil
}

// get by phone
func (s *Service) GetByPhone(phone string) (*DTO, error) {
	member, err := s.member.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	return s.FromDBModel(member), nil
}

func (s *Service) GetMembersByGiftCode(gift string, limit, offset int) ([]*DTO, error) {
	var members []*DTO
	members, err := s.RetrieveFromRedis(keyRdb + gift)
	if err != nil {
		wallets, err := s.wallet.GetByDiscountCodeWithPagination(gift, limit, offset)
		if err != nil {
			return nil, err
		}
		for _, w := range wallets {
			member, err := s.member.GetById(w.MemberID)
			if err != nil {
				return nil, err
			}
			members = append(members, s.FromDBModel(member))
		}
		err = s.UpdateOrInsertInRedis(keyRdb+gift, members, time.Minute*10)
		if err != nil {
			log.Error().Err(err).Msg("failed to update or insert in redis")
		}
	}
	return members, nil
}
