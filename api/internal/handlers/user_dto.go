package handlers

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"` // Should be hashed before storing
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Email string `json:"email"`
}
