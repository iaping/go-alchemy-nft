package nft

type Api struct {
	c *Client
}

func (a *Api) request(path, method string, p interface{}) IRequest {
	return &Request{
		Path:   path,
		Method: method,
		Param:  p,
	}
}
