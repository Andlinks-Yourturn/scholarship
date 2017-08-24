package models



type Ipfs struct {
	IpfsId   string
	Address    string
	Info	Info
}

type Info struct {
	firstName	string
	LastName	string
	Birth	string
	Major	string
	GPA		int8
	Rank	int16
	Score      int64
}
