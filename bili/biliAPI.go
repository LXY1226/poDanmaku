package bili

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const BiliHost = "https://api.live.bilibili.com"

const UserAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Mobile Safari/537.36 Edg/92.0.902.67"

type (
	BiliResp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		//TTL     int    `json:"ttl"`
		Data json.RawMessage `json:"data"`
	}

	// DanmuInfo /xlive/web-room/v1/index/getDanmuInfo?RoomID=%roomId&type=0
	// where RoomID=RoomID
	DanmuInfo struct {
		//Group            string  `json:"group"`
		//BusinessID       int     `json:"business_id"`
		//RefreshRowFactor float64 `json:"refresh_row_factor"`
		//RefreshRate      int     `json:"refresh_rate"`
		//MaxDelay         int     `json:"max_delay"`
		Token    json.RawMessage `json:"token"`
		HostList []HostList      `json:"host_list"`
	}

	HostList struct {
		Host string `json:"host"`
		//Port    int    `json:"port"`
		//WssPort int    `json:"wss_port"`
		//WsPort  int    `json:"ws_port"`
	}

	UserInfo struct {
		UID   uint32 `json:"uid"`
		Uname string `json:"uname"`
		Face  string `json:"face"`
		//Rank              string `json:"rank"`
		//PlatformUserLevel int    `json:"platform_user_level"`
		//MobileVerify      int    `json:"mobile_verify"`
		//Identification    int    `json:"identification"`
		//OfficialVerify    struct {
		//	Type int    `json:"type"`
		//	Desc string `json:"desc"`
		//	Role int    `json:"role"`
		//} `json:"official_verify"`
		//VipType int `json:"vip_type"`
		//Gender  int `json:"gender"`
	}

	// AnchorInRoom /live_user/v1/UserInfo/get_anchor_in_room?roomid=%roomId
	AnchorInRoom struct {
		Info UserInfo `json:"info"`
		//Level struct {
		//	UID         int    `json:"uid"`
		//	Cost        int    `json:"cost"`
		//	Rcost       int    `json:"rcost"`
		//	UserScore   string `json:"user_score"`
		//	Vip         int    `json:"vip"`
		//	VipTime     string `json:"vip_time"`
		//	Svip        int    `json:"svip"`
		//	SvipTime    string `json:"svip_time"`
		//	UpdateTime  string `json:"update_time"`
		//	MasterLevel struct {
		//		Level            int    `json:"level"`
		//		Color            int    `json:"color"`
		//		Current          []int  `json:"current"`
		//		Next             []int  `json:"next"`
		//		AnchorScore      int    `json:"anchor_score"`
		//		UpgradeScore     int    `json:"upgrade_score"`
		//		MasterLevelColor int    `json:"master_level_color"`
		//		Sort             string `json:"sort"`
		//	} `json:"master_level"`
		//	UserLevel   int `json:"user_level"`
		//	Color       int `json:"color"`
		//	AnchorScore int `json:"anchor_score"`
		//} `json:"level"`
		//San int `json:"san"`
	}

	// RoomInfo /room/v1/Room/get_info?id=%roomId

	RoomInfo struct {
		//Uid       int `json:"uid"`
		RoomId    uint32 `json:"room_id"`
		ShortId   uint32 `json:"short_id"`
		Attention int    `json:"attention"`
		Online    int    `json:"online"`
		//IsPortrait       bool     `json:"is_portrait"`
		Description    string `json:"description"`
		LiveStatus     int    `json:"live_status"`
		AreaId         int    `json:"area_id"`
		ParentAreaId   int    `json:"parent_area_id"`
		ParentAreaName string `json:"parent_area_name"`
		//OldAreaId        int      `json:"old_area_id"`
		Background string `json:"background"`
		Title      string `json:"title"`
		UserCover  string `json:"user_cover"`
		//Keyframe         string   `json:"keyframe"`
		IsStrictRoom bool `json:"is_strict_room"`
		//LiveTime         string   `json:"live_time"`
		Tags     string `json:"tags"`
		IsAnchor int    `json:"is_anchor"`
		//RoomSilentType   string   `json:"room_silent_type"`
		//RoomSilentLevel  int      `json:"room_silent_level"`
		//RoomSilentSecond int      `json:"room_silent_second"`
		AreaName string `json:"area_name"`
		//Pendants         string   `json:"pendants"`
		//AreaPendants     string   `json:"area_pendants"`
		//HotWords         []string `json:"hot_words"`
		//HotWordsStatus   int      `json:"hot_words_status"`
		//Verify           string   `json:"verify"`
		//NewPendants      struct {
		//	Frame struct {
		//		Name       string `json:"name"`
		//		Value      string `json:"value"`
		//		Position   int    `json:"position"`
		//		Desc       string `json:"desc"`
		//		Area       int    `json:"area"`
		//		AreaOld    int    `json:"area_old"`
		//		BgColor    string `json:"bg_color"`
		//		BgPic      string `json:"bg_pic"`
		//		UseOldArea bool   `json:"use_old_area"`
		//	} `json:"frame"`
		//	Badge       interface{} `json:"badge"`
		//	MobileFrame struct {
		//		Name       string `json:"name"`
		//		Value      string `json:"value"`
		//		Position   int    `json:"position"`
		//		Desc       string `json:"desc"`
		//		Area       int    `json:"area"`
		//		AreaOld    int    `json:"area_old"`
		//		BgColor    string `json:"bg_color"`
		//		BgPic      string `json:"bg_pic"`
		//		UseOldArea bool   `json:"use_old_area"`
		//	} `json:"mobile_frame"`
		//	MobileBadge interface{} `json:"mobile_badge"`
		//} `json:"new_pendants"`
		//UpSession            string `json:"up_session"`
		//PkStatus             int    `json:"pk_status"`
		//PkId                 int    `json:"pk_id"`
		//BattleId             int    `json:"battle_id"`
		//AllowChangeAreaTime  int    `json:"allow_change_area_time"`
		//AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
		//StudioInfo           struct {
		//	Status     int           `json:"status"`
		//	MasterList []interface{} `json:"master_list"`
		//} `json:"studio_info"`
	}
)

