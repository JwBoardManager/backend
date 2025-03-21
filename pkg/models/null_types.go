package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// NullString trata valores NULL corretamente no JSON
type NullString struct {
	sql.NullString
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// NullInt64 trata valores NULL corretamente no JSON
type NullInt64 struct {
	sql.NullInt64
}

func (ns NullInt64) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.Int64)
}

// NullTime trata valores NULL corretamente no JSON
type NullTime struct {
	sql.NullTime
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(nt.Time)
}

// 🔹 Converte um ponteiro `*int64` para `sql.NullInt64`
func ToNullInt64(ptr *int64) sql.NullInt64 {
	if ptr != nil {
		return sql.NullInt64{Int64: *ptr, Valid: true}
	}
	return sql.NullInt64{Valid: false}
}

// 🔹 Converte um ponteiro `*sql.NullInt64` para `int64`
func FromNullInt64(ptr *sql.NullInt64) int64 {
	if ptr != nil && ptr.Valid {
		return ptr.Int64
	}
	return 0
}

// 🔹 Converte um ponteiro `*string` para `sql.NullString`
func ToNullString(ptr *string) sql.NullString {
	if ptr != nil {
		return sql.NullString{String: *ptr, Valid: true}
	}
	return sql.NullString{Valid: false}
}

// 🔹 Converte um ponteiro `*time.Time` para `sql.NullTime`
func ToNullTime(ptr *time.Time) sql.NullTime {
	if ptr != nil {
		return sql.NullTime{Time: *ptr, Valid: true}
	}
	return sql.NullTime{Valid: false}
}
