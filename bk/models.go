package bk

// TestPingResponse returns API heartbeat status code
type TestPingResponse struct {
	Status int
}

// TestResponse returns test JSON results from jsonplaceholder.typicode.com
type TestResponse struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// AudiencesListResult is a collection of custom  audiences
type AudiencesListResult struct {
	Audiences []struct {
		Owner struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"owner"`
		Campaigns []interface{} `json:"campaigns"`
		Price     interface{}   `json:"price"`
		Labels    []interface{} `json:"labels"`
		Status    string        `json:"status"`
		Segments  struct {
		} `json:"segments"`
		IDPackages    []interface{} `json:"idPackages"`
		CountryCodes  []string      `json:"countryCodes"`
		Recency       int           `json:"recency"`
		Retargeting   bool          `json:"retargeting"`
		Prospecting   bool          `json:"prospecting"`
		Notes         string        `json:"notes"`
		Name          string        `json:"name"`
		ID            int           `json:"id"`
		Attributes    []interface{} `json:"attributes"`
		SharingType   string        `json:"sharing_type"`
		CampaignCount int           `json:"campaign_count"`
		CreatedAt     string        `json:"created_at"`
		UpdatedAt     string        `json:"updated_at"`
		UsageType     string        `json:"usage_type"`
		DeviceType    string        `json:"device_type"`
	} `json:"audiences"`
	TotalCount int `json:"total_count"`
}

// AudienceDetailResult describes a single audience from BK
type AudienceDetailResult struct {
	Owner struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	Campaigns []interface{} `json:"campaigns"`
	Price     float64       `json:"price"`
	Labels    []string      `json:"labels"`
	Status    string        `json:"status"`
	Segments  struct {
		AND []struct {
			AND []struct {
				OR []struct {
					Cat    int           `json:"cat"`
					Name   string        `json:"name"`
					Path   string        `json:"path"`
					Status string        `json:"status"`
					Price  float64       `json:"price"`
					Freq   []interface{} `json:"freq"`
				} `json:"OR"`
			} `json:"AND"`
		} `json:"AND"`
	} `json:"segments"`
	IDPackages    []interface{} `json:"idPackages"`
	CountryCodes  []string      `json:"countryCodes"`
	Recency       int           `json:"recency"`
	Retargeting   bool          `json:"retargeting"`
	Prospecting   bool          `json:"prospecting"`
	Notes         string        `json:"notes"`
	Name          string        `json:"name"`
	ID            int           `json:"id"`
	Attributes    []interface{} `json:"attributes"`
	SharingType   string        `json:"sharing_type"`
	CampaignCount int           `json:"campaign_count"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	UsageType     string        `json:"usage_type"`
	DeviceType    string        `json:"device_type"`
}

// SegmentReachResult describes a DMP segment reach call
type SegmentReachResult struct {
	UnknownCategories []struct {
		Price  int           `json:"price,omitempty"`
		Reach  int           `json:"reach"`
		Cat    int           `json:"cat"`
		Freq   []interface{} `json:"freq,omitempty"`
		Status string        `json:"status,omitempty"`
	} `json:"unknownCategories"`
	Multiplier int `json:"multiplier"`
	Reach      int `json:"reach"`
	PriceFloor int `json:"priceFloor"`
	AND        []struct {
		Reach      int `json:"reach"`
		PriceFloor int `json:"priceFloor"`
		AND        []struct {
			OR []struct {
				Reach      int `json:"reach"`
				PriceFloor int `json:"priceFloor"`
				AND        []struct {
					Price      int           `json:"price,omitempty"`
					Reach      int           `json:"reach"`
					PriceFloor int           `json:"priceFloor"`
					Cat        int           `json:"cat,omitempty"`
					Freq       []interface{} `json:"freq,omitempty"`
					Status     string        `json:"status,omitempty"`
					Nvars      int           `json:"nvars"`
					OR         []struct {
						Reach      int64 `json:"reach"`
						PriceFloor int   `json:"priceFloor"`
						Cat        int   `json:"cat"`
						Nvars      int   `json:"nvars"`
					} `json:"OR,omitempty"`
				} `json:"AND"`
				Nvars int `json:"nvars"`
			} `json:"OR"`
			Reach      int `json:"reach"`
			PriceFloor int `json:"priceFloor"`
			Nvars      int `json:"nvars"`
		} `json:"AND"`
		Nvars int `json:"nvars"`
	} `json:"AND"`
	CPUTime    int    `json:"_cpuTime"`
	Namespaces []int  `json:"namespaces"`
	Nvars      int    `json:"nvars"`
	Status     string `json:"status"`
}

// TaxonomyResult is a collection of categories in the BK taxonomy
type TaxonomyResult struct {
	NodeList []struct {
		NodeName    string `json:"nodeName"`
		FullPath    string `json:"fullPath"`
		Buyable     bool   `json:"buyable"`
		Size        int    `json:"size"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		Vertical    string `json:"vertical"`
		Leaf        bool   `json:"leaf"`
		NodeID      int    `json:"nodeID"`
		ParentID    int    `json:"parentID"`
	} `json:"nodeList"`
}

// OwnerViewCategoryResult is a collection of BuyerViewCategory types
type OwnerViewCategoryResult struct {
	Count        int  `json:"count"`
	HasMore      bool `json:"hasMore"`
	Items        []*OwnerViewCategory
	Limit        int `json:"limit"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

// OwnerViewCategory describes a buyer view specific  category object
type OwnerViewCategory struct {
	CategoryType                    string        `json:"categoryType"`
	Description                     string        `json:"description"`
	ID                              int           `json:"id"`
	IsCountableFlag                 bool          `json:"isCountableFlag"`
	IsExplicitPriceFloorFlag        bool          `json:"isExplicitPriceFloorFlag"`
	IsForNavigationOnlyFlag         bool          `json:"isForNavigationOnlyFlag"`
	IsIncludeForAnalyticsFlag       bool          `json:"isIncludeForAnalyticsFlag"`
	IsLeafFlag                      bool          `json:"isLeafFlag"`
	IsMutuallyExclusiveChildrenFlag bool          `json:"isMutuallyExclusiveChildrenFlag"`
	IsPublicFlag                    bool          `json:"isPublicFlag"`
	Links                           []interface{} `json:"links"`
	Name                            string        `json:"name"`
	NamespaceID                     int           `json:"namespaceId"`
	OwnershipType                   string        `json:"ownershipType"`
	ParentCategory                  struct {
		ID int `json:"id"`
	} `json:"parentCategory"`
	Partner struct {
		ID int `json:"id"`
	} `json:"partner"`
	PathFromRoot struct {
		Ids   []int    `json:"ids"`
		Names []string `json:"names"`
	} `json:"pathFromRoot"`
	PriceFloor float32 `json:"priceFloor"`
	SoftFloor  int     `json:"softFloor"`
	SortOrder  int     `json:"sortOrder"`
	Status     string  `json:"status"`
	Vertical   struct {
		Name string `json:"name"`
	} `json:"vertical"`
	VisibilityStatus string `json:"visibilityStatus"`
}

// SiteResult is a collection of Site types
type SiteResult struct {
	Sites      []*Site
	TotalCount int `json:"total_count"`
}

// Site describes a site object
type Site struct {
	AllowedBuyers []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"allowed_buyers"`
	AnalyticsOnly             string      `json:"analytics_only"` // REVIEW: should be bool
	BlockedCountries          []string    `json:"blocked_countries"`
	CreatedAt                 string      `json:"created_at"`                  // REVIEW: should be time
	UpdatedAt                 string      `json:"updated_at"`                  // REVIEW: should be time
	DataTransferBoostEnabled  string      `json:"data_transfer_boost_enabled"` // REVIEW: should be bool
	DataTransferBoostInterval int         `json:"data_transfer_boost_interval"`
	DataTransferLimit         int         `json:"data_transfer_limit"`
	GroupID                   int         `json:"group_id"`
	ID                        int         `json:"id"`
	Labels                    interface{} `json:"labels"`
	Name                      string      `json:"name"`
	Notes                     string      `json:"notes"`
	Partner                   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"partner"`
	PrivateData      string `json:"private_data"`
	TransactionScope string `json:"transaction_scope"`
	Type             int    `json:"type"`
}
