package ontap

import (
	"net/http"
)

// QuotaRuleSpace defines space limits for a quota rule
type QuotaRuleSpace struct {
	HardLimit *int `json:"hard_limit,omitempty"`
	SoftLimit *int `json:"soft_limit,omitempty"`
}

// QuotaRuleFiles defines file limits for a quota rule
type QuotaRuleFiles struct {
	HardLimit *int `json:"hard_limit,omitempty"`
	SoftLimit *int `json:"soft_limit,omitempty"`
}

// QuotaRuleQtree defines qtree information for a quota rule
type QuotaRuleQtree struct {
	Links *struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// QuotaRuleUser defines a user in a quota rule
type QuotaRuleUser struct {
	Name string `json:"name,omitempty"`
}

// QuotaRuleGroup defines a group in a quota rule
type QuotaRuleGroup struct {
	Name string `json:"name,omitempty"`
}

// QuotaRule represents a quota rule in ONTAP
type QuotaRule struct {
	Resource
	Svm         *Resource        `json:"svm,omitempty"`
	Volume      *Resource        `json:"volume,omitempty"`
	Type        *string          `json:"type,omitempty"` // "user", "group", or "tree"
	Users       *[]QuotaRuleUser `json:"users,omitempty"`
	Group       *QuotaRuleGroup  `json:"group,omitempty"`
	Qtree       *QuotaRuleQtree  `json:"qtree,omitempty"`
	UserMapping *string          `json:"user_mapping,omitempty"`
	Space       *QuotaRuleSpace  `json:"space,omitempty"`
	Files       *QuotaRuleFiles  `json:"files,omitempty"`
}

// QuotaRuleResponse represents the response for quota rule operations
type QuotaRuleResponse struct {
	BaseResponse
	QuotaRules []QuotaRule `json:"records,omitempty"`
	Job        *Resource   `json:"job,omitempty"`
}

// QuotaRuleGetIter retrieves all quota rules with pagination
func (c *Client) QuotaRuleGetIter(parameters []string) (quotaRules []QuotaRule, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/quota/rules"
	reqParameters := parameters
	for {
		r := QuotaRuleResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, quotaRule := range r.QuotaRules {
			quotaRules = append(quotaRules, quotaRule)
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

// QuotaRuleGet retrieves a specific quota rule by its UUID
func (c *Client) QuotaRuleGet(uuid string, parameters []string) (*QuotaRule, *RestResponse, error) {
	path := "/api/storage/quota/rules/" + uuid
	r := QuotaRule{}
	req, err := c.NewRequest("GET", path, parameters, nil)
	if err != nil {
		return nil, nil, err
	}
	res, err := c.Do(req, &r)
	if err != nil {
		return nil, nil, err
	}
	return &r, res, nil
}

// QuotaRuleCreate creates a new quota rule
func (c *Client) QuotaRuleCreate(quotaRule *QuotaRule, parameters []string) (*QuotaRule, *RestResponse, error) {
	path := "/api/storage/quota/rules"
	r := QuotaRuleResponse{}
	req, err := c.NewRequest("POST", path, parameters, quotaRule)
	if err != nil {
		return nil, nil, err
	}
	res, err := c.Do(req, &r)
	if err != nil {
		return nil, nil, err
	}
	if len(r.QuotaRules) > 0 {
		return &r.QuotaRules[0], res, nil
	}
	return nil, res, nil
}

// QuotaRuleModify modifies an existing quota rule
func (c *Client) QuotaRuleModify(uuid string, quotaRule *QuotaRule, parameters []string) (*RestResponse, error) {
	path := "/api/storage/quota/rules/" + uuid
	req, err := c.NewRequest("PATCH", path, parameters, quotaRule)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// QuotaRuleDelete deletes a quota rule
func (c *Client) QuotaRuleDelete(uuid string, parameters []string) (*RestResponse, error) {
	path := "/api/storage/quota/rules/" + uuid
	req, err := c.NewRequest("DELETE", path, parameters, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
