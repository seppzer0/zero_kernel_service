package lineage

import "context"

// Client -- a client structure
type Client struct {
	Device string
}

func (c *Client) GetRomList(ctx context.Context) {}
