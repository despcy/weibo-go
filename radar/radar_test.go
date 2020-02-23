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
			info, err := RequestUserInfo(user.id)
			if err != nil {
				println(err.Error())
				break
			}
			s := string(info.screen_name + " https://weibo.com/" + info.id + " " + info.gender + " " + info.created_at)

			println(s)

			//每次block要缓 3min
		}
		j++

	}
}
