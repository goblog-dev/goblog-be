package model

func ValidateWhere(where *Where) (whereNew *Where) {
	whereNew = where

	if whereNew == nil {
		whereNew = &Where{
			Parameter: "",
			Values:    []any{},
			Order:     "",
			Limit:     "",
		}
	}

	return
}
