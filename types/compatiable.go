package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Uint16Array struct {
	Data  []uint16
	Valid bool
}

func (u *Uint16Array) Scan(value interface{}) error {
	if value == nil {
		u.Valid = false
		return nil
	}
	u.Valid = true
	e := json.Unmarshal(value.([]byte), u)
	return e
}

func (u Uint16Array) Value() (driver.Value, error) {
	r, e := json.Marshal(u.Data)
	if e != nil {
		return nil, e
	}
	if string(r) == "null" {
		return nil, e
	}
	return r, e
}

func (u *Uint16Array) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "" || dataStr == "null" {
		u.Valid = false
		return nil
	}
	if err := json.Unmarshal(data, &u.Data); err != nil {
		return err
	}
	u.Valid = true
	return nil
}

func (u Uint16Array) MarshalJSON() ([]byte, error) {
	if !u.Valid {
		return json.Marshal(make([]uint16, 0))
	}
	return json.Marshal(u.Data)
}

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StringArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}

type StringInterfaceMap map[string]interface{}

func (s *StringInterfaceMap) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s StringInterfaceMap) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type DeletedAt sql.NullTime
