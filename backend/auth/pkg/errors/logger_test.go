package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_logger_errorArgs(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "2 errors",
			args: args{
				err: Wrapf(
					Errorf("main %v", Token("token1", "val1")),
					"first wrap %v",
					Token("token2", "val2"),
				),
			},
			want: []any{"token2", "val2", "token1", "val1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{}

			got := l.errorArgs(tt.args.err)

			assert.Equal(t, tt.want, got)
		})
	}
}
