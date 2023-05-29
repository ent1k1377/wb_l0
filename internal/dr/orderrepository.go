package dr

import "github.com/ent1k1377/wb_l0/internal/model"

type OrderRepository struct {
	dr *DR
}

func (r OrderRepository) Create(o *model.Order) (*model.Order, error) {
	return nil, nil
}
