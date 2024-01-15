package transaction

func (s *Service) Create(r *CreateRequest) (*DTO, error) {
	t := s.FromCreateRequest(r)
	err := s.transaction.Insert(t)
	if err != nil {
		return nil, err
	}
	return s.FromDBModel(t), nil
}

func (s *Service) GetByID(id int64) (*DTO, error) {
	t, err := s.transaction.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.FromDBModel(t), nil
}

func (s *Service) GetByWalletID(walletID int64) ([]*DTO, error) {
	ts, err := s.transaction.GetByWalletID(walletID)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		result = append(result, s.FromDBModel(t))
	}
	return result, nil
}

func (s *Service) GetByWalletIDWithPagination(walletID int64, limit, offset int) ([]*DTO, error) {
	ts, err := s.transaction.GetByWalletIDWithPagination(walletID, limit, offset)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		result = append(result, s.FromDBModel(t))
	}
	return result, nil
}

func (s *Service) GetByWalletIDAndType(walletID int64, transactionType Type) ([]*DTO, error) {
	ts, err := s.transaction.GetByWalletIDAndType(walletID, TypeToDBType(transactionType))
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		result = append(result, s.FromDBModel(t))
	}
	return result, nil
}

func (s *Service) GetByWalletIDAndDiscountCode(walletID int64, discountCode string) ([]*DTO, error) {
	ts, err := s.transaction.GetByWalletIDAndDiscountCode(walletID, discountCode)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		result = append(result, s.FromDBModel(t))
	}
	return result, nil
}

func (s *Service) GetByWalletIDAndTypeAndDiscountCode(walletID int64, transactionType Type, discountCode string) ([]*DTO, error) {
	ts, err := s.transaction.GetByWalletIDAndTypeAndDiscountCode(walletID, TypeToDBType(transactionType), discountCode)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		result = append(result, s.FromDBModel(t))
	}
	return result, nil
}

func (s *Service) GetByDiscountCodeWithPagination(discountCode string, limit, offset int) ([]*DTO, error) {
	ts, err := s.transaction.GetByDiscountCodeWithPagination(discountCode, limit, offset)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		result = append(result, s.FromDBModel(t))
	}
	return result, nil
}

func (s *Service) DeleteByWalletID(walletID int64) error {
	return s.transaction.DeleteByWalletID(walletID)
}

func (s *Service) Delete(id int64) error {
	return s.transaction.DeleteByID(id)
}

func (s *Service) GetBalance(walletID int64) (int64, error) {
	return s.transaction.GetBalance(walletID)
}
