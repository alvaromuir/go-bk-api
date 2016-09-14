package main

import (
	"fmt"
	"time"

	"github.com/alvaromuir/go-bk-api/bk"
)

func main() {
	start := time.Now()

	bk.Ping()
	// bk.TableAudiences("status=shared")
	// bk.DetailAudience("12345")
	// bk.TableTaxonomy("parentId=123456")
	// bk.TableSites()
	// bk.GetSegmentsReach("{\"AND\":[{\"AND\": [{\"OR\": [{\"cat\": 17}]}]}]}", "0")

	elapsed := time.Since(start)
	fmt.Println("")
	fmt.Printf("Operation took %s\n", elapsed-(elapsed%10000000))
}
