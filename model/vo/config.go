package vo

type ConfigVo struct {
	BaseVo
	Key    string `json:"key"`
	Value  string `json:"value"`
	Remark string `json:"remark"`
}
