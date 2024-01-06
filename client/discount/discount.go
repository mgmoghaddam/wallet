package discount

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wallet/internal/serr"
)

type Client struct {
	net *http.Client

	address string
}

func New(address string) *Client {
	c := &Client{address: address}
	c.net = &http.Client{Timeout: time.Second * 10}
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

func (c *Client) GetGiftByCode(code string) (*Gift, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/gift/%s", c.address, code), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.net.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", res.StatusCode)
	}
	var g Gift
	err = json.NewDecoder(res.Body).Decode(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (c *Client) UseGift(code string) (*Gift, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/gift/use/%s", c.address, code), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.net.Do(req)
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
	var g Gift
	err = json.NewDecoder(res.Body).Decode(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil

}
