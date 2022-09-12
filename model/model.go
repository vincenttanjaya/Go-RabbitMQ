package model

type PublishRequest struct {
	Message Message
	Channel string
}

type Message struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Created int64  `json:"created,omitempty"`
}

type PublishCreateRequest struct {
	Name string `json:"name"`
}

type PublishDeleteRequest struct {
	Uuid string `json:"id"`
}

type PublishResponse struct {
	Message string `json:"message"`
}

type GetAllTodoResponse struct {
	Created []Message `json:"created"`
	Deleted []Message `json:"deleted"`
}
