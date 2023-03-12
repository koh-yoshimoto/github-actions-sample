package main

import (
	"context"
	"testing"
)

func Test_HandleRequest(t *testing.T) {
	tests := map[string]struct {
		in      Payload
		want    string
		hasErr  bool
		wantErr string
	}{
		"success": {
			in: Payload{
				Body: `{ "name": "World" }`,
			},
			want:    "Hello World from Lambda!",
			hasErr:  false,
			wantErr: "",
		},

		"error_bad_request": {
			in: Payload{
				Body: `{ "naem": "World" }`,
			},
			want:    "Hello World from Lambda!",
			hasErr:  true,
			wantErr: "bad request",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			res, err := HandleRequest(ctx, tt.in)
			t.Log(res)

			if err != nil && tt.hasErr {
				if err.Error() != tt.wantErr {
					t.Errorf("want %s, but got %s", tt.want, err.Error())
				}
				return
			}

			if err != nil {
				t.Error(err)
			}

			if res.Message != tt.want {
				t.Errorf("want %s, but got %s", tt.want, res.Message)
			}
		})
	}
}
