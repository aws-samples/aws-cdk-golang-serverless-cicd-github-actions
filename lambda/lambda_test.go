package main

import (
	"context"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	tests := []struct {
		name    string
		event   *MyEvent
		want    string
		wantErr bool
	}{
		{
			name:    "Valid input",
			event:   &MyEvent{Name: "John"},
			want:    "Hello John!",
			wantErr: false,
		},
		{
			name:    "Nil event",
			event:   nil,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(context.Background(), tt.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if *got != tt.want {
				t.Errorf("HandleRequest() = %v, want %v", *got, tt.want)
			}
		})
	}
}

