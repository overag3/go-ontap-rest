package ontap

import (
	"fmt"
	"net/http"
)

type Encryption struct {
	Enabled bool   `json:"enabled"`
	KeyId   string `json:"key_id,omitempty"`
	ReKey   string `json:"rekey,omitempty"`
	State   string `json:"state,omitempty"`
	Status  struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"status,omitempty"`
	Type string `json:"type,omitempty"`
}

type Autosize struct {
	GrowThreshold   int    `json:"grow_threshold"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	Mode            string `json:"mode,omitempty"`
	ShrinkThreshold int    `json:"shrink_threshold"`
}

type Nas struct {
	ExportPolicy    *ExportPolicyRef `json:"export_policy,omitempty"`
	Gid             int              `json:"gid"`
	Path            string           `json:"path,omitempty"`
	SecurityStyle   string           `json:"security_style,omitempty"`
	Uid             int              `json:"uid"`
	UnixPermissions int              `json:"unix_permissions"`
}

type Efficiency struct {
	ApplicationIoSize string         `json:"application_io_size,omitempty"`
	Compaction        string         `json:"compaction,omitempty"`
	CrossVolumeDedupe string         `json:"cross_volume_dedupe,omitempty"`
	Dedupe            string         `json:"dedupe,omitempty"`
	Policy            *NameReference `json:"policy,omitempty"`
	Schedule          string         `json:"schedule,omitempty"`
}

type VolumeSpaceGuarantee struct {
	Honored *bool  `json:"honored,omitempty"`
	Type    string `json:"type,omitempty"`
}

type QosPolicy struct {
	Resource
	MaxThroughputIops int `json:"max_throughput_iops"`
	MaxThroughputMbps int `json:"max_throughput_mbps"`
	MinThroughputIops int `json:"min_throughput_iops"`
	MinThroughputMbps int `json:"min_throughput_mbps"`
}

type Qos struct {
	Policy QosPolicy `json:"policy"`
}

type Quota struct {
	Enabled bool   `json:"enabled"`
	State   string `json:"state,omitempty"`
}

type VolumeSnapshotSettings struct {
	AutodeleteEnabled *bool `json:"autodelete_enabled,omitempty"`
	ReservePercent    *int  `json:"reserve_percent,omitempty"`
	Used              *int  `json:"used,omitempty"`
}

type VolumeSpace struct {
	AfsTotal                          *int  `json:"afs_total,omitempty"`
	Available                         *int  `json:"available,omitempty"`
	CapacityTierFootprint             *int  `json:"capacity_tier_footprint,omitempty"`
	DedupeMetafilesFootprint          *int  `json:"dedupe_metafiles_footprint,omitempty"`
	DedupeMetafilesTemporaryFootprint *int  `json:"dedupe_metafiles_temporary_footprint,omitempty"`
	DelayedFreeFootprint              *int  `json:"delayed_free_footprint,omitempty"`
	ExpectedAvailable                 *int  `json:"expected_available,omitempty"`
	FilesystemSize                    *int  `json:"filesystem_size,omitempty"`
	FilesystemSizeFixed               *bool `json:"filesystem_size_fixed,omitempty"`
	FractionalReserve                 *int  `json:"fractional_reserve,omitempty"`
	FullThresholdPercent              *int  `json:"full_threshold_percent,omitempty"`
	IsUsedStale                       *bool `json:"is_used_stale,omitempty"`
	LargeSizeEnabled                  *bool `json:"large_size_enabled,omitempty"`
	LocalTierFootprint                *int  `json:"local_tier_footprint,omitempty"`
	LogicalSpace                      *struct {
		Enforcement     bool `json:"enforcement"`
		Reporting       bool `json:"reporting"`
		Used            int  `json:"used"`
		UsedByAfs       int  `json:"used_by_afs"`
		UsedBySnapshots int  `json:"used_by_snapshots"`
		UsedPercent     int  `json:"used_percent"`
	} `json:"logical_space,omitempty"`
	Metadata                   *int                    `json:"metadata,omitempty"`
	NearlyFullThresholdPercent *int                    `json:"nearly_full_threshold_percent,omitempty"`
	OverProvisioned            *int                    `json:"over_provisioned,omitempty"`
	OverwriteReserve           *int                    `json:"overwrite_reserve,omitempty"`
	OverwriteReserveUsed       *int                    `json:"overwrite_reserve_used,omitempty"`
	PerformanceTierFootprint   *int                    `json:"performance_tier_footprint,omitempty"`
	PhysicalUsed               *int                    `json:"physical_used,omitempty"`
	PhysicalUsedPercent        *int                    `json:"physical_used_percent,omitempty"`
	Size                       *int                    `json:"size,omitempty"`
	SizeAvailableForSnapshots  *int                    `json:"size_available_for_snapshots,omitempty"`
	Snapshot                   *VolumeSnapshotSettings `json:"snapshot,omitempty"`
	SnapshotSpill              *int                    `json:"snapshot_spill,omitempty"`
	TotalFootprint             *int                    `json:"total_footprint,omitempty"`
	Used                       *int                    `json:"used,omitempty"`
	UsedByAfs                  *int                    `json:"used_by_afs,omitempty"`
	UserData                   *int                    `json:"user_data,omitempty"`
}

