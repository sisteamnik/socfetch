package socfetch

import (
	"encoding/json"
	"fmt"
	"github.com/ungerik/go-dry"
	"net/url"
	"strconv"
	"time"
)

var DefaultVkApi = &VkApi{
	BaseUrl: "https://api.vk.com/method/",
	Version: "5.29",
}

type VkApi struct {
	Version string
	BaseUrl string
}

func (api *VkApi) Feed(id string) (rsp []Media) {
	method := "wall.get"
	v := url.Values{"owner_id": []string{id}}
	bts, err := api.req(method, v)
	if err != nil {
		return
	}
	vkp := VkPosts{}
	err = json.Unmarshal(bts, &vkp)
	if err != nil {
		panic(err)
	}

	for _, v := range vkp.Response.Items {
		rsp = append(rsp, Media(VkMedia{Post: v}))
	}

	return rsp
}

func (api *VkApi) req(method string, v url.Values) ([]byte, error) {
	u, err := url.Parse(api.BaseUrl + method)
	if err != nil {
		return []byte{}, err
	}
	v.Set("v", api.Version)
	u.RawQuery = v.Encode()
	bts, err := dry.FileGetBytes(u.String())
	return bts, err
}

type VkPosts struct {
	Response Response `json:"response"`
}

type Response struct {
	Count int      `json:"count"`
	Items []VkPost `json:"items"`
}

type VkPost struct {
	Id          int                 `json:"id"`
	Date        int                 `json:"date"`
	Owner_Id    int                 `json:"owner_id"`
	From_Id     int                 `json:"from_id"`
	Post_Type   string              `json:"post_type"`
	Text        string              `json:"text"`
	Attachments []VkPostAttachement `json:"attachments"`
}

type VkPostAttachement struct {
	Type  string `json:"type"`
	Photo VkPhoto
}

type VkPhoto struct {
	Pid      int    `json:"pid"`
	OwnerId  int    `json:"owner_id"`
	Src      string `json:"src"`
	Src_Big  string `json:"src_big"`
	Src_xBig string `json:"src_xbig"`
}

type VkMedia struct {
	Post VkPost
}

func (vk VkMedia) Type() string {
	return vkCheckType(vk.Post)
}

func (vk VkMedia) Text() string {
	return vk.Post.Text
}

func (vk VkMedia) Created() time.Time {
	i, err := strconv.ParseInt(fmt.Sprint(vk.Post.Date), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}

func vkCheckType(post VkPost) string {
	return "post"
}
