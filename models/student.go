package models



type Student struct {
	IpfsId   string
	Address    string
	Stu_Info	Info
}

type Stu_Info struct {
	firstName	string
	LastName	string
	Birth	string
	Major	string
	GPA		int8
	Rank	int16
	Score      int64
}