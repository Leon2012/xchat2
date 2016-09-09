package protocol

type LoginArg struct {
	Token    string
	DeviceId string
}

type LoginReply struct {
	Uid      int64
	Nickname string
}
