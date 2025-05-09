package phoneconfirm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_digitCodeGenerator_Generate(t *testing.T) {
	validDigits := []byte("0123456789")
	type fields struct {
		cfg Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "no error 10 digit code",
			fields: fields{
				cfg: Config{
					Length: 10,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "no error 5 digit code",
			fields: fields{
				cfg: Config{
					Length: 5,
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &digitCodeGenerator{
				cfg: tt.fields.cfg,
			}
			got, err := g.Generate()

			tt.wantErr(t, err)
			for i := range got.value {
				assert.Contains(t, validDigits, got.value[i])
			}
			assert.Equal(t, tt.fields.cfg.ExpirationTime, got.ExpiresIn)
		})
	}
}
