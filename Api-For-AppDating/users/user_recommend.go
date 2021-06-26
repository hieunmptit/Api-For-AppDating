package users

import (
	"Api/dto"
	"Api/handler"
	"Api/interfaces"
)

func GetUsername(id uint) *interfaces.User {
	db := handler.ConnectDB()
	user := &interfaces.User{}
	if db.Where("id = ? ", id).First(&user).RecordNotFound() {
		return nil
	}
	db.Close()
	return user
}

func Recommendation(userId uint) (dto.GetListUserResponse, error) {
	profile := GetProfile(userId)
	recomnendlist := make([]dto.ResponseRecommend, 0)
	for i := 0; i < 100000; i++ {
		k := uint(i)
		if GetProfile(k).Age == profile.PAge && GetProfile(k).Location == profile.Location && GetProfile(k).Gender == profile.Gender && k != userId && profile.Status[k] == 0 {
			recomnendlist = append(recomnendlist, dto.ResponseRecommend{
				Username: GetUsername(k).Username,
				Name:     GetProfile(k).Name,
				Gender:   GetProfile(k).Gender,
				Location: GetProfile(k).Location,
			})
		}
	}
	pageList := make([]dto.Page, 0)
	if len(recomnendlist) > 5 {

		for i := 0; i < len(recomnendlist)-50; i += 50 {
			p := dto.Page{
				NUser:         50,
				Recomnendlist: recomnendlist[i : i+50],
			}
			pageList = append(pageList, p)

		}
		p := dto.Page{
			NUser:         (len(recomnendlist)-1)/50 + 1,
			Recomnendlist: recomnendlist[(len(recomnendlist) - len(recomnendlist)%50):],
		}
		pageList = append(pageList, p)
	} else {
		p := dto.Page{
			NUser:         len(recomnendlist),
			Recomnendlist: recomnendlist[0:],
		}
		pageList = append(pageList, p)
	}

	return dto.GetListUserResponse{
		NPage:    len(pageList),
		PageList: pageList,
	}, nil
}
