package rule_engine

type UserOfferContext struct {
	UserOfferInput  *UserOfferInput
	UserOfferOutput *UserOfferOutput
}

func (uoc *UserOfferContext) RuleName() string {
	return "user_offers"
}

func (uoc *UserOfferContext) RuleInput() RuleInput {
	return uoc.UserOfferInput
}

func (uoc *UserOfferContext) RuleOutput() RuleOutput {
	return uoc.UserOfferOutput
}

// User data attributes
type UserOfferInput struct {
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Age         int     `json:"age"`
	Gender      string  `json:"gender"`
	Income      float64 `json:"income"`
	HasCar      bool    `json:"has_car"`
	HasChildren bool    `json:"has_children"`
	LovesTravel bool    `json:"loves_travel"`
	CarType     string  `json:"has_sport_car"`
}

func (u *UserOfferInput) DataKey() string {
	return "InputData"
}

// Offer output object
type UserOfferOutput struct {
	OfferType string `json:"offer_type"`
}

func (u *UserOfferOutput) DataKey() string {
	return "OutputData"
}

func NewUserOfferContext() *UserOfferContext {
	return &UserOfferContext{
		UserOfferInput:  &UserOfferInput{},
		UserOfferOutput: &UserOfferOutput{},
	}
}
