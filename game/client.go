package game

import (
	"github.com/z46-dev/arctic"
	"github.com/z46-dev/genn/shared"
)

func NewClient(g *Game, c *arctic.ServerClient) (client *Client) {
	return
}

func (c *Client) OnMessage(r *shared.Reader) {}

func (c *Client) Remove() {}