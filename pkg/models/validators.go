package models

type SignupUser struct {
	Rollno    string `json:"rollno" validate:"required"`
	Firstname string `json:"firstname" validate:"required,omitempty,min=1,max=30"`
	Lastname  string `json:"lastname" validate:"required,omitempty,min=1,max=30"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,omitempty,min=6"`
}

type LoginUser struct {
	Rollno   string `json:"rollno" validate:"required"`
	Password string `json:"password" validate:"required,omitempty,min=6"`
}

// type SignupUser struct {
// 	Rollno    string `json:"rollno" `
// 	Firstname string `json:"firstname" `
// 	Lastname  string `json:"lastname" `
// 	Email     string `json:"email" `
// 	Password  string `json:"password" `
// }

func (s SignupUser) GetUserDetials() Users {
	return Users{
		Rollno:    s.Rollno,
		Firstname: s.Firstname,
		Lastname:  s.Lastname,
		Email:     s.Email,
		Branch:    s.Rollno[2:5],
		IsAdmin:   false,
		Password:  s.Password,
	}
}
