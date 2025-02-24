package ontap

import (
	"net/http"
)

type QtreeGroup struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type QtreeUser struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Qtree struct {
	Resource
	Group         *QtreeGroup `json:"group,omitempty"`
	Id            int         `json:"id"`
	Path          string      `json:"path,omitempty"`
	QosPolicy     *Resource   `json:"qos_policy,omitempty"`
	SecurityStyle string      `json:"security_style,omitempty"`
	Statistics    struct {
		Status  string `json:"status,omitempty"`
		IOPSRaw struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"iops_raw,omitempty"`
		ThroughputRaw struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"throughput_raw,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
	} `json:"statistics,omitempty"`
	Svm             *Resource    `json:"svm,omitempty"`
	UnixPermissions int          `json:"unix_permissions,omitempty"`
	User            *[]QtreeUser `json:"user,omitempty"`
	Volume          *Resource    `json:"volume,omitempty"`
}

type QtreeResponse struct {
	BaseResponse
	Qtrees []Qtree `json:"records,omitempty"`
}

func (c *Client) QtreeGetIter(parameters []string) (qtrees []Qtree, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/qtrees"
	reqParameters := parameters
	for {
		r := QtreeResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, qtree := range r.Qtrees {
			qtrees = append(qtrees, qtree)
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

func (c *Client) QtreeGet(href string, parameters []string) (*Qtree, *RestResponse, error) {
	r := Qtree{}
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
