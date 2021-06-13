package models

type ApiEventsPage struct {
	Items []struct {
		ID                 int         `json:"id"`
		Lang               string      `json:"lang"`
		Title              string      `json:"title"`
		Text               string      `json:"text"`
		DateFrom           string      `json:"date_from"`
		DateTo             string      `json:"date_to"`
		DateFromTimestamp  int         `json:"date_from_timestamp"`
		DateToTimestamp    int         `json:"date_to_timestamp"`
		Date               string      `json:"date"`
		DateTimestamp      int         `json:"date_timestamp"`
		CreatedAt          string      `json:"created_at"`
		CreatedAtTimestamp int         `json:"created_at_timestamp"`
		UpdatedAt          string      `json:"updated_at"`
		UpdatedAtTimestamp int         `json:"updated_at_timestamp"`
		PublishedAt        string      `json:"published_at"`
		Status             string      `json:"status"`
		Search             int         `json:"search"`
		YaRss              int         `json:"ya_rss"`
		Free               int         `json:"free"`
		IsPowered          int         `json:"is_powered"`
		Label              interface{} `json:"label"`
		OivID              int         `json:"oiv_id"`
		Restriction        struct {
			Age int `json:"age"`
		} `json:"restriction"`
		HasImage            int         `json:"has_image"`
		ActiveFrom          interface{} `json:"active_from"`
		ActiveTo            interface{} `json:"active_to"`
		ActiveFromTimestamp interface{} `json:"active_from_timestamp"`
		ActiveToTimestamp   interface{} `json:"active_to_timestamp"`
		IconID              interface{} `json:"icon_id"`
		Kind                struct {
			ID    string `json:"id"`
			Title string `json:"title"`
			Type  int    `json:"type"`
		} `json:"kind"`
		IsOivPublication int    `json:"is_oiv_publication"`
		EbsID            int    `json:"ebs_id"`
		EbsType          string `json:"ebs_type"`
		EbsTitle         string `json:"ebs_title"`
		EbsAgentUID      string `json:"ebs_agent_uid"`
		Image            struct {
			ID        int64       `json:"id"`
			Title     interface{} `json:"title"`
			Copyright interface{} `json:"copyright"`
			Original  struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"original"`
			Small struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"small"`
			Middle struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"middle"`
			Big struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"big"`
			Download struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"download"`
			Thumb struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"thumb"`
			Show struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"show"`
			OneX1M struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"1x1_m"`
			OneX1S struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"1x1_s"`
			TwoX1B struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"2x1_b"`
			TwoX1M struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"2x1_m"`
			TwoX1S struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"2x1_s"`
			ThreeX1B struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"3x1_b"`
			ThreeX1M struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"3x1_m"`
			ThreeX1S struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"3x1_s"`
			FourX1B struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"4x1_b"`
			FourX1M struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"4x1_m"`
			FourX1S struct {
				ID     int64       `json:"id"`
				Title  interface{} `json:"title"`
				Src    string      `json:"src"`
				Width  int         `json:"width"`
				Height int         `json:"height"`
				Type   string      `json:"type"`
			} `json:"4x1_s"`
		} `json:"image"`
		Organizations []int `json:"organizations"`
	} `json:"items"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
	} `json:"_links"`
	Meta struct {
		TotalCount  int `json:"totalCount"`
		PageCount   int `json:"pageCount"`
		CurrentPage int `json:"currentPage"`
		PerPage     int `json:"perPage"`
	} `json:"_meta"`
}
