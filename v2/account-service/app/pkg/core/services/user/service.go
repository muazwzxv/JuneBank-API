package user

import (
	pub "account-service/app/events/publish"
	"account-service/app/pkg/core/domain"
)

var (
	USERCREATED = "user created"
)

func (svr *userService) Create(createUser domain.CreateUser) error {
	user, err := svr.userRepository.Save(createUser)
	if err != nil {
		return err
	}

	userByte, err := user.ToJsonStr()
	if err != nil {
		return err
	}

	payload := &pub.Payload{
		Data:  userByte,
		Event: USERCREATED,
	}

	// Do this asyncly
	go svr.Publish(payload, []string{"KEY"})

	return nil
}

func (svr *userService) Get(id uint64) (*domain.User, error) {
	return svr.userRepository.GetByID(id)
}

func (svr *userService) Update(id uint64, update domain.CreateUser) (*domain.User, error) {
	return svr.userRepository.Update(id, update)
}

func (svr *userService) Delete(id uint64) (*domain.User, error) {
	return svr.userRepository.DeleteByID(id)
}

func (svr *userService) IsUserExist(id uint64) (bool, error) {
	return svr.userRepository.IsUserExist(id)
}
