package ontap

import (
	"net/http"
)

type QuotaReportFiles struct {
	HardLimit *int `json:"hard_limit,omitempty"`
	SoftLimit *int `json:"soft_limit,omitempty"`
	Used      *struct {
		HardLimitPercent int `json:"hard_limit_percent,omitempty"`
		SoftLimitPercent int `json:"soft_limit_percent,omitempty"`
		Total            int `json:"total,omitempty"`
	} `json:"used,omitempty"`
}

type QuotaReportSpace struct {
	HardLimit *int `json:"hard_limit,omitempty"`
	SoftLimit *int `json:"soft_limit,omitempty"`
	Used      *struct {
		HardLimitPercent int `json:"hard_limit_percent,omitempty"`
		SoftLimitPercent int `json:"soft_limit_percent,omitempty"`
		Total            int `json:"total,omitempty"`
	} `json:"used,omitempty"`
}

type QuotaReportGroup struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type QuotaReportQtree struct {
	Links *struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type QuotaReportUser struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type QuotaReport struct {
	Resource
	Files  *QuotaReportFiles  `json:"files,omitempty"`
	Group  *QuotaReportGroup  `json:"groups,omitempty"`
	Index  *int               `json:"index,omitempty"`
	Qtree  *QuotaReportQtree  `json:"qtree,omitempty"`
	Space  *QuotaReportSpace  `json:"space,omitempty"`
	Svm    *Resource          `json:"svm,omitempty"`
	Type   *string            `json:"type,omitempty"`
	Users  *[]QuotaReportUser `json:"users,omitempty"`
	Volume *Resource          `json:"volume,omitempty"`
}

type QuotaReportResponse struct {
	BaseResponse
	QuotaReports []QuotaReport `json:"records,omitempty"`
}

func (c *Client) QuotaReportGetIter(parameters []string) (quotaReports []QuotaReport, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/quota/reports"
	reqParameters := parameters
	for {
		r := QuotaReportResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, quotaReport := range r.QuotaReports {
			quotaReports = append(quotaReports, quotaReport)
		}
		if r.IsPaginate() {
			path = r.GetNextRef()
			reqParameters = []string{}
		} else {
			break
		}
	}
	return
}

func (c *Client) QuotaReportGet(href string, parameters []string) (*QuotaReport, *RestResponse, error) {
	r := QuotaReport{}
	req, err := c.NewRequest("GET", href, parameters, nil)
	if err != nil {
		return nil, nil, err
	}
	res, err := c.Do(req, &r)
	if err != nil {
		return nil, nil, err
	}
	return &r, res, nil
}
