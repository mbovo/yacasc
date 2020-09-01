package v1

import (
  "reflect"
  "testing"
)

func TestDownload(t *testing.T) {
  type args struct {
    c *Command
  }
  tests := []struct {
    name string
    args args
    want Result
  }{
    // TODO: Add test cases.
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := Download(tt.args.c); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("Download() = %v, want %v", got, tt.want)
      }
    })
  }
}
