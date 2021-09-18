package config

type CVEResults struct {
	ResultsPerPage uint `json:"resultsPerPage"`
	StartIndex     uint `json:"startIndex"`
	TotalResults   uint `json:"totalResults"`
	Results        struct {
		CVEDataType    string `json:"CVE_data_type"`
		CVEDataFormat  string `json:"CVE_data_format"`
		CVEDataVersion string `json:"CVE_data_version"`
		CVEDataTime    string `json:"CVE_data_timestamp"`
		CVEItems       []struct {
			CVE struct {
				DataType    string `json:"data_type"`
				DataFormat  string `json:"data_format"`
				DataVersion string `json:"data_version"`
				DataMeta    struct {
					CVEID       string `json:"ID"`
					CVEAssigner string `json:"ASSIGNER"`
				} `json:"CVE_data_meta"`
				ProblemType struct {
					ProblemTypeData []struct {
						CWE []struct {
							Language string `json:"lang"`
							Value    string `json:"value"`
						} `json:"description"`
					} `json:"problemtype_data"`
				} `json:"problemtype"`
				References struct {
					ReferenceData []struct {
						URL             string   `json:"url"`
						Name            string   `json:"name"`
						ReferenceSource string   `json:"refsource"`
						Tags            []string `json:"tags"`
					} `json:"reference_data"`
				} `json:"references"`
				Description struct {
					DescriptionData []struct {
						Language string `json:"lang"`
						Value    string `json:"value"`
					} `json:"description_data"`
				} `json:"description"`
			} `json:"cve"`
			Configurations struct {
				CVEDataVersion string `json:"CVE_data_version"`
				Nodes          []struct {
					Operator string   `json:"operator"`
					Children []string `json:"children"`
					CPEMatch []struct {
						Vulnerable        bool     `json:"vulnerable"`
						CPE23URI          string   `json:"cpe23Uri"`
						VersionEndExclude string   `json:"versionEndExcluding"`
						CPEName           []string `json:"cpe_name"`
					} `json:"cpe_match"`
				} `json:"nodes"`
			} `json:"configurations"`
			Impact struct {
				BaseMetricV3 struct {
					CVSSV3 struct {
						Version            string  `json:"version"`
						VectorString       string  `json:"vectorString"`
						AttackVector       string  `json:"attackVector"`
						AttackComplexity   string  `json:"attackComplexity"`
						PrivilegesRequired string  `json:"privilegesRequired"`
						UserInteraction    string  `json:"userInteraction"`
						Scope              string  `json:"scope"`
						Confidentiality    string  `json:"confidentialityImpact"`
						Integrity          string  `json:"integrityImpact"`
						Availability       string  `json:"availibilityImpact"`
						BaseScore          float64 `json:"baseScore"`
						BaseSeverity       string  `json:"baseSeverity"`
					} `json:"cvssV3"`
					Exploitability float64 `json:"exploitabilityScore"`
					ImpactScore    float64 `json:"impactScore"`
				} `json:"baseMetricV3"`
				BaseMetricV2 struct {
					CVSSV2 struct {
						Version               string  `json:"version"`
						VectorString          string  `json:"vectorString"`
						AccessVector          string  `json:"accessVector"`
						AccessComplexity      string  `json:"accessComplexity"`
						Authentication        string  `json:"authentication"`
						ConfidentialityImpact string  `json:"confidentialityImpact"`
						IntegrityImpact       string  `json:"integrityImpact"`
						AvailabilityImpact    string  `json:"availabilityImpact"`
						BaseScore             float64 `json:"baseScore"`
					} `json:"cvssV2"`
					Severity        string  `json:"severity"`
					Exploitability  float64 `json:"exploitabilityScore"`
					ImpactScore     float64 `json:"impactScore"`
					AcInsufInfo     bool    `json:"acInsufInfo"`
					ObtainAllPriv   bool    `json:"obtainAllPrivilege"`
					ObtainUserPriv  bool    `json:"obtainUserPrivilege"`
					ObtainOtherPriv bool    `json:"obtainOtherPrivilege"`
					UserInteraction bool    `json:"userInteractionRequired"`
				} `json:"baseMetricV2"`
			} `json:"impact"`
			PublishedDate string `json:"publishedDate"`
			LastModified  string `json:"lastModifiedDate"`
		} `json:"CVE_Items"`
	} `json:"result"`
}
