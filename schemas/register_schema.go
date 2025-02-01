package schemas

type RegisterSchema struct {
	Username     string `json:"username" validate:"required,usernamevalid"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8,max=32"`
	FirstName    string `json:"first_name" validate:"required,min=2,max=50,alpha"`
	LastName     string `json:"last_name" validate:"required,min=2,max=50,alpha"`
	DateOfBirth  string `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	Gender       string `json:"gender" validate:"required,oneof=male female other"`
	MobileNumber string `json:"mobile_number" validate:"required,len=10,numeric"`
}
