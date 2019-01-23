package structs

type Link struct {
	Id     int     `json:"id"`
	Url    string  `json:"url"`
	Title  string  `json:"title"`
	Status int     `json:"status"`
	Cost   float64 `json:"cost"`
	Size   int     `json:"size"`
}

type Task struct {
	Id         int      `json:"id"`
	Url        string   `json:"url"`
	Title      string   `json:"title"`
	Size       int      `json:"size"`
	Cost       float64  `json:"float"`
	Image      int `json:"image"`
	Href       int `json:"href"`
	Video      int `json:"video"`
	CreateTime string   `json:"create_time"`
}
