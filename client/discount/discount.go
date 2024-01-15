package discount

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wallet/internal/serr"
)

type Client interface {
	GetGiftByCode(code string) (*Gift, error)
	UseGift(code string) (*Gift, error)
}

type HTTPClient struct {
	address    string
	httpClient *http.Client
}

func NewHTTPClient(address string) *HTTPClient {
	c := &HTTPClient{address: address}
	c.httpClient = &http.Client{Timeout: time.Second * 10}
	return c
}

type Gift struct {
	Id             int    `json:"id"`
	Code           string `json:"code"`
	GiftAmount     int64  `json:"giftAmount"`
	UsageLimit     int64  `json:"usageLimit"`
	UsedCount      int64  `json:"usedCount"`
	ExpirationDate string `json:"expirationDate"`
	StartDateTime  string `json:"startDateTime"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

func (r *HTTPClient) GetGiftByCode(code string) (*Gift, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/gift/%s", r.address, code), nil)
	if err != nil {
		return nil, err
	}
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var result map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, serr.ValidationErr("getGift", result["message"].(string), serr.ErrDiscountClient)
	}
	var gift Gift
	err = json.NewDecoder(res.Body).Decode(&gift)
	if err != nil {
		return nil, err
	}
	return &gift, nil
}

func (r *HTTPClient) UseGift(code string) (*Gift, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/gift/use/%s", r.address, code), nil)
	if err != nil {
		return nil, err
	}
	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var result map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			return nil, err
		}
		return nil, serr.ValidationErr("useGift", result["message"].(string), serr.ErrDiscountClient)
	}
	var gift Gift
	err = json.NewDecoder(res.Body).Decode(&gift)
	if err != nil {
		return nil, err
	}
	return &gift, nil
}
