package types

// Donation - ожидаемый json пожертвования
type Donation struct {
	ComicsID  int     `json:"comicsId"`
	Amount    float64 `json:"donationAmount"`
	UserEmail string  `json:"userEmail"`
	UserName  string  `json:"userName"`

	RegularPaymentEnabled bool `json:"areRegularPaymentsEnabled"`
	SubscribedToReport    bool `json:"isSubscribedToGetReport"`
	SubscribedToProgress  bool `json:"isSubscribedToTrackProgress"`

	Personalisation *Personalisation `json:"personalisation"`
}

// Personalisation - ожидаемый json персонализации пожертвования
type Personalisation struct {
	Gender        int    `json:"characterGender"`
	Name          string `json:"characterName"`
	CostumeColor  string `json:"costumeColor"`
	SignboardText string `json:"previewName"`
}

// CompanyInfo - json ответа на получение информации по компании
type CompanyInfo struct {
	TerminationAmount float64 `json:"terminationAmount"`
	CollectedAmount   float64 `json:"collectedAmount"`
	DayRemains        float64 `json:"dayRemains"`
}
