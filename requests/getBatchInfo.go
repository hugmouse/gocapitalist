package requests

import (
	"strconv"
)

type GetBatchInfo struct {
	BatchID     int
	PageSite    int
	StartOffset int
}

func (r *GetBatchInfo) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	logParams["operation"] = "get_batch_info"
	params["operation"] = "get_batch_info"

	if r.BatchID != 0 {
		logParams["batch_id"] = r.BatchID
		params["batch_id"] = strconv.Itoa(r.BatchID)
	}

	if r.PageSite != 0 {
		logParams["page_size"] = r.PageSite
		params["page_size"] = strconv.Itoa(r.PageSite)
	}

	if r.StartOffset != 0 {
		logParams["start_offset"] = r.StartOffset
		params["start_offset"] = strconv.Itoa(r.StartOffset)
	}

	return params, logParams
}
