package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	customErr := Error("errorIs custom error")
	customWithTokenErr := Errorf("errorIs custom error %v", Token("token", "my custom token"))
	customWithTokensErr := Errorf(
		"errorIs custom error %v and %v",
		Token("token1", "my first custom token"),
		Token("token2", "my second custom token"),
	)

	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "check message",
			args: args{
				err: Error("my custom error"),
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something: my custom error")
			},
		},
		{
			name: "check unwrap 1",
			args: args{
				err: customErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something: errorIs custom error")
			},
		},
		{
			name: "check unwrap 2",
			args: args{
				err: customWithTokenErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something: errorIs custom error [token]")
			},
		},
		{
			name: "check unwrap 3",
			args: args{
				err: customWithTokensErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something: errorIs custom error [token1] and [token2]")
			},
		},
		{
			name: "several wrapped check unwrap",
			args: args{
				err: Wrap(customWithTokensErr, "do something 1"),
				msg: "do something 2",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something 2: do something 1: errorIs custom error [token1] and [token2]")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, Wrap(tt.args.err, tt.args.msg), fmt.Sprintf("Wrap(%v, %v)", tt.args.err, tt.args.msg))
		})
	}
}

func TestWrapf(t *testing.T) {
	customErr := Error("errorIs custom error")
	customWithTokenErr := Errorf("errorIs custom error %v", Token("token", "my custom token"))
	customWithTokensErr := Errorf(
		"errorIs custom error %v and %v",
		Token("token1", "my first custom token"),
		Token("token2", "my second custom token"),
	)

	type args struct {
		err    error
		msg    string
		tokens []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "check message",
			args: args{
				err: Error("my custom error"),
				msg: "do something %v",
				tokens: []any{
					Token("token", "my token"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something [token]: my custom error")
			},
		},
		{
			name: "check unwrap 1",
			args: args{
				err: customErr,
				msg: "do something %v and %v",
				tokens: []any{
					Token("token1", "my token 1"),
					Token("token2", "my token 2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something [token1] and [token2]: errorIs custom error")
			},
		},
		{
			name: "check unwrap 2",
			args: args{
				err: customWithTokenErr,
				msg: "do something %v",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something %v: errorIs custom error [token]")
			},
		},
		{
			name: "check unwrap 3",
			args: args{
				err: customWithTokensErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something: errorIs custom error [token1] and [token2]")
			},
		},
		{
			name: "several wrapped check unwrap",
			args: args{
				err: Wrapf(customWithTokensErr, "do something with %v", Token("tokenWrap1", "val1")),
				msg: "do something for %v",
				tokens: []any{
					Token("tokenWrap2", "val2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something for [tokenWrap2]: do something with [tokenWrap1]: errorIs custom error [token1] and [token2]")
			},
		},
		{
			name: "several wrapped same token check message",
			args: args{
				err: Wrapf(
					Errorf(
						"errorIs custom error %v and %v",
						Token("token1", "my first custom token"),
						Token("token2", "my second custom token"),
					),
					"do something with %v",
					Token("token1", "val1"),
				),
				msg: "do something for %v",
				tokens: []any{
					Token("token2", "val2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something for [token2#]: do something with [token1#]: errorIs custom error [token1] and [token2]")
			},
		},
		{
			name: "several wrapped same token check message 2",
			args: args{
				err: Wrapf(
					Wrapf(
						Errorf(
							"errorIs custom error %v and %v",
							Token("token1", "my first custom token"),
							Token("token2", "my second custom token"),
						),
						"internal %v",
						Token("token1", "val1"),
					),
					"do something with %v",
					Token("token1", "val1"),
				),
				msg: "do something for %v",
				tokens: []any{
					Token("token2", "val2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "do something for [token2#]: do something with [token1##]: internal [token1#]: errorIs custom error [token1] and [token2]")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, Wrapf(tt.args.err, tt.args.msg, tt.args.tokens...), fmt.Sprintf("Wrap(%v, %v)", tt.args.err, tt.args.msg))
		})
	}
}

func TestWrapFail(t *testing.T) {
	customErr := Error("errorIs custom error")
	customWithTokenErr := Errorf("errorIs custom error %v", Token("token", "my custom token"))
	customWithTokensErr := Errorf(
		"errorIs custom error %v and %v",
		Token("token1", "my first custom token"),
		Token("token2", "my second custom token"),
	)

	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "check message",
			args: args{
				err: Error("my custom error"),
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something: my custom error")
			},
		},
		{
			name: "check unwrap 1",
			args: args{
				err: customErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something: errorIs custom error")
			},
		},
		{
			name: "check unwrap 2",
			args: args{
				err: customWithTokenErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something: errorIs custom error [token]")
			},
		},
		{
			name: "check unwrap 3",
			args: args{
				err: customWithTokensErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something: errorIs custom error [token1] and [token2]")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, WrapFail(tt.args.err, tt.args.msg), fmt.Sprintf("Wrap(%v, %v)", tt.args.err, tt.args.msg))
		})
	}
}

func TestWrapFailf(t *testing.T) {
	customErr := Error("errorIs custom error")
	customWithTokenErr := Errorf("errorIs custom error %v", Token("token", "my custom token"))
	customWithTokensErr := Errorf(
		"errorIs custom error %v and %v",
		Token("token1", "my first custom token"),
		Token("token2", "my second custom token"),
	)

	type args struct {
		err    error
		msg    string
		tokens []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "check message",
			args: args{
				err: Error("my custom error"),
				msg: "do something %v",
				tokens: []any{
					Token("token", "my token"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something [token]: my custom error")
			},
		},
		{
			name: "check unwrap 1",
			args: args{
				err: customErr,
				msg: "do something %v and %v",
				tokens: []any{
					Token("token1", "my token 1"),
					Token("token2", "my token 2"),
				},
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something [token1] and [token2]: errorIs custom error")
			},
		},
		{
			name: "check unwrap 2",
			args: args{
				err: customWithTokenErr,
				msg: "do something %v",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something %v: errorIs custom error [token]")
			},
		},
		{
			name: "check unwrap 3",
			args: args{
				err: customWithTokensErr,
				msg: "do something",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "failed to do something: errorIs custom error [token1] and [token2]")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, WrapFailf(tt.args.err, tt.args.msg, tt.args.tokens...), fmt.Sprintf("Wrap(%v, %v)", tt.args.err, tt.args.msg))
		})
	}
}
