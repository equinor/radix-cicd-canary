package model

// ApplicationRegistration describe an application
type ApplicationRegistration struct {
	Name         string   `json:"name"`
	Repository   string   `json:"repository"`
	SharedSecret string   `json:"sharedSecret"`
	AdGroups     []string `json:"adGroups"`
	PublicKey    string   `json:"publicKey,omitempty"`
	PrivateKey   string   `json:"privateKey,omitempty"`
}
