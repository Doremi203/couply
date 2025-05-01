package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_messageString(t *testing.T) {
	type args struct {
		format string
		tokens []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no tokens",
			args: args{
				format: "my error message",
			},
			want: "my error message",
		},
		{
			name: "with tokens 1",
			args: args{
				format: "my error message with %v",
				tokens: []any{
					Token("token", "value"),
				},
			},
			want: "my error message with [token]",
		},
		{
			name: "with tokens 2",
			args: args{
				format: "my error message with %v and %v",
				tokens: []any{
					Token("token1", "value1"),
					Token("token2", 2),
				},
			},
			want: "my error message with [token1] and [token2]",
		},
		{
			name: "with tokens 3",
			args: args{
				format: "my error message with %s and %s",
				tokens: []any{
					Token("token1", "value1"),
					Token("token2", 2),
				},
			},
			want: "my error message with [token1] and [token2]",
		},
		{
			name: "with default args",
			args: args{
				format: "my error message with %s and %s and %d",
				tokens: []any{
					"value1", "value2", 2,
				},
			},
			want: "my error message with value1 and value2 and 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := message{
				format: tt.args.format,
				tokens: tt.args.tokens,
			}
			got := msg.String()

			assert.Equal(t, tt.want, got)
		})
	}
}
