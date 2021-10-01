package cliq

import (
	zoho "github.com/schmorrison/Zoho"
	"math/rand"
)

const ()

// API is used for interacting with the Zoho expense API
// the exposed methods are primarily access to expense modules which provide access to expense Methods
type API struct {
	*zoho.Zoho
	id      string
	channel string
}

// New returns a *cliq.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho, c string) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			id = append(id, keyspace[rand.Intn(len(keyspace))])
		}
		return string(id)
	}()
	API := &API{
		Zoho:    z,
		id:      id,
		channel: c,
	}
	return API
}
