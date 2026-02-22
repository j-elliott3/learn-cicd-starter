package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header // Using the type from net/http
		want    string
		wantErr bool
	}{
		{
			name: "valid key",
			headers: http.Header{
				"Authorization": []string{"ApiKey key"},
			},
			want:    "key",
			wantErr: false,
		},
		{
			name:    "missing header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		{
			name: "empty value",
			headers: http.Header{
				"Authorization": []string{""},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "multiple headers - returns the first",
			headers: http.Header{
				"Authorization": []string{"ApiKey first-key", "ApiKey second-key"},
			},
			want:    "first-key",
			wantErr: false,
		},
	}

	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			got, err := GetAPIKey(tCase.headers)
			if (err != nil) != tCase.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tCase.wantErr)
				return
			}
			if got != tCase.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tCase.want)
			}
		})
	}
}