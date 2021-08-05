package common

func NewPageResp(totalRows, limit int32) *PageResp {
	resp := &PageResp{TotalRows: totalRows}
	if limit > 0 {
		resp.TotalPage = resp.TotalRows / limit
		if resp.TotalRows%limit > 0 {
			resp.TotalPage++
		}
	}
	return resp
}
