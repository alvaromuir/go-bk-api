package bk

import (
	"fmt"
	"log"
)

// Ping returns a status from a BK API call
func Ping() {
	resp, err := PingService()
	CheckError(err, "ERROR: %s")
	fmt.Printf("BK Status: %v \n", resp.Status)
}

// TableAudiences returns a list of audiences from BK, determined by params string
func TableAudiences(params string) {
	resp, rslt, err := Audiences(params)
	CheckError(err, "ERROR: %s")
	if resp.StatusCode == 200 {
		fmt.Printf("%s %6s %7v %s\n",
			"Status", "AudienceID", "Device", "Name")
		for _, item := range rslt.Audiences {
			fmt.Printf("%s %6v %7v %s\n", item.Status, item.ID, item.DeviceType, item.Name)
		}
		fmt.Println("----------------------------------------------------------------")
		fmt.Printf("Total: %v\n", len(rslt.Audiences))
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}

// DetailAudience returns a specific from BK, determined by audienceID
func DetailAudience(audienceID string) {
	resp, rslt, err := Audience(audienceID)
	CheckError(err, "ERROR: %s")
	if resp.StatusCode == 200 {
		fmt.Println(rslt)
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}

// GetSegmentsReach provide reach of segments, denoted by queryArgs and a deviceTypeID
func GetSegmentsReach(segmentsQuery string, deviceTypeID string) {
	resp, rslt, err := Segments(segmentsQuery, deviceTypeID)
	CheckError(err, "ERROR: %s")
	if resp.StatusCode == 200 {
		fmt.Println(rslt)
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}

// TableTaxonomy returns a list of categories from BK, determined by parentID
func TableTaxonomy(queryArgs string) {
	resp, rslt, err := Taxonomy(queryArgs)
	CheckError(err, "ERROR: %s")
	if resp.StatusCode == 200 {
		fmt.Printf("%s \t%13s \t%12s\n",
			"ParentID", "NodeID", "Name")
		for _, item := range rslt.NodeList {
			fmt.Printf("%d \t%-10d \t%s\n",
				item.ParentID, item.NodeID, item.NodeName)
		}
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}

// TableSites returns a list of BK tagged websites associted with this account
func TableSites() {
	resp, rslt, err := Sites()
	if err != nil {
		log.Fatalf("An error has occured calling an API.")
	}
	if resp.StatusCode == 200 {
		for _, item := range rslt.Sites {
			fmt.Printf("#%d \t%s \t%s\n",
				item.ID, item.UpdatedAt, item.Name)
		}
	} else {
		log.Fatalf("ERROR - returned: %s", resp.Status)
	}
}
