package types

// DonationRequest - ожидаемый json пожертвования
type DonationRequest struct {
	Donation     *Donation     `json:"donation"`
	Character    *Character    `json:"character"`
	Subscription *Subscription `json:"subscriptions"`
}

// Donation - ожидаемы json пожертвования
type Donation struct {
	Amount                float64 `json:"amount"`
	DirectionID           int     `json:"directionId"`
	UserEmail             string  `json:"userEmail"`
	RegularPaymentEnabled bool    `json:"areRegularPaymentsEnabled"`
}

// Character - ожидаемый json персонализации пожертвования
type Character struct {
	Gender       int    `json:"characterGender"`
	Name         string `json:"name"`
	CostumeColor string `json:"costumeColor"`
	HairColor    string `json:"hairColor"`
}

type Subscription struct {
	GetReport     bool `json:"getReport"`
	TrackProgress bool `json:"trackProgress"`
}

// CompanyInfo - json ответа на получение информации по компании
type CompanyInfo struct {
	TerminationAmount float64 `json:"terminationAmount"`
	CollectedAmount   float64 `json:"collectedAmount"`
	DonationCount     int     `json:"donationCount"`
	DayRemains        float64 `json:"dayRemains"`
}
