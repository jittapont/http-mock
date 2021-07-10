package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var someError = errors.New("some error")

func Test_makeGetRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	type want struct {
		resp *http.Response
	}
	type args struct {
		client Client
		req    *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr error
	}{
		{
			name: "case 1",
			args: args{
				req: func() *http.Request {
					req, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
					if err != nil {
						t.Errorf("error in creating request: %s", err.Error())
					}
					return req
				}(),
				client: func() Client {
					client := NewMockClient(ctrl)
					req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
					resp := &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}
					client.EXPECT().Do(req).Times(1).Return(resp, nil)
					return client
				}(),
			},
			want: want{
				resp: func() *http.Response {
					return &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}
				}(),
			},
			wantErr: nil,
		},
		{
			name: "case 2",
			args: args{
				req: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "https://www.example.com", nil)
					return req
				}(),
				client: func() Client {
					client := NewMockClient(ctrl)
					req, _ := http.NewRequest(http.MethodGet, "https://www.example.com", nil)
					resp := &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}
					client.EXPECT().Do(req).Times(1).Return(resp, someError)
					return client
				}(),
			},
			want: want{
				resp: func() *http.Response {
					return &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}
				}(),
			},
			wantErr: someError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := makeGetRequest(tt.args.client, tt.args.req)
			assert.ErrorIs(t, err, tt.wantErr, "makeGetRequest erorr mismatch")
			assert.Equal(t, tt.want.resp, resp, "makeGetRequest response mismatch")
		})
	}
}
