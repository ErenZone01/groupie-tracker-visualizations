package main

type categorie struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Search struct{
	Artists   []artists
}

type artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type locations struct {
	Id        	int      `json:"id"`
	Locations 	[]string `json:"locations"`
	Date 		string
}

type dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type IndexLocation struct {
	Index []locations `json:"index"`
}

type relations struct {
	Id             int `json:"id"`
	DatesLocations map[string][]string
}

type IndexDates struct{
	Index []dates `json:"index"`
}

type IndexRel struct{
	Index []relations
}

type AllItem struct {
	Artists   []artists
	Location  []locations
	Dates     []dates
	Relations []RelaTab
	Id        int
	Searchs    Search
}

type RelaTab struct {
	DatesLocations []string
}