func getAPI[T any](path, query string, v ...interface{}) *T {
	apiBucket.Use()
	req, _ := http.NewRequest("GET", BiliHost, nil)
	req.Header.Set("User-Agent", UserAgent)
	req.URL.Path = path
	req.URL.RawQuery = fmt.Sprintf(query, v...)
	var sleep time.Duration
	for {
		if sleep > 0 {
			time.Sleep(sleep)
		}
		sleep = 1 * time.Second
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("请求API出错：", err)
			continue
		}

		var biliResp BiliResp
		data, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(data, &biliResp)
		if err != nil {
			log.Println("解析API出错：", err)
			if resp.StatusCode == 429 {
				continue
			}
			return nil
		}
		if biliResp.Code != 0 {
			log.Println("API返回错误：", biliResp.Message, biliResp.Code)
			continue
		}
		var ret T
		err = json.Unmarshal(biliResp.Data, &ret)
		if err != nil {
			log.Println("API返回正常但解析出错", err)
			log.Println(path, query)
		}
		return &ret
	}
}

func GetDanmuInfo(RoomID uint32) *DanmuInfo {
	return getAPI[DanmuInfo]("/xlive/web-room/v1/index/getDanmuInfo", "id=%d&type=0", RoomID)
}

func GetUserInfo(RoomID uint32) *UserInfo {
	data := getAPI[AnchorInRoom]("/live_user/v1/UserInfo/get_anchor_in_room", "roomid=%d", RoomID)
	return &data.Info
}

func GetRoomInfo(RoomID uint32) *RoomInfo {
	return getAPI[RoomInfo]("/room/v1/Room/get_info", "id=%d", RoomID)
}
