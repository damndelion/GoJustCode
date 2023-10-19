package repo

type (
	UserRepoInterface interface {
		GetAll()
		GetById()
		CreateUser()
		Update()
		Delete()
	}
	UserInfoRepoInterface interface {
		GetAll()
		GetById()
		CreateUser()
		Update()
		Delete()
	}
	UserCredentialsRepoInterface interface {
		GetAll()
		GetById()
		CreateUser()
		Update()
		Delete()
	}
)
