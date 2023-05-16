package handler

import (
	"net/http"
	"testing"

	"github.com/nann-e-backend/api/usecase"
)

func TestHandler_Register(t *testing.T) {
	t.Parallel()

	type fields struct {
		UC usecase.IUsecase
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			h := Handler{
				UC: test.fields.UC,
			}
			if err := h.Register(test.args.w, test.args.r); (err != nil) != test.wantErr {
				t.Errorf("Handler.Register() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
