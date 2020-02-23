//Package scans userInfo
package radar

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/buger/jsonparser"
)

type User struct {
	Lat               string
	Lon               string
	Id                string //   cards[0]->card_group[i]->user->id
	Screen_name       string // cards[0]->card_group[i]->user->screen_name
	Profile_image_url string // cards[0]->card_group[i]->user->pro..
	Avatar_large      string // cards[0]->card_group[i]->user->ava...
	Verified          bool   // cards[0]->card_group[i]->user->v
	Desc1             string //    cards[0]->card_group[i]->desc1
	Desc2             string //  cards[0]->card_group[i]->desc2
}

type UserInfo struct {
	Id                 string
	Screen_name        string
	Province           string
	City               string
	Location           string
	Description        string
	Url                string
	Profile_image_url  string
	Gender             string
	Followers_count    int64
	Friends_count      int64
	Pagefriends_count  int64
	Statuses_count     int64
	Video_status_count int64
	Favourites_count   int64
	Created_at         string
	Verified           bool
	Avatar_large       string
	Avatar_hd          string
	Bi_followers_count int64
	Lang               string
}

func SearchUser(latitude string, lontitude string, page string) ([]*User, error) {
	requestURL := "https://api.weibo.cn/2/guest/cardlist?lat=" + latitude + ",&lon=" + lontitude + "&page=" + page + "&count=20&containerid=2317120015_111_1"
	//println(requestURL)
	client := &http.Client{}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:72.0) Gecko/20100101 Firefox/72.0")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := make([]*User, 50)
	i := 0

	if resp.StatusCode != 200 {
		return nil, errors.New("Radar Request Blocked By Server")
	}
	cardIndex := "[1]"
	if page != "1" {
		cardIndex = "[0]"
	}

	_, err = jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		userID, err := jsonparser.GetInt(value, "user", "id")
		if err != nil {
			log.Println(err)

		}
		screenName, err := jsonparser.GetString(value, "user", "screen_name")
		if err != nil {
			log.Println(err)
		}
		profileImageUrl, err := jsonparser.GetString(value, "user", "profile_image_url")
		if err != nil {
			log.Println(err)
		}
		avatarLarge, err := jsonparser.GetString(value, "user", "avatar_large")
		if err != nil {
			log.Println(err)
		}
		verified, err := jsonparser.GetBoolean(value, "user", "verified")
		if err != nil {
			log.Println(err)
		}
		desc1, err := jsonparser.GetString(value, "desc1")
		if err != nil {
			log.Println(err)
		}
		desc2, err := jsonparser.GetString(value, "desc2")
		if err != nil {
			log.Println(err)
		}
		users[i] = &User{
			Lat:               latitude,
			Lon:               lontitude,
			Id:                strconv.FormatInt(userID, 10),
			Screen_name:       screenName,
			Profile_image_url: profileImageUrl,
			Avatar_large:      avatarLarge,
			Verified:          verified,
			Desc1:             desc1,
			Desc2:             desc2,
		}

		i++

	}, "cards", cardIndex, "card_group")

	return users[:i], err
}

func RequestUserInfo(uid string) (*UserInfo, error) {
	requestURL := "https://api.weibo.cn/2/profile?uid=" + uid
	//println(requestURL)
	client := &http.Client{}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("UserInfo Request Blocked By Server")
	}

	value := body
	id, err := jsonparser.GetString(value, "userInfo", "idstr")
	if err != nil {
		return nil, err
	}
	screen_name, err := jsonparser.GetString(value, "userInfo", "screen_name")
	if err != nil {
		return nil, err
	}
	province, err := jsonparser.GetString(value, "userInfo", "province")
	if err != nil {
		return nil, err
	}
	city, err := jsonparser.GetString(value, "userInfo", "city")
	if err != nil {
		return nil, err
	}
	location, err := jsonparser.GetString(value, "userInfo", "location")
	if err != nil {
		return nil, err
	}
	description, err := jsonparser.GetString(value, "userInfo", "description")
	if err != nil {
		return nil, err
	}
	url, err := jsonparser.GetString(value, "userInfo", "url")
	if err != nil {
		return nil, err
	}
	gender, err := jsonparser.GetString(value, "userInfo", "gender")
	if err != nil {
		return nil, err
	}
	profile_image_url, err := jsonparser.GetString(value, "userInfo", "profile_image_url")
	if err != nil {
		return nil, err
	}
	followers_count, err := jsonparser.GetInt(value, "userInfo", "followers_count")
	if err != nil {
		return nil, err
	}
	friends_count, err := jsonparser.GetInt(value, "userInfo", "friends_count")
	if err != nil {
		return nil, err
	}

	pagefriends_count, err := jsonparser.GetInt(value, "userInfo", "pagefriends_count")
	if err != nil {
		return nil, err
	}
	statuses_count, err := jsonparser.GetInt(value, "userInfo", "statuses_count")
	if err != nil {
		return nil, err
	}
	video_status_count, err := jsonparser.GetInt(value, "userInfo", "video_status_count")
	if err != nil {
		return nil, err
	}
	favourites_count, err := jsonparser.GetInt(value, "userInfo", "favourites_count")
	if err != nil {
		return nil, err
	}
	created_at, err := jsonparser.GetString(value, "userInfo", "created_at")
	if err != nil {
		return nil, err
	}
	verified, err := jsonparser.GetBoolean(value, "userInfo", "verified")
	if err != nil {
		return nil, err
	}
	avatar_large, err := jsonparser.GetString(value, "userInfo", "avatar_large")
	if err != nil {
		return nil, err
	}
	avatar_hd, err := jsonparser.GetString(value, "userInfo", "avatar_hd")
	if err != nil {
		return nil, err
	}
	bi_followers_count, err := jsonparser.GetInt(value, "userInfo", "bi_followers_count")
	if err != nil {
		return nil, err
	}
	lang, err := jsonparser.GetString(value, "userInfo", "lang")
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		Id:                 id,
		Screen_name:        screen_name,
		Province:           province,
		City:               city,
		Location:           location,
		Description:        description,
		Url:                url,
		Profile_image_url:  profile_image_url,
		Gender:             gender,
		Followers_count:    followers_count,
		Friends_count:      friends_count,
		Pagefriends_count:  pagefriends_count,
		Statuses_count:     statuses_count,
		Video_status_count: video_status_count,
		Favourites_count:   favourites_count,
		Created_at:         created_at,
		Verified:           verified,
		Avatar_large:       avatar_large,
		Avatar_hd:          avatar_hd,
		Bi_followers_count: bi_followers_count,
		Lang:               lang,
	}, nil

}
