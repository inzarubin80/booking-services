package models

type MessageResponse struct {
    Message      string `json:"message"`
    Error        bool  `json:"Error"` 
}