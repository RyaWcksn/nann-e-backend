package interfaces

import "github.com/nann-e-backend/entities"

type IAi interface {
	Register(r entities.RegisterEntity) (resp *entities.RegisterEntityResponse, err error)
}
