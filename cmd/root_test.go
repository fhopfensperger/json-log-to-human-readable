package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_noArgs(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "to many args",
			args: args{
				cmd:  rootCmd,
				args: []string{"arg1"},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				cmd:  rootCmd,
				args: []string{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := noArgs(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("noArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// dummy test
func Test_isInputFromPipe(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "false",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInputFromPipe(); got != tt.want {
				t.Errorf("isInputFromPipe() = %v, want %v", got, tt.want)
			}
		})
	}
}

// dummy test
func Test_runCommand(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runCommand(); (err != nil) != tt.wantErr {
				t.Errorf("runCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--version"})
	Execute("0.0.0")

	out, _ := io.ReadAll(b)
	assert.Equal(t, "v0.0.0\n", string(out))
}

func Test_toHumanReadable(t *testing.T) {
	quarkus, _ := os.Open(filepath.Join("..", "test-quarkus.json"))
	uber, _ := os.Open(filepath.Join("..", "test-uber-zap.json"))
	spring, _ := os.Open(filepath.Join("..", "test-spring-boot.json"))
	dotnet, _ := os.Open(filepath.Join("..", "test-dotnet.json"))

	type args struct {
		r io.Reader
	}

	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name:    "quarkus",
			args:    args{quarkus},
			wantW:   "INFO 2020-07-14T09:38:14.977Z\torg.acme.MyClass\tsample output\n",
			wantErr: false,
		},
		{
			name:    "quarkus",
			args:    args{strings.NewReader("123")},
			wantW:   "123\n",
			wantErr: false,
		},
		{
			name:    "uber zap",
			args:    args{uber},
			wantW:   "information 2020-08-26 12:45:05.143377065 +0000 UTC\tcontroller-runtime.controller\tmsg: Reconciler error\tcontroller: scaledobject-controller\trequest: default/azure-servicebus-queue-scaledobject\n",
			wantErr: false,
		},
		{
			name:    "uber zap", // could not decode
			args:    args{strings.NewReader("123")},
			wantW:   "123\n",
			wantErr: false,
		},
		{
			name:    "spring boot",
			args:    args{spring},
			wantW:   "INFO 2020-07-15T19:09:39.983Z\torg.acme.MyClass\tMy log message\n",
			wantErr: false,
		},
		{
			name:    "spring boot", // could not decode
			args:    args{strings.NewReader("123")},
			wantW:   "123\n",
			wantErr: false,
		},
		{
			name:    "dotnet",
			args:    args{dotnet},
			wantW:   "Information 2021-03-19T13:01:52.734Z\tUserManagementSvc.Authorization.UserRealmRoleAuthorizationHandler\tRole authorization requirement satisfied\n",
			wantErr: false,
		},
		{
			name:    "dotnet", // could not decode
			args:    args{strings.NewReader("123")},
			wantW:   "123\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "uber zap":
				uberZapInput = true
				dotnetInput = false
				springBootInput = false
			case "spring boot":
				uberZapInput = false
				dotnetInput = false
				springBootInput = true
			case "dotnet":
				uberZapInput = false
				dotnetInput = true
				springBootInput = false
			default:
				uberZapInput = false
				dotnetInput = false
				springBootInput = false
			}
			w := &bytes.Buffer{}
			err := toHumanReadable(tt.args.r, w)
			if (err != nil) != tt.wantErr {
				t.Errorf("toHumanReadable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("toHumanReadable() gotW = \n%v, want \n%v", gotW, tt.wantW)
			}
		})
	}
}
