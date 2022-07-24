package neisgo

// NewNeisClient returns new neis api wrapper client.
// Key is optional. If key is missing(""), it will use sample key(5 data only).
func NewNeisClient(key string) *NeisClient {
	return &NeisClient{
		http: newHTTPClient(key),
	}
}

// GetSchoolInfo returns Schools
// All arguments are optional. If arguments are missing(""), it will return all schools.
func (c NeisClient) GetSchoolInfo(officeCode string, schoolCode string, schoolName string, schoolType string, location string, foundationName string) ([]School, error) {
	var schools []School

	response, err := c.http.get("/schoolInfo", &Parameters{"ATPT_OFCDC_SC_CODE": officeCode, "SD_SCHUL_CODE": schoolCode, "SCHUL_NM": schoolName, "SCHUL_KND_SC_NM": schoolType, "LCTN_SC_NM": location, "FOND_SC_NM": foundationName})
	if err != nil {
		return nil, err
	}

	if err := getError(response); err != nil {
		return nil, err
	}

	for _, value := range response.GetArray("schoolInfo", "1", "row") {
		schools = append(schools, School{
			OfficeCode:       string(value.GetStringBytes("ATPT_OFCDC_SC_CODE")),
			OfficeName:       string(value.GetStringBytes("ATPT_OFCDC_SC_NM")),
			Code:             string(value.GetStringBytes("SD_SCHUL_CODE")),
			Name:             string(value.GetStringBytes("SCHUL_NM")),
			EnglishName:      string(value.GetStringBytes("ENG_SCHUL_NM")),
			Type:             string(value.GetStringBytes("SCHUL_KND_SC_NM")),
			ZipCode:          string(value.GetStringBytes("ORG_RDNZC")),
			Address:          string(value.GetStringBytes("ORG_RDNMA")),
			DetailedAddress:  string(value.GetStringBytes("ORG_RDNDA")),
			TelNumber:        string(value.GetStringBytes("ORG_TELNO")),
			Homepage:         string(value.GetStringBytes("HMPG_ADRES")),
			Gender:           string(value.GetStringBytes("COEDU_SC_NM")),
			FaxNumber:        string(value.GetStringBytes("ORG_FAXNO")),
			HighSchoolType:   string(value.GetStringBytes("HS_SC_NM")),
			SPHighSchoolType: string(value.GetStringBytes("SPCLY_PURPS_HS_ORD_NM")),
			FoundationDate:   string(value.GetStringBytes("FOND_YMD")),
			FoundationDay:    string(value.GetStringBytes("FOAS_MEMRD")),
		})
	}

	return schools, nil
}
