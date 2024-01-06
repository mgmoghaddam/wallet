package member

import "time"

type DTO struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

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
	wallets, err := s.wallet.GetByDiscountCodeWithPagination(gift, limit, offset)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, w := range wallets {
		member, err := s.member.GetById(w.MemberID)
		if err != nil {
			return nil, err
		}
		result = append(result, s.FromDBModel(member))
	}
	return result, nil
}
