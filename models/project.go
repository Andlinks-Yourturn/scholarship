package models



type Project struct {
	IpfsId   string
	Address    string
	Pro_Info	Pro_Info
}

type Pro_Info struct {
	Time    string
	Birth	string
	Major	string
	GPA		int8
	Rank	int16
	Score      int64
}