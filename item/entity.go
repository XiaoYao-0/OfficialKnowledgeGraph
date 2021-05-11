package item

// 院校
type University struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// 地域
// Level 0: 国家, 1: 省级, 2: 市级, 3: 县级
type Area struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// 官员 0:男 1:女 2:未知
type Official struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	BirthYear   int    `json:"birth_year"`
	Nationality string `json:"nationality"`
}

// 职位
// Level 0: 国家级正职, 1: 国家级副职, 2: 省部级正职, 3: 省部级副职, 4: 厅局级正职, 5: 厅局级副职
// Level 6: 县处级正职, 7: 县处级副职, 8: 乡科级正职, 9: 乡科级副职
type Position struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
