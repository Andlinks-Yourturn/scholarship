package models



type Student struct {
	name   string
	criteria	Criteria
}

type Criteria struct {
	age     int
	rank	int
	major	string
	amount	int
	GPA		int8
	Score      int64
}