type Volume struct {
	Resource
	AccessTimeEnabled *bool      `json:"access_time_enabled,omitempty"`
	Aggregates        []Resource `json:"aggregates,omitempty"`
	Analytics         *struct {
		ScanProgress      int    `json:"scan_progress,omitempty"`
		State             string `json:"state,omitempty"`
		Supported         bool   `json:"supported"`
		UnsupportedReason *struct {
			Code    string `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"unsupported_reason,omitempty"`
	} `json:"analytics,omitempty"`
	Application *Resource `json:"application,omitempty"`
	Autosize    *Autosize `json:"autosize,omitempty"`
	Clone       *struct {
		IsFlexclone          bool     `json:"is_flexclone"`
		ParentSnapshot       Resource `json:"parent_snapshot,omitempty"`
		ParentSvm            Resource `json:"parent_svm,omitempty"`
		ParentVolume         Resource `json:"parent_volume,omitempty"`
		SplitCompletePercent int      `json:"split_complete_percent"`
		SplitEstimate        int      `json:"split_estimate"`
		SplitInitiated       int      `json:"split_initiated"`
	} `json:"clone,omitempty"`
	CloudRetrievalPolicy string `json:"cloud_retrieval_policy,omitempty"`
	Comment              string `json:"comment,omitempty"`
	ConsistencyGroup     *struct {
		Name string `json:"name,omitempty"`
	} `json:"consistency_group,omitempty"`
	ConstituentsPerAggregate *int        `json:"constituents_per_aggregate,omitempty"`
	CreateTime               string      `json:"create_time,omitempty"`
	Efficiency               *Efficiency `json:"efficiency,omitempty"`
	Encryption               *Encryption `json:"encryption,omitempty"`
	ErrorState               *struct {
		HasBadBlocks   bool `json:"has_bad_blocks"`
		IsInconsistent bool `json:"is_inconsistent"`
	} `json:"error_state,omitempty"`
	Files *struct {
		Maximum int `json:"maximum"`
		Used    int `json:"used,omitempty"`
	} `json:"files,omitempty"`
	FlexcacheEndpointType string                `json:"flexcache_endpoint_type,omitempty"`
	Guarantee             *VolumeSpaceGuarantee `json:"guarantee,omitempty"`
	IsObjectStore         *bool                 `json:"is_object_store,omitempty"`
	IsSvmRoot             *bool                 `json:"is_svm_root,omitempty"`
	Language              string                `json:"language,omitempty"`
	Metric                *struct {
		Resource
		Cloud *struct {
			Duration string `json:"duration,omitempty"`
			Iops     *struct {
				Other int `json:"other"`
				Read  int `json:"read"`
				Total int `json:"total"`
				Write int `json:"write"`
			} `json:"iops,omitempty"`
			Latency *struct {
				Other int `json:"other"`
				Read  int `json:"read"`
				Total int `json:"total"`
				Write int `json:"write"`
			} `json:"latency,omitempty"`
			Status    string `json:"status,omitempty"`
			Timestamp string `json:"timestamp,omitempty"`
		} `json:"cloud,omitempty"`
		Duration  string `json:"duration,omitempty"`
		Flexcache *struct {
			CacheMissPercent int    `json:"cache_miss_percent"`
			Duration         string `json:"duration,omitempty"`
			Status           string `json:"status,omitempty"`
			Timestamp        string `json:"timestamp,omitempty"`
		} `json:"flexcache,omitempty"`
		Iops *struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"iops,omitempty"`
		Latency *struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"latency,omitempty"`
		Status     string `json:"status,omitempty"`
		Throughput *struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"throughput,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
	} `json:"metric,omitempty"`
	Movement *struct {
		CutoverWindow        int      `json:"cutover_window"`
		DestinationAggregate Resource `json:"destination_aggregate,omitempty"`
		PercentComplete      int      `json:"percent_complete"`
		State                string   `json:"state,omitempty"`
		TieringPolicy        string   `json:"tiering_policy,omitempty"`
	} `json:"movement,omitempty"`
	Nas                *Nas   `json:"nas,omitempty"`
	Qos                *Qos   `json:"qos,omitempty"`
	QueueForEncryption *bool  `json:"queue_for_encryption,omitempty"`
	Quota              *Quota `json:"quota,omitempty"`
	Size               *int   `json:"size,omitempty"`
	Snaplock           *struct {
		AppendModeEnabled   bool   `json:"append_mode_enabled"`
		AutocommitPeriod    string `json:"autocommit_period,omitempty"`
		ComplianceClockTime string `json:"compliance_clock_time,omitempty"`
		ExpiryTime          string `json:"expiry_time,omitempty"`
		IsAuditLog          string `json:"is_audit_log"`
		LitigationCount     int    `json:"litigation_count"`
		PrivilegedDelete    string `json:"privileged_delete,omitempty"`
		Retention           *struct {
			Default string `json:"default,omitempty"`
			Maximum string `json:"maximum,omitempty"`
			Minimum string `json:"minimum,omitempty"`
		} `json:"retention,omitempty"`
		Type                          string `json:"type,omitempty"`
		UnspecifiedRetentionFileCount int    `json:"unspecified_retention_file_count"`
	} `json:"snaplock,omitempty"`
	Snapmirror *struct {
		IsProtected bool `json:"is_protected"`
	} `json:"snapmirror,omitempty"`
	SnapshotPolicy *Resource    `json:"snapshot_policy,omitempty"`
	Space          *VolumeSpace `json:"space,omitempty"`
	State          string       `json:"state,omitempty"`
	Statistics     *struct {
		Cloud *struct {
			Duration string `json:"duration,omitempty"`
			IopsRaw  *struct {
				Other int `json:"other"`
				Read  int `json:"read"`
				Total int `json:"total"`
				Write int `json:"write"`
			} `json:"iops_raw,omitempty"`
			LatencyRaw *struct {
				Other int `json:"other"`
				Read  int `json:"read"`
				Total int `json:"total"`
				Write int `json:"write"`
			} `json:"latency_raw,omitempty"`
			Status    string `json:"status,omitempty"`
			Timestamp string `json:"timestamp,omitempty"`
		} `json:"cloud,omitempty"`
		FlexcacheRaw *struct {
			CacheMissBlocks       int    `json:"cache_miss_blocks"`
			ClientRequestedBlocks int    `json:"client_requested_blocks"`
			Status                string `json:"status,omitempty"`
			Timestamp             string `json:"timestamp,omitempty"`
		} `json:"flexcache_raw,omitempty"`
		IopsRaw *struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"iops_raw,omitempty"`
		LatencyRaw *struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"latency_raw,omitempty"`
		Status        string `json:"status,omitempty"`
		ThroughputRaw *struct {
			Other int `json:"other"`
			Read  int `json:"read"`
			Total int `json:"total"`
			Write int `json:"write"`
		} `json:"throughput_raw,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
	} `json:"statistics,omitempty"`
	Style   string    `json:"style,omitempty"`
	Svm     *Resource `json:"svm,omitempty"`
	Tiering *struct {
		MinCoolingDays int      `json:"min_cooling_days"`
		ObjectTags     []string `json:"object_tags,omitempty"`
		Policy         string   `json:"policy,omitempty"`
		Supported      bool     `json:"supported"`
	} `json:"tiering,omitempty"`
	Type                  string `json:"type,omitempty"`
	UseMirroredAggregates *bool  `json:"use_mirrored_aggregates,omitempty"`
}

type VolumeResponse struct {
	BaseResponse
	Volumes []Volume `json:"records,omitempty"`
}

func (c *Client) VolumeGetIter(parameters []string) (volumes []Volume, res *RestResponse, err error) {
	var req *http.Request
	path := "/api/storage/volumes"
	reqParameters := parameters
	for {
		r := VolumeResponse{}
		req, err = c.NewRequest("GET", path, reqParameters, nil)
		if err != nil {
			return
		}
		res, err = c.Do(req, &r)
		if err != nil {
			return
		}
		for _, volume := range r.Volumes {
			volumes = append(volumes, volume)
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

func (c *Client) VolumeGet(href string, parameters []string) (*Volume, *RestResponse, error) {
	r := Volume{}
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

func (c *Client) VolumeCreate(volume *Volume, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	var job *Job
	path := "/api/storage/volumes"
	jobLink := JobLinkResponse{}
	if req, err = c.NewRequest("POST", path, parameters, volume); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil {
		if job != nil && job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}

func (c *Client) VolumeModify(href string, volume *Volume, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	var job *Job
	jobLink := JobLinkResponse{}
	if req, err = c.NewRequest("PATCH", href, parameters, volume); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil {
		if job != nil && job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}

func (c *Client) VolumeDelete(href string, parameters []string) (res *RestResponse, err error) {
	var req *http.Request
	var job *Job
	jobLink := JobLinkResponse{}
	if req, err = c.NewRequest("DELETE", href, parameters, nil); err != nil {
		return
	}
	if res, err = c.Do(req, &jobLink); err != nil {
		return
	}
	if job, err = c.JobWaitUntilComplete(jobLink.JobLink.GetRef()); err == nil {
		if job != nil && job.State == "failure" {
			err = fmt.Errorf("Error: REST code=%d, REST message=\"%s\"", job.Code, job.Message)
		}
	}
	return
}
