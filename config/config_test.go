package config

import (
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestReadProperties(t *testing.T) {
	tests := []struct {
		name         string
		want         Config
		environments []string
		shouldPanic  bool
	}{
		{
			name: "ok",
			want: Config{
				App: app{Name: "transaction"},
				Database: database{
					Host: "localhost",
					Port: "5432",
					User: "local_development",
					Pwd:  "local_development",
					Base: "transaction",
				},
				ServerHTTP: serverHTTP{
					Port:         8080,
					WriteTimeout: time.Duration(15000000000),
					ReadTimeout:  time.Duration(15000000000),
				},
			},
			environments: []string{"../.env"},
			shouldPanic:  false,
		},
		{
			name:         "empty values",
			want:         Config{},
			environments: []string{},
			shouldPanic:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()

			if len(tt.environments) > 0 {
				err := godotenv.Overload(tt.environments...)
				assert.NoError(t, err)
			}

			if tt.shouldPanic {
				assert.Panics(t, func() {
					ReadProperties()
				})
				return
			}

			got := ReadProperties()
			assert.Equal(t, tt.want, got)
		})
	}
}
