package client

import (
	"reflect"
	"testing"
)

func TestArgs(t *testing.T) {
	cmd := "foo"

	t.Run("should contain a Command and Params", func(t *testing.T) {
		params := []string{"bar", "baz"}
		args := &Args{Command: cmd, Params: params}

		if got, want := args.Command, cmd; got != want {
			t.Fatalf("args.Command = %q, want %q", got, want)
		}

		if !reflect.DeepEqual(args.Params, params) {
			t.Fatalf("args.Params != params")
		}
	})

	t.Run("newArgs", func(t *testing.T) {
		t.Run("with no arguments", func(t *testing.T) {
			args := newArgs([]string{})

			if got, want := args.Command, ""; got != want {
				t.Fatalf("args.Command = %q, want %q", got, want)
			}
		})

		t.Run("with command", func(t *testing.T) {
			args := newArgs([]string{cmd})

			if got, want := args.Command, cmd; got != want {
				t.Fatalf("args.Command = %q, want %q", got, want)
			}
		})

		t.Run("with command and params", func(t *testing.T) {
			args := newArgs([]string{cmd, "param1", "param2"})

			if got, want := args.Command, cmd; got != want {
				t.Fatalf("args.Command = %q, want %q", got, want)
			}
		})
	})

	t.Run("IsParamsEmpty", func(t *testing.T) {
		t.Run("with no params", func(t *testing.T) {
			args := newArgs([]string{cmd})

			if !args.IsParamsEmpty() {
				t.Fatalf("args.IsParamsEmpty() is false")
			}
		})

		t.Run("with params", func(t *testing.T) {
			args := newArgs([]string{cmd, "bar", "baz"})

			if args.IsParamsEmpty() {
				t.Fatalf("args.IsParamsEmpty() is true")
			}
		})
	})
}
