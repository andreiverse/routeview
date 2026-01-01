package peer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Facility represents a physical data center building
type Facility struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

// NetFac represents the link between an ASN and a Facility
type NetFac struct {
	FacID int `json:"fac_id"`
}

type PeeringDBResponse struct {
	Data []NetFac `json:"data"`
}

// CleanASN converts "AS12345" or "12345" to the integer 12345
func CleanASN(asnStr string) (int, error) {
	clean := strings.TrimPrefix(strings.ToUpper(asnStr), "AS")
	return strconv.Atoi(clean)
}

func GetSharedFacilities(asn1Str, asn2Str string) ([]int, error) {
	asn1, err := CleanASN(asn1Str)
	if err != nil {
		return nil, err
	}
	asn2, err := CleanASN(asn2Str)
	if err != nil {
		return nil, err
	}

	// Fetch presence for both ASNs
	facs1, err := fetchNetworkFacilities(asn1)
	if err != nil {
		return nil, err
	}
	facs2, err := fetchNetworkFacilities(asn2)
	if err != nil {
		return nil, err
	}

	// Find Intersection
	var shared []int
	presenceMap := make(map[int]bool)
	for _, id := range facs1 {
		presenceMap[id] = true
	}

	for _, id := range facs2 {
		if presenceMap[id] {
			shared = append(shared, id)
		}
	}

	return shared, nil
}

func fetchNetworkFacilities(asn int) ([]int, error) {
	url := fmt.Sprintf("https://www.peeringdb.com/api/netfac?net__asn=%d", asn)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pdbResp PeeringDBResponse

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	fmt.Printf("RAW PDB RESPONSE for AS%d: %s\n", asn, string(bodyBytes))

	if err := json.Unmarshal(bodyBytes, &pdbResp); err != nil {
		return nil, err
	}

	var ids []int
	for _, nf := range pdbResp.Data {
		ids = append(ids, nf.FacID)
	}
	return ids, nil
}
