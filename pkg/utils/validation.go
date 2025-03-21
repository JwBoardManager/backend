package utils

import (
	"database/sql"
	"encoding/json"
)

// NullInt64JSON permite serializar `sql.NullInt64` corretamente para JSON
type NullInt64JSON sql.NullInt64

func (n NullInt64JSON) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil) // Se for inválido, retorna NULL no JSON
	}
	return json.Marshal(n.Int64) // Se for válido, retorna o número normalmente
}
