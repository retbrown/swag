package data

import (
	"github.com/retbrown/swag/testdata/alias_type/types"
	"time"
)

type TimeContainer struct {
	Name      types.StringAlias `json:"name"`
	Timestamp time.Time         `json:"timestamp"`
	CreatedAt types.DateOnly    `json:"created_at"`
}
