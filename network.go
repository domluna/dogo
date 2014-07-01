package dogo

type Networks struct {
	V4 []V4 `json:"v4,omitempty"`
	V6 []V6 `json:"v6,omitempty"`
}

type V4 struct {
	IP      string `json:"ip_address,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Type    string `json:"type,omitempty"`
}

type V6 struct {
	IP      string `json:"ip_address,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Type    string `json:"type,omitempty"`
}
