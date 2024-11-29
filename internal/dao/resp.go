package dao

type ScanRsp struct {
	Items []interface{} `json:"items"`
	Page  int64         `json:"page"`
	Size  int64         `json:"size"`
	Total int64         `json:"total"`
}

type ExecRsp struct {
	RowsAffected int64         `json:"rows_affected"`
	KeysAffected []interface{} `json:"keys_affected,omitempty"`
}
