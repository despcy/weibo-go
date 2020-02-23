package album

import (
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
)

func TestDownload(t *testing.T) {
	uid := "1669879400"
	client := NewAlbumClient(uid)
	os.Mkdir(uid, os.ModePerm)
	var wg sync.WaitGroup
	for client.HasNext {
		photos, err := client.RequestNextPage()

		if err != nil {
			println(err.Error())
		}

		for _, p := range photos {
			wg.Add(1)
			println(p.pic_small)
			go DownloadFile(&wg, uid, p.pic_small, p.pic_id+".jpg")

		}

	}
	wg.Wait()
}

func DownloadFile(wg *sync.WaitGroup, filepath string, url string, filename string) error {

	// Get the data
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {

		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath + "/" + filename)
	if err != nil {

		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	println("downloaded" + url)
	return err
}
