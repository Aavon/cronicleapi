package cronicleapi

// # common api resp
type CommonRsp struct {
	// float64 or string
	Code interface{}    `json:"code"`
	// optional
	Description string  `json:"description"`
}

// # get_event request
type GetEventReq struct {
	// event id
	Id    string `json:"id"`
	// event title(N)
	Title string `json:"title"`
}

type GetEventRsp struct {
	// float64 or string
	Code interface{}    `json:"code"`
	// optional
	Description string  `json:"description"`
	// Event
	Event *EventObject  `json:"event"`
}

// # create_event response
type CreateEventRsp struct {
	// float64 or string
	Code interface{}    `json:"code"`
	Id   string         `json:"id"`
	// optional
	Description string  `json:"description"`
}

// # delete_event request
type DeleteEventReq struct {
	// event id
	Id    string `json:"id"`
}

// # run_event response
type RunEventRsp struct {
	CommonRsp
	Ids   []string `json:"ids"`
}