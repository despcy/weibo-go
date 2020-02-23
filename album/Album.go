//package album get photo info from a uid
package album

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

type AlbumClient struct {
	Uid        string
	HasNext    bool
	CurSinceId string
}

//cards->[1]->card_group->[1]->pics
type PhotoInfo struct {
	Pic_small  string
	Pic_middle string
	Pic_big    string
	Pic_mw2000 string
	Photo_tag  int64
	Pic_id     string
	Video      string
	Pic_type   string
	Blog       *MBlog
}

type MBlog struct {
	Id                    string
	Mid                   string
	Text                  string
	IsLongText            bool
	AuthorId              string
	AuthorScreenName      string
	AuthorProfileImageURL string
}

//https://api.weibo.cn/2/guest/cardlist?containerid=2318261669879400_-_mobile_profile_album_-_Index&since_id=
func NewAlbumClient(uid string) *AlbumClient {
	return &AlbumClient{
		Uid:        uid,
		HasNext:    true,
		CurSinceId: "",
	}
}
func (c *AlbumClient) ResetPointer() {
	c.CurSinceId = ""
	c.HasNext = true
}

//for c.hasNext ...
func (c *AlbumClient) RequestNextPage() ([]*PhotoInfo, error) {
	if !c.HasNext {
		return nil, errors.New("Album Page End")

	}
	requestURL := "https://api.weibo.cn/2/guest/cardlist?containerid=231826" + c.Uid + "_-_mobile_profile_album_-_Index&since_id=" + c.CurSinceId
	println(requestURL)
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

	if resp.StatusCode != 200 {
		return nil, errors.New("Photo Info Request Blocked By Server")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	info, err := c.processPhotoJson(body)

	if err != nil {
		return nil, err
	}
	//println(string(body))
	newsinceId, datatype, _, err := jsonparser.Get(body, "cardlistInfo", "since_id")
	if err != nil {
		return nil, err
	}
	if datatype.String() != "string" {
		c.HasNext = false
		c.CurSinceId = ""
	} else {
		c.HasNext = true
		c.CurSinceId = string(newsinceId)
	}
	return info, nil

}

func (c *AlbumClient) processPhotoJson(body []byte) ([]*PhotoInfo, error) {
	photos := make([]*PhotoInfo, 50)
	i := 0
	albumIndex := "[0]"
	if c.CurSinceId == "" {
		albumIndex = "[1]"
	}
	//println(string(body))
	_, err := jsonparser.ArrayEach(body, func(val []byte, dataType jsonparser.ValueType, offset int, err error) {
		_, err = jsonparser.ArrayEach(val, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			pic_small, err := jsonparser.GetString(value, "pic_small")
			pic_middle, err := jsonparser.GetString(value, "pic_middle")
			pic_big, err := jsonparser.GetString(value, "pic_big")
			pic_mw2000, err := jsonparser.GetString(value, "pic_mw2000")
			photo_tag, err := jsonparser.GetInt(value, "photo_tag")
			pic_id, err := jsonparser.GetString(value, "pic_id")
			video, err := jsonparser.GetString(value, "video")
			pic_type, err := jsonparser.GetString(value, "pic_type")

			Id, err := jsonparser.GetString(value, "mblog", "id")
			Mid, err := jsonparser.GetString(value, "mblog", "mid")
			Text, err := jsonparser.GetString(value, "mblog", "text")
			IsLongText, err := jsonparser.GetBoolean(value, "mblog", "isLongText")
			AuthorId, err := jsonparser.GetString(value, "mblog", "user", "idstr")
			AuthorScreenName, err := jsonparser.GetString(value, "mblog", "user", "screen_name")
			AuthorProfileImageURL, err := jsonparser.GetString(value, "mblog", "user", "profile_image_url")
			photos[i] = &PhotoInfo{
				Pic_small:  pic_small,
				Pic_middle: pic_middle,
				Pic_big:    pic_big,
				Pic_mw2000: pic_mw2000,
				Photo_tag:  photo_tag,
				Pic_id:     pic_id,
				Video:      video,
				Pic_type:   pic_type,
				Blog: &MBlog{
					Id:                    Id,
					Mid:                   Mid,
					Text:                  Text,
					IsLongText:            IsLongText,
					AuthorId:              AuthorId,
					AuthorScreenName:      AuthorScreenName,
					AuthorProfileImageURL: AuthorProfileImageURL,
				},
			}
			i++

		}, "pics")
		if err != nil {
			log.Println(err)
		}
	}, "cards", albumIndex, "card_group")

	if err != nil {
		log.Println(err)
	}

	return photos[:i], nil
}
