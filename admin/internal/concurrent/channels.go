package concurrent

import "service/admin/demo/internal/domain/users/models"

type StoreChannels struct {
	ChanUsers chan models.Users
	ChanEmail chan string
}
