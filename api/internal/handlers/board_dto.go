package handlers

type CreateBoardRequest struct {
	Name string `json:"name"`
}

type UpdateBoardRequest struct {
	Name string `json:"name"`
}

type BoardResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
