package patient

type GetPatientInfoOut struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Snils      string `json:"snils"`
	Email      string `json:"email"`
}
