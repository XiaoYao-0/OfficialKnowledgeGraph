package item

// 院校
type University struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// 地域
type Area struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// 官员
type Official struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	BirthYear   int    `json:"birth_year"`
	Nationality string `json:"nationality"`
}

// 职位
type Position struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
