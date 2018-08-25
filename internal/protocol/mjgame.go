package protocol

type Error int32

const (
	Error_UsernameOrPwdError         Error = 13009
	Error_UsernameEmpty              Error = 13010
	Error_RegistError                Error = 13012
	Error_PwdEmpty                   Error = 13011
	Error_PhoneRegisted              Error = 13016
	Error_PhoneNumberError           Error = 13017
	Error_PwdFormatError             Error = 13018
	Error_PhoneNumberEnpty           Error = 13019
	Error_UserDataNotExist           Error = 14001
	Error_WechatLoingFailReAuth      Error = 14003
	Error_GetWechatUserInfoFail      Error = 14004
	Error_HTTP_LOGIN_EXPIRE          Error = 14006
	Error_HTTP_LOGIN_TOKEN_FAIL      Error = 14007
	Error_HTTP_LOGIN_USERDATA_NIL    Error = 14008
	Error_HTTP_LOGIN_USERID_NIL      Error = 14009
	Error_HTTP_LOGIN_USER_NOT_REGIST Error = 14010
	Error_NotInRoomCannotLeave       Error = 20002
	Error_GameStartedCannotLeave     Error = 20001
	Error_NotYourTurn                Error = 20003
	Error_BuyAlready                 Error = 20004
	Error_NoOperate                  Error = 20005
	Error_NoStarted                  Error = 20006
	Error_GameRoundIllegal           Error = 20008
	Error_GameMaiziIllegal           Error = 20009
	Error_CardValueZero              Error = 20007
	Error_NotInRoom                  Error = 20018
	Error_NotEnoughCoin              Error = 20019
	Error_StartedNotKick             Error = 20023
	Error_CreateRoomFail             Error = 30012
	Error_RoomNotExist               Error = 30016
	Error_RoomFull                   Error = 30018
	Error_NotInPrivateRoom           Error = 30019
	Error_RunningNotVote             Error = 30032
	Error_VotingCantLaunchVote       Error = 30033
	Error_NotVoteTime                Error = 30034
	Error_NameTooLong                Error = 40002
	Error_SexValueRangeout           Error = 40004
	Error_FeedfackError              Error = 40005
	Error_NotEnough_ROOM_CARD        Error = 40007
	Error_NoticeListEnpty            Error = 40008
	Error_DataOutOfRange             Error = 40009
	Error_IpayOrderFail              Error = 62001
	Error_PostboxEmpty               Error = 69001
	Error_PostNotExist               Error = 69002
	Error_AppendixNotExist           Error = 69003
	Error_PrivateRecordEmpty         Error = 68004
)

type RoomData struct {
	RoomID     int32  `json:"room_id"`
	RoomType   int32  `json:"room_type"`
	RoomName   string `json:"room_name"`
	Count      int32  `json:"count"`
	Invitecode string `json:"invitecode"`
	Round      int32  `json:"round"`
	UserID     string `json:"user_id"`
	TotalRound int32  `json:"total_round"`
}

type CCreatePrivateRoom struct {
	RoomName string `json:"room_name"`
	RoomType int32  `json:"room_type"`
	Ante     int32  `json:"ante"`
	Round    int32  `json:"round"`
	Payment  int32  `json:"payment"`
}

type SCreatePrivateRoom struct {
	RoomData RoomData `json:"room_data"`
}

type CEnterSocialRoom struct {
}
