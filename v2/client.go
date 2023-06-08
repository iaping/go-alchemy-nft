package nft

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/valyala/fasthttp"
)

type Client struct {
	host, key string
	C         *fasthttp.Client
}

func New(h, k string, c *fasthttp.Client) *Client {
	if c == nil {
		c = DefaultFasthttpClient()
	}

	return &Client{
		host: h,
		key:  k,
		C:    c,
	}
}

func DefaultFasthttpClient() *fasthttp.Client {
	return &fasthttp.Client{
		Name:                "go-alchemy-nft",
		MaxConnsPerHost:     16,
		MaxIdleConnDuration: 20 * time.Second,
		ReadTimeout:         10 * time.Second,
		WriteTimeout:        10 * time.Second,
	}
}

func (c *Client) Do(ireq IRequest, iresp IResponse) error {
	url, body, err := c.format(ireq)
	if err != nil {
		return err
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	if ireq.IsPost() {
		req.Header.Set("content-type", "application/json")
		req.SetBody(body)
	}

	req.Header.Set("accept", "application/json")
	req.Header.SetMethod(ireq.GetMethod())
	req.SetRequestURI(url)

	if err := c.C.Do(req, resp); err != nil {
		return err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("code: %d, desc: %s", resp.StatusCode(), string(resp.Body()))
	}

	return json.Unmarshal(resp.Body(), &iresp)
}

func (c *Client) format(req IRequest) (string, []byte, error) {
	url := fmt.Sprintf("%s/nft/v2/%s/%s", c.host, c.key, req.GetPath())
	if req.IsPost() {
		body, err := json.Marshal(req.GetParam())
		if err != nil {
			return "", nil, err
		}
		return url, body, nil
	}

	v, err := query.Values(req.GetParam())
	if err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("%s?%s", url, v.Encode()), nil, nil
}

func (c *Client) NewGetNFTs() *GetNFTs {
	return NewGetNFTs(c)
}

func (c *Client) NewGetContractsForOwner() *GetContractsForOwner {
	return NewGetContractsForOwner(c)
}

func (c *Client) NewIsHolderOfCollection() *IsHolderOfCollection {
	return NewIsHolderOfCollection(c)
}

func (c *Client) NewGetOwnersForToken() *GetOwnersForToken {
	return NewGetOwnersForToken(c)
}

func (c *Client) NewGetOwnersForCollection() *GetOwnersForCollection {
	return NewGetOwnersForCollection(c)
}

func (c *Client) NewGetNFTMetadata() *GetNFTMetadata {
	return NewGetNFTMetadata(c)
}

func (c *Client) NewGetContractMetadata() *GetContractMetadata {
	return NewGetContractMetadata(c)
}

func (c *Client) NewInvalidateContract() *InvalidateContract {
	return NewInvalidateContract(c)
}

func (c *Client) NewReingestContract() *ReingestContract {
	return NewReingestContract(c)
}

func (c *Client) NewSearchContractMetadata() *SearchContractMetadata {
	return NewSearchContractMetadata(c)
}

func (c *Client) NewGetNFTsForCollection() *GetNFTsForCollection {
	return NewGetNFTsForCollection(c)
}

func (c *Client) NewGetNFTMetadataBatch() *GetNFTMetadataBatch {
	return NewGetNFTMetadataBatch(c)
}

func (c *Client) NewGetContractMetadataBatch() *GetContractMetadataBatch {
	return NewGetContractMetadataBatch(c)
}
