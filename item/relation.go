package item

// 位于
type UniversityArea struct {
	UniversityID int64 `json:"university_id"`
	AreaID       int64 `json:"area_id"`
}

// 属于
type AreaArea struct {
	RootAreaID  int64 `json:"root_area_id"`
	ChildAreaID int64 `json:"child_area_id"`
}

// 毕业于
type OfficialUniversity struct {
	OfficialID   int64 `json:"official_id"`
	UniversityID int64 `json:"university_id"`
}

// 生长于
type OfficialArea struct {
	OfficialID int64 `json:"official_id"`
	AreaID     int64 `json:"area_id"`
}

// 任职(StartYear-EndYear)
type OfficialPosition struct {
	OfficialID int64 `json:"official_id"`
	PositionID int64 `json:"position_id"`
	StartYear  int   `json:"start_year"`
	EndYear    int   `json:"end_year"`
}

// 设于
type PositionArea struct {
	PositionID int64 `json:"position_id"`
	AreaID     int64 `json:"area_id"`
}
