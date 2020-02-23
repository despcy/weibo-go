package radar

import (
	"os"
	"strconv"
	"testing"
)

func TestRadar(t *testing.T) {
	j := 1
	for {
		users, err := SearchUser("38.080000", "114.01234", strconv.Itoa(j))
		if err != nil {
			println(err.Error())
			os.Exit(0)
		}

		for _, user := range users {
			info, err := RequestUserInfo(user.Id)
			if err != nil {
				println(err.Error())
				break
			}
			s := string(info.Screen_name + " https://weibo.com/" + info.Id + " " + info.Gender + " " + info.Created_at)

			println(s)

			//每次block要缓 3min
		}
		j++

	}
}
