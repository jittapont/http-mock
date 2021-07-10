package roundtripper

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_makeGetRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		roundTrip RoundTripper
		req       *http.Request
	}
	type want struct {
		resp *http.Response
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
				roundTrip: func() RoundTripper {
					roundTrip := NewMockRoundTripper(ctrl)
					req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
					resp := &http.Response{
						Status:     "200 OK",
						StatusCode: 200,
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}
					roundTrip.EXPECT().RoundTrip(req).Return(resp, nil)
					return roundTrip
				}(),
				req: func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
					return req
				}(),
			},
			want: want{
				resp: &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader("")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{
				Transport: tt.args.roundTrip,
			}
			got, err := makeGetRequest(client, tt.args.req)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want.resp, got)
		})
	}
}
