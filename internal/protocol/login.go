package protocol

//LoginToGameServerRequest login request
type LoginToGameServerRequest struct {
	Name    string `json:"name"`
	UID     string `json:"uid"`
	HeadURL string `json:"headUrl"`
	Sex     int    `json:"sex"` //[0]未知 [1]男 [2]女
	FangKa  int    `json:"fangka"`
	IP      string `json:"ip"`
}

// JoinResponse represents the result of joining lobby
type JoinResponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}
