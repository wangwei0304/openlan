package schema

type IPSecTunnel struct {
	Left      string `json:"local"`
	LeftId    string `json:"localid"`
	LeftPort  int    `json:"localport"`
	Right     string `json:"remote"`
	RightId   string `json:"remoteid"`
	RightPort int    `json:"remoteport"`
	Transport string `json:"transport"`
	Secret    string `json:"secret"`
}
