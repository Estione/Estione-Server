package usecases

import "estione/internal/estione/domain"

type ChatRepository interface {
	CreateChat() *domain.Chat
	UpdateChat()
}
