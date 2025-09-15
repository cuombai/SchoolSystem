package admin

func CreateAndLogUser(req CreateUserRequest, adminID string) error {
    return AddUser(req, adminID)
}
