package cliq

import (
	"math/rand"
	"reflect"
	"testing"

	zoho "github.com/schmorrison/Zoho"
)

var (
	redirectURL = "https://example.org"
)

func TestAPI_PostMessageChannel(t *testing.T) {
	z := zoho.New()

	// ZohoCliq.Webhooks.CREATE,ZohoCliq.Channels.CREATE,ZohoCliq.Channels.READ,ZohoCliq.Channels.UPDATE,ZohoCliq.Channels.DELETE
	if err := z.GenerateTokenRequest("1000.HWGFZJVV8VLS6FZDQ4B9HWQMKLO4VY", "1e09dfa0978d4696ab0310c9a4a40dfff0d99a2733",
		"1000.fcca04caf6a62db4bd2307c00a856305.d5d28ead0c6314f669fcfdb3938da93d", redirectURL); err != nil {
		t.Fatal(err)
	}

	type fields struct {
		Zoho    *zoho.Zoho
		id      string
		channel string
	}
	type args struct {
		request PostMessageChannelData
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData PostMessageChannelResponse
		wantErr  bool
	}{
		{
			name: "test",
			fields: fields{
				Zoho: z,
				id: func() string {
					var id []byte
					keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
					for i := 0; i < 25; i++ {
						id = append(id, keyspace[rand.Intn(len(keyspace))])
					}
					return string(id)
				}(),
				channel: "zpeclouddevops",
			},
			args: args{PostMessageChannelData{
				Text: "testing golang integration",
			}},
			wantData: PostMessageChannelResponse{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &API{
				Zoho:    tt.fields.Zoho,
				id:      tt.fields.id,
				channel: tt.fields.channel,
			}
			gotData, err := c.PostMessageChannel(tt.args.request)
			t.Log(gotData.MessageId)
			if (err != nil) != tt.wantErr {
				t.Errorf("API.PostMessageChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("API.PostMessageChannel() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
