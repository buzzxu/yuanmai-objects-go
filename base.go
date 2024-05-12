package main

type (
	R[T any] struct {
		Code    int    `json:"code"`
		Error   bool   `json:"error"`
		Key     string `json:"key,omitempty"`
		Message string `json:"message,omitempty"`
		Data    T      `json:"data,omitempty"`
	}
	Id[ID comparable] struct {
		ID        ID           `json:"id" gorm:"column:id;primaryKey;autoIncrement;not null" mapstructure:"id"`
		Creater   string       `json:"creater,omitempty" gorm:"column:creater" `
		Updater   string       `json:"updater,omitempty" gorm:"column:updater"`
		Deleted   bool         `json:"deleted" gorm:"column:deleted;default:false"`
		CreatedAt TimeDateTime `json:"createdAt,omitempty" gorm:"column:created_at" mapstructure:"createdAt"`
		UpdatedAt TimeDateTime `json:"updatedAt,omitempty" gorm:"column:updated_at" mapstructure:"updatedAt"`
	}

	District struct {
		Name     string     `json:"name"`
		Level    string     `json:"level"`
		AdCode   string     `json:"adCode"`
		CityCode string     `json:"cityCode"`
		Geo      []string   `json:"geo"`
		Children []District `json:"children"`
	}
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Province  string  `json:"province"`
		City      string  `json:"city"`
		County    string  `json:"county"`
		District  string  `json:"district"`
	}
	Region struct {
		ProvinceId string   `json:"provinceId"`
		Province   string   `json:"province"`
		CityId     string   `json:"cityId"`
		City       string   `json:"city"`
		County     string   `json:"county"`
		CountyId   string   `json:"countyId"`
		District   string   `json:"district"`
		Town       string   `json:"town"`
		TownId     string   `json:"townId"`
		Village    string   `json:"village"`
		Geo        []string `json:"geo"`
		Address    string   `json:"address"`
	}
)

// R_ 自定义成功
func R_[T any](data T) *R[T] {
	return &R[T]{
		Code:  200,
		Error: false,
		Data:  data,
	}
}

// R__ 自定义错误
func R__[T any](code int, message string) *R[T] {
	return &R[T]{
		Code:    code,
		Error:   true,
		Message: message,
	}
}

func (d *District) Find(cityCode string) []District {
	if d.Children != nil {
		for _, child := range d.Children {
			if child.CityCode == cityCode {
				return child.Children
			}
		}
	}
	return []District{}
}
