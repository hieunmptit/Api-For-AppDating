package users

func Like(id uint, pid uint) string {
	k := int(pid)
	z := int(id)
	user1 := GetProfile(id)
	user1.Status[k] = 1
	user2 := GetProfile(pid)
	if user1.Status[k] == user2.Status[z] && user2.Status[z] == 1 {
		return string("matched")
	}
	return string("liked")
}
