package wallet

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"time"
	"wallet/client/discount"
	"wallet/service/transaction"
	"wallet/storage/wallet"
)

const walletPrefix = "WALLET:%s"

type Service struct {
	wallet      wallet.Storage
	transaction *transaction.Service
	rdb         *redis.Client

	discount *discount.Client

	mu   sync.Mutex
	inTx bool
}

func New(
	wallet wallet.Storage,
	transaction *transaction.Service,
	discount *discount.Client,
	rdb *redis.Client,
) *Service {
	return &Service{
		wallet:      wallet,
		transaction: transaction,
		discount:    discount,
		rdb:         rdb,
	}
}

func (s *Service) WithTX(tx *sql.Tx) (*Service, error) {
	service := *s
	w, err := s.wallet.WithTX(tx)
	if err != nil {
		return nil, err
	}
	service.wallet = w
	service.inTx = true
	return &service, nil
}

func (s *Service) ToDBModel(w *DTO) *wallet.Wallet {
	return &wallet.Wallet{
		ID:       w.ID,
		MemberID: w.MemberID,
		Balance:  w.Balance,
	}

}

func (s *Service) FromDBModel(w *wallet.Wallet) *DTO {
	return &DTO{
		ID:        w.ID,
		MemberID:  w.MemberID,
		Balance:   w.Balance,
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
	}
}

func (s *Service) FromCreateRequest(r *CreateRequest) *wallet.Wallet {
	return &wallet.Wallet{
		WalletName: r.WalletName,
		MemberID:   r.MemberID,
		Balance:    r.Balance,
	}
}

func (s *Service) UpdateOrInsertInRedis(key string, g *DTO, exp time.Duration) error {
	err := s.rdb.Set(context.Background(), key, g, exp).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RetrieveFromRedis(key string) (*DTO, error) {
	w, err := s.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	walletRecord := &DTO{}
	err = json.Unmarshal([]byte(w), &walletRecord)
	if err != nil {
		return nil, err
	}
	return walletRecord, nil
}

func (s *Service) RemoveGiftFromRedis(w *DTO) {
	key := fmt.Sprintf(walletPrefix, strconv.FormatInt(w.MemberID, 10))
	s.RemoveWithKey(key)
}

func (s *Service) RemoveWithKey(k string) {
	s.rdb.Del(context.Background(), k)
}

func (w *DTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}
