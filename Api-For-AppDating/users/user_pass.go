package users

func Pass(id uint, pid uint) string {
	k := int(pid)
	user := GetProfile(id)
	user.Status[k] = 2
	return string("Passed")
}
