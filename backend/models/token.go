package Models

type Token struct {
	Email     string
	First_name  string
	Last_name   string
	Birth_date  string
	Gender     string
	Pp_id      string
	Save_Token bool
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}
