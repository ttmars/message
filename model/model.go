package model

import (
	"net/http"
	"time"
)

type Item struct {
	Name        string
	Link        string
	Description string
	Badge       string
}

var M = make(map[string][]Item)
var DFClient = &http.Client{Timeout: time.Second * 10}
var Num = 20
var GithubNum = 10

type ZWTXStruct struct {
	Data struct {
		Total int `json:"total"`
		List  []struct {
			GUID      string `json:"guid"`
			ID        string `json:"id"`
			Time      string `json:"time"`
			Title     string `json:"title"`
			Length    string `json:"length"`
			Image     string `json:"image"`
			FocusDate int64  `json:"focus_date"`
			Brief     string `json:"brief"`
			URL       string `json:"url"`
			Mode      int    `json:"mode"`
		} `json:"list"`
	} `json:"data"`
}

type TiebaStruct struct {
	Data struct {
		BangHeadPic  string `json:"bang_head_pic"`
		UserHisTopic struct {
			ModuleTitle string        `json:"module_title"`
			TopicList   []interface{} `json:"topic_list"`
		} `json:"user_his_topic"`
		SugTopic struct {
			ModuleTitle string        `json:"module_title"`
			TopicList   []interface{} `json:"topic_list"`
		} `json:"sug_topic"`
		BangTopic struct {
			ModuleTitle string `json:"module_title"`
			TopicList   []struct {
				TopicID            int    `json:"topic_id"`
				TopicName          string `json:"topic_name"`
				TopicDesc          string `json:"topic_desc"`
				Abstract           string `json:"abstract"`
				TopicPic           string `json:"topic_pic"`
				Tag                int    `json:"tag"`
				DiscussNum         int    `json:"discuss_num"`
				IdxNum             int    `json:"idx_num"`
				CreateTime         int    `json:"create_time"`
				ContentNum         int    `json:"content_num"`
				TopicAvatar        string `json:"topic_avatar"`
				IsVideoTopic       string `json:"is_video_topic"`
				TopicURL           string `json:"topic_url"`
				TopicDefaultAvatar string `json:"topic_default_avatar"`
			} `json:"topic_list"`
		} `json:"bang_topic"`
		ManualTopic struct {
			ModuleTitle string        `json:"module_title"`
			TopicList   []interface{} `json:"topic_list"`
		} `json:"manual_topic"`
		Timestamp int64 `json:"timestamp"`
	} `json:"data"`
	Errno  int    `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type Douban1Struct struct {
	Subjects []struct {
		EpisodesInfo string `json:"episodes_info"`
		Rate         string `json:"rate"`
		CoverX       int    `json:"cover_x"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		Playable     bool   `json:"playable"`
		Cover        string `json:"cover"`
		ID           string `json:"id"`
		CoverY       int    `json:"cover_y"`
		IsNew        bool   `json:"is_new"`
	} `json:"subjects"`
}

type BiliStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Note string `json:"note"`
		List []struct {
			Aid       int    `json:"aid"`
			Videos    int    `json:"videos"`
			Tid       int    `json:"tid"`
			Tname     string `json:"tname"`
			Copyright int    `json:"copyright"`
			Pic       string `json:"pic"`
			Title     string `json:"title"`
			Pubdate   int    `json:"pubdate"`
			Ctime     int    `json:"ctime"`
			Desc      string `json:"desc"`
			State     int    `json:"state"`
			Duration  int    `json:"duration"`
			Rights    struct {
				Bp            int `json:"bp"`
				Elec          int `json:"elec"`
				Download      int `json:"download"`
				Movie         int `json:"movie"`
				Pay           int `json:"pay"`
				Hd5           int `json:"hd5"`
				NoReprint     int `json:"no_reprint"`
				Autoplay      int `json:"autoplay"`
				UgcPay        int `json:"ugc_pay"`
				IsCooperation int `json:"is_cooperation"`
				UgcPayPreview int `json:"ugc_pay_preview"`
				NoBackground  int `json:"no_background"`
				ArcPay        int `json:"arc_pay"`
				PayFreeWatch  int `json:"pay_free_watch"`
			} `json:"rights"`
			Owner struct {
				Mid  int    `json:"mid"`
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"owner"`
			Stat struct {
				Aid      int `json:"aid"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Favorite int `json:"favorite"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Dislike  int `json:"dislike"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			Dynamic   string `json:"dynamic"`
			Cid       int    `json:"cid"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			ShortLinkV2 string `json:"short_link_v2"`
			FirstFrame  string `json:"first_frame"`
			PubLocation string `json:"pub_location"`
			Bvid        string `json:"bvid"`
			Score       int    `json:"score"`
			EnableVt    int    `json:"enable_vt"`
			MissionID   int    `json:"mission_id,omitempty"`
			SeasonID    int    `json:"season_id,omitempty"`
			UpFromV2    int    `json:"up_from_v2,omitempty"`
			RedirectURL string `json:"redirect_url,omitempty"`
			Others      []struct {
				Aid       int    `json:"aid"`
				Videos    int    `json:"videos"`
				Tid       int    `json:"tid"`
				Tname     string `json:"tname"`
				Copyright int    `json:"copyright"`
				Pic       string `json:"pic"`
				Title     string `json:"title"`
				Pubdate   int    `json:"pubdate"`
				Ctime     int    `json:"ctime"`
				Desc      string `json:"desc"`
				State     int    `json:"state"`
				Attribute int    `json:"attribute"`
				Duration  int    `json:"duration"`
				MissionID int    `json:"mission_id"`
				Rights    struct {
					Bp            int `json:"bp"`
					Elec          int `json:"elec"`
					Download      int `json:"download"`
					Movie         int `json:"movie"`
					Pay           int `json:"pay"`
					Hd5           int `json:"hd5"`
					NoReprint     int `json:"no_reprint"`
					Autoplay      int `json:"autoplay"`
					UgcPay        int `json:"ugc_pay"`
					IsCooperation int `json:"is_cooperation"`
					UgcPayPreview int `json:"ugc_pay_preview"`
					NoBackground  int `json:"no_background"`
					ArcPay        int `json:"arc_pay"`
					PayFreeWatch  int `json:"pay_free_watch"`
				} `json:"rights"`
				Owner struct {
					Mid  int    `json:"mid"`
					Name string `json:"name"`
					Face string `json:"face"`
				} `json:"owner"`
				Stat struct {
					Aid      int `json:"aid"`
					View     int `json:"view"`
					Danmaku  int `json:"danmaku"`
					Reply    int `json:"reply"`
					Favorite int `json:"favorite"`
					Coin     int `json:"coin"`
					Share    int `json:"share"`
					NowRank  int `json:"now_rank"`
					HisRank  int `json:"his_rank"`
					Like     int `json:"like"`
					Dislike  int `json:"dislike"`
					Vt       int `json:"vt"`
					Vv       int `json:"vv"`
				} `json:"stat"`
				Dynamic   string `json:"dynamic"`
				Cid       int    `json:"cid"`
				Dimension struct {
					Width  int `json:"width"`
					Height int `json:"height"`
					Rotate int `json:"rotate"`
				} `json:"dimension"`
				SeasonID    int    `json:"season_id"`
				AttributeV2 int    `json:"attribute_v2"`
				ShortLinkV2 string `json:"short_link_v2"`
				FirstFrame  string `json:"first_frame"`
				PubLocation string `json:"pub_location"`
				Bvid        string `json:"bvid"`
				Score       int    `json:"score"`
				EnableVt    int    `json:"enable_vt"`
			} `json:"others,omitempty"`
		} `json:"list"`
	} `json:"data"`
}

type ToolStruct struct {
	Kind1 []struct {
		Name        string `json:"name"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Badge       string `json:"badge"`
	} `json:"kind1"`
	Kind2 []struct {
		Name        string `json:"name"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Badge       string `json:"badge"`
	} `json:"kind2"`
	Kind3 []struct {
		Name        string `json:"name"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Badge       string `json:"badge"`
	} `json:"kind3"`
	Kind4 []struct {
		Name        string `json:"name"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Badge       string `json:"badge"`
	} `json:"kind4"`
}

type CSDNStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Period            string      `json:"period"`
		HotRankScore      string      `json:"hotRankScore"`
		PcHotRankScore    string      `json:"pcHotRankScore"`
		LoginUserIsFollow bool        `json:"loginUserIsFollow"`
		NickName          string      `json:"nickName"`
		AvatarURL         string      `json:"avatarUrl"`
		UserName          string      `json:"userName"`
		ArticleTitle      string      `json:"articleTitle"`
		ArticleDetailURL  string      `json:"articleDetailUrl"`
		CommentCount      string      `json:"commentCount"`
		FavorCount        string      `json:"favorCount"`
		ViewCount         string      `json:"viewCount"`
		HotComment        interface{} `json:"hotComment"`
		PicList           []string    `json:"picList"`
		IsNew             interface{} `json:"isNew"`
		ProductID         string      `json:"productId"`
		ProductType       string      `json:"productType"`
		RecommendType     string      `json:"recommendType"`
		ReportData        interface{} `json:"report_data"`
	} `json:"data"`
}

type AcfunStruct struct {
	Result   int    `json:"result"`
	HostName string `json:"host-name"`
	RankList []struct {
		GroupID           string `json:"groupId"`
		UserID            int    `json:"userId"`
		DougaID           string `json:"dougaId"`
		UserImg           string `json:"userImg"`
		ChannelName       string `json:"channelName"`
		ContentID         int    `json:"contentId"`
		VideoCover        string `json:"videoCover"`
		FansCount         int    `json:"fansCount"`
		ChannelID         int    `json:"channelId"`
		ContributeTime    int64  `json:"contributeTime"`
		ContentTitle      string `json:"contentTitle"`
		ContentDesc       string `json:"contentDesc,omitempty"`
		UserSignature     string `json:"userSignature"`
		ContributionCount int    `json:"contributionCount"`
		DanmuCount        int    `json:"danmuCount"`
		ContentType       int    `json:"contentType"`
		IsFollowing       bool   `json:"isFollowing"`
		UserName          string `json:"userName"`
		Duration          int    `json:"duration"`
		ViewCount         int    `json:"viewCount"`
		StowCount         int    `json:"stowCount"`
		BananaCount       int    `json:"bananaCount"`
		CreateTimeMillis  int64  `json:"createTimeMillis"`
		DanmakuCount      int    `json:"danmakuCount"`
		VideoList         []struct {
			Priority             int    `json:"priority"`
			DanmakuCount         int    `json:"danmakuCount"`
			VisibleType          int    `json:"visibleType"`
			Title                string `json:"title"`
			SourceStatus         int    `json:"sourceStatus"`
			SizeType             int    `json:"sizeType"`
			DanmakuGuidePosition int    `json:"danmakuGuidePosition"`
			DanmakuCountShow     string `json:"danmakuCountShow"`
			UploadTime           int64  `json:"uploadTime"`
			DurationMillis       int    `json:"durationMillis"`
			FileName             string `json:"fileName"`
			ID                   string `json:"id"`
		} `json:"videoList"`
		RecoReason struct {
			Desc       string      `json:"desc"`
			Href       string      `json:"href"`
			Tag        string      `json:"tag"`
			LayoutType int         `json:"layoutType"`
			DescType   interface{} `json:"descType"`
			Type       int         `json:"type"`
		} `json:"recoReason"`
		CommentCountRealValue int `json:"commentCountRealValue"`
		CommentCount          int `json:"commentCount"`
		User                  struct {
			Action         int    `json:"action"`
			Href           string `json:"href"`
			HeadURL        string `json:"headUrl"`
			FanCountValue  int    `json:"fanCountValue"`
			FollowingCount string `json:"followingCount"`
			HeadCdnUrls    []struct {
				URL            string `json:"url"`
				FreeTrafficCdn bool   `json:"freeTrafficCdn"`
			} `json:"headCdnUrls"`
			IsJoinUpCollege      bool   `json:"isJoinUpCollege"`
			FollowingCountValue  int    `json:"followingCountValue"`
			ContributeCountValue int    `json:"contributeCountValue"`
			FanCount             string `json:"fanCount"`
			SocialMedal          struct {
			} `json:"socialMedal"`
			Gender          int    `json:"gender"`
			VerifiedTypes   []int  `json:"verifiedTypes"`
			NameColor       int    `json:"nameColor"`
			IsFollowing     bool   `json:"isFollowing"`
			AvatarImage     string `json:"avatarImage"`
			UserHeadImgInfo struct {
				Width          int  `json:"width"`
				Height         int  `json:"height"`
				Size           int  `json:"size"`
				Type           int  `json:"type"`
				Animated       bool `json:"animated"`
				ThumbnailImage struct {
					CdnUrls []struct {
						URL            string `json:"url"`
						FreeTrafficCdn bool   `json:"freeTrafficCdn"`
					} `json:"cdnUrls"`
				} `json:"thumbnailImage"`
				ThumbnailImageCdnURL string `json:"thumbnailImageCdnUrl"`
			} `json:"userHeadImgInfo"`
			IsFollowed           bool   `json:"isFollowed"`
			Name                 string `json:"name"`
			Signature            string `json:"signature"`
			FollowingStatus      int    `json:"followingStatus"`
			ContributeCount      string `json:"contributeCount"`
			AvatarFramePcImg     string `json:"avatarFramePcImg"`
			AvatarFrameMobileImg string `json:"avatarFrameMobileImg"`
			AvatarFrame          int    `json:"avatarFrame"`
			SexTrend             int    `json:"sexTrend"`
			ID                   string `json:"id"`
		} `json:"user"`
		LikeCount       int    `json:"likeCount"`
		CreateTime      string `json:"createTime"`
		ShareURL        string `json:"shareUrl"`
		Title           string `json:"title"`
		CoverURL        string `json:"coverUrl"`
		OriginalDeclare int    `json:"originalDeclare,omitempty"`
		TagList         []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"tagList,omitempty"`
		HasHotComment               bool   `json:"hasHotComment"`
		CommentCountTenThousandShow string `json:"commentCountTenThousandShow"`
		IsLike                      bool   `json:"isLike"`
		CoverCdnUrls                []struct {
			URL            string `json:"url"`
			FreeTrafficCdn bool   `json:"freeTrafficCdn"`
		} `json:"coverCdnUrls"`
		CoverImgInfo struct {
			Width          int  `json:"width"`
			Height         int  `json:"height"`
			Size           int  `json:"size"`
			Type           int  `json:"type"`
			Animated       bool `json:"animated"`
			ThumbnailImage struct {
				CdnUrls []struct {
					URL            string `json:"url"`
					FreeTrafficCdn bool   `json:"freeTrafficCdn"`
				} `json:"cdnUrls"`
			} `json:"thumbnailImage"`
			ThumbnailImageCdnURL string `json:"thumbnailImageCdnUrl"`
		} `json:"coverImgInfo"`
		PicShareURL            string `json:"picShareUrl"`
		IsDislike              bool   `json:"isDislike"`
		ViewCountShow          string `json:"viewCountShow"`
		ShareCountShow         string `json:"shareCountShow"`
		StowCountShow          string `json:"stowCountShow"`
		DanmakuCountShow       string `json:"danmakuCountShow"`
		BananaCountShow        string `json:"bananaCountShow"`
		GiftPeachCountShow     string `json:"giftPeachCountShow"`
		DisableEdit            bool   `json:"disableEdit"`
		BelongToSpecifyArubamu bool   `json:"belongToSpecifyArubamu"`
		CommentCountShow       string `json:"commentCountShow"`
		LikeCountShow          string `json:"likeCountShow"`
		GiftPeachCount         int    `json:"giftPeachCount"`
		DurationMillis         int    `json:"durationMillis"`
		Status                 int    `json:"status"`
		Description            string `json:"description"`
		Channel                struct {
			ParentID   int    `json:"parentId"`
			ParentName string `json:"parentName"`
			Name       string `json:"name"`
			ID         int    `json:"id"`
		} `json:"channel"`
		ShareCount         int  `json:"shareCount"`
		IsFavorite         bool `json:"isFavorite"`
		SuperUbb           bool `json:"superUbb"`
		IsRewardSupportted bool `json:"isRewardSupportted"`
		IsThrowBanana      bool `json:"isThrowBanana"`
		StaffContribute    bool `json:"staffContribute,omitempty"`
	} `json:"rankList"`
	RequestID string `json:"requestId"`
}

type JuejinStruct struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
	Data   []struct {
		ItemType int `json:"item_type"`
		ItemInfo struct {
			ArticleID   string `json:"article_id"`
			ArticleInfo struct {
				ArticleID      string  `json:"article_id"`
				UserID         string  `json:"user_id"`
				CategoryID     string  `json:"category_id"`
				TagIds         []int64 `json:"tag_ids"`
				VisibleLevel   int     `json:"visible_level"`
				LinkURL        string  `json:"link_url"`
				CoverImage     string  `json:"cover_image"`
				IsGfw          int     `json:"is_gfw"`
				Title          string  `json:"title"`
				BriefContent   string  `json:"brief_content"`
				IsEnglish      int     `json:"is_english"`
				IsOriginal     int     `json:"is_original"`
				UserIndex      float64 `json:"user_index"`
				OriginalType   int     `json:"original_type"`
				OriginalAuthor string  `json:"original_author"`
				Content        string  `json:"content"`
				Ctime          string  `json:"ctime"`
				Mtime          string  `json:"mtime"`
				Rtime          string  `json:"rtime"`
				DraftID        string  `json:"draft_id"`
				ViewCount      int     `json:"view_count"`
				CollectCount   int     `json:"collect_count"`
				DiggCount      int     `json:"digg_count"`
				CommentCount   int     `json:"comment_count"`
				HotIndex       int     `json:"hot_index"`
				IsHot          int     `json:"is_hot"`
				RankIndex      float64 `json:"rank_index"`
				Status         int     `json:"status"`
				VerifyStatus   int     `json:"verify_status"`
				AuditStatus    int     `json:"audit_status"`
				MarkContent    string  `json:"mark_content"`
				DisplayCount   int     `json:"display_count"`
			} `json:"article_info"`
			AuthorUserInfo struct {
				UserID            string `json:"user_id"`
				UserName          string `json:"user_name"`
				Company           string `json:"company"`
				JobTitle          string `json:"job_title"`
				AvatarLarge       string `json:"avatar_large"`
				Level             int    `json:"level"`
				Description       string `json:"description"`
				FolloweeCount     int    `json:"followee_count"`
				FollowerCount     int    `json:"follower_count"`
				PostArticleCount  int    `json:"post_article_count"`
				DiggArticleCount  int    `json:"digg_article_count"`
				GotDiggCount      int    `json:"got_digg_count"`
				GotViewCount      int    `json:"got_view_count"`
				PostShortmsgCount int    `json:"post_shortmsg_count"`
				DiggShortmsgCount int    `json:"digg_shortmsg_count"`
				Isfollowed        bool   `json:"isfollowed"`
				FavorableAuthor   int    `json:"favorable_author"`
				Power             int    `json:"power"`
				StudyPoint        int    `json:"study_point"`
				University        struct {
					UniversityID string `json:"university_id"`
					Name         string `json:"name"`
					Logo         string `json:"logo"`
				} `json:"university"`
				Major struct {
					MajorID  string `json:"major_id"`
					ParentID string `json:"parent_id"`
					Name     string `json:"name"`
				} `json:"major"`
				StudentStatus           int  `json:"student_status"`
				SelectEventCount        int  `json:"select_event_count"`
				SelectOnlineCourseCount int  `json:"select_online_course_count"`
				Identity                int  `json:"identity"`
				IsSelectAnnual          bool `json:"is_select_annual"`
				SelectAnnualRank        int  `json:"select_annual_rank"`
				AnnualListType          int  `json:"annual_list_type"`
				ExtraMap                struct {
				} `json:"extraMap"`
				IsLogout   int           `json:"is_logout"`
				AnnualInfo []interface{} `json:"annual_info"`
			} `json:"author_user_info"`
			Category struct {
				CategoryID      string `json:"category_id"`
				CategoryName    string `json:"category_name"`
				CategoryURL     string `json:"category_url"`
				Rank            int    `json:"rank"`
				BackGround      string `json:"back_ground"`
				Icon            string `json:"icon"`
				Ctime           int    `json:"ctime"`
				Mtime           int    `json:"mtime"`
				ShowType        int    `json:"show_type"`
				ItemType        int    `json:"item_type"`
				PromoteTagCap   int    `json:"promote_tag_cap"`
				PromotePriority int    `json:"promote_priority"`
			} `json:"category"`
			Tags []struct {
				ID               int    `json:"id"`
				TagID            string `json:"tag_id"`
				TagName          string `json:"tag_name"`
				Color            string `json:"color"`
				Icon             string `json:"icon"`
				BackGround       string `json:"back_ground"`
				ShowNavi         int    `json:"show_navi"`
				Ctime            int    `json:"ctime"`
				Mtime            int    `json:"mtime"`
				IDType           int    `json:"id_type"`
				TagAlias         string `json:"tag_alias"`
				PostArticleCount int    `json:"post_article_count"`
				ConcernUserCount int    `json:"concern_user_count"`
			} `json:"tags"`
			UserInteract struct {
				ID        int64 `json:"id"`
				Omitempty int   `json:"omitempty"`
				UserID    int   `json:"user_id"`
				IsDigg    bool  `json:"is_digg"`
				IsFollow  bool  `json:"is_follow"`
				IsCollect bool  `json:"is_collect"`
			} `json:"user_interact"`
			Org struct {
				OrgInfo    interface{} `json:"org_info"`
				OrgUser    interface{} `json:"org_user"`
				IsFollowed bool        `json:"is_followed"`
			} `json:"org"`
			ReqID  string `json:"req_id"`
			Status struct {
				PushStatus int `json:"push_status"`
			} `json:"status"`
			AuthorInteract interface{} `json:"author_interact"`
		} `json:"item_info"`
	} `json:"data"`
	Cursor  string `json:"cursor"`
	Count   int    `json:"count"`
	HasMore bool   `json:"has_more"`
}

type ZhihuStruct struct {
	InitialState struct {
		Common struct {
			Ask struct {
			} `json:"ask"`
		} `json:"common"`
		Loading struct {
			Global struct {
				Count int `json:"count"`
			} `json:"global"`
			Local struct {
				TopstoryGetHotList bool `json:"topstory/getHotList/"`
			} `json:"local"`
		} `json:"loading"`
		Club struct {
			Tags struct {
			} `json:"tags"`
			Admins struct {
				Data []interface{} `json:"data"`
			} `json:"admins"`
			Members struct {
				Data []interface{} `json:"data"`
			} `json:"members"`
			Explore struct {
				CandidateSyncClubs struct {
				} `json:"candidateSyncClubs"`
			} `json:"explore"`
			Profile struct {
			} `json:"profile"`
			Checkin struct {
			} `json:"checkin"`
			Comments struct {
				Paging struct {
				} `json:"paging"`
				Loading struct {
				} `json:"loading"`
				Meta struct {
				} `json:"meta"`
				Ids struct {
				} `json:"ids"`
			} `json:"comments"`
			PostList struct {
				Paging struct {
				} `json:"paging"`
				Loading struct {
				} `json:"loading"`
				Ids struct {
				} `json:"ids"`
			} `json:"postList"`
			Recommend struct {
				Data []interface{} `json:"data"`
			} `json:"recommend"`
			Silences struct {
				Data []interface{} `json:"data"`
			} `json:"silences"`
			Application struct {
				Profile interface{} `json:"profile"`
			} `json:"application"`
		} `json:"club"`
		Entities struct {
			Users struct {
			} `json:"users"`
			Questions struct {
			} `json:"questions"`
			Answers struct {
			} `json:"answers"`
			Articles struct {
			} `json:"articles"`
			Columns struct {
			} `json:"columns"`
			Topics struct {
			} `json:"topics"`
			Roundtables struct {
			} `json:"roundtables"`
			Favlists struct {
			} `json:"favlists"`
			Comments struct {
			} `json:"comments"`
			Notifications struct {
			} `json:"notifications"`
			Ebooks struct {
			} `json:"ebooks"`
			Activities struct {
			} `json:"activities"`
			Feeds struct {
			} `json:"feeds"`
			Pins struct {
			} `json:"pins"`
			Promotions struct {
			} `json:"promotions"`
			Drafts struct {
			} `json:"drafts"`
			Chats struct {
			} `json:"chats"`
			Posts struct {
			} `json:"posts"`
			Clubs struct {
			} `json:"clubs"`
			ClubTags struct {
			} `json:"clubTags"`
			Zvideos struct {
			} `json:"zvideos"`
			ZvideoContributions struct {
			} `json:"zvideoContributions"`
			Briefs struct {
			} `json:"briefs"`
		} `json:"entities"`
		CurrentUser string `json:"currentUser"`
		Account     struct {
			LockLevel struct {
			} `json:"lockLevel"`
			UnlockTicketStatus bool          `json:"unlockTicketStatus"`
			UnlockTicket       interface{}   `json:"unlockTicket"`
			Challenge          []interface{} `json:"challenge"`
			ErrorStatus        bool          `json:"errorStatus"`
			Message            string        `json:"message"`
			IsFetching         bool          `json:"isFetching"`
			AccountInfo        struct {
			} `json:"accountInfo"`
			URLToken struct {
				Loading bool `json:"loading"`
			} `json:"urlToken"`
			CardUserInfo struct {
				VipInfo struct {
				} `json:"vipInfo"`
			} `json:"cardUserInfo"`
			HandleWidget struct {
			} `json:"handleWidget"`
			WidgetList   []interface{} `json:"widgetList"`
			UserWidgetID string        `json:"userWidgetId"`
		} `json:"account"`
		Settings struct {
			SocialBind   interface{} `json:"socialBind"`
			InboxMsg     interface{} `json:"inboxMsg"`
			Notification struct {
			} `json:"notification"`
			Email struct {
			} `json:"email"`
			PrivacyFlag  interface{} `json:"privacyFlag"`
			BlockedUsers struct {
				IsFetching bool `json:"isFetching"`
				Paging     struct {
					PageNo   int `json:"pageNo"`
					PageSize int `json:"pageSize"`
				} `json:"paging"`
				Data []interface{} `json:"data"`
			} `json:"blockedUsers"`
			BlockedFollowees struct {
				IsFetching bool `json:"isFetching"`
				Paging     struct {
					PageNo   int `json:"pageNo"`
					PageSize int `json:"pageSize"`
				} `json:"paging"`
				Data []interface{} `json:"data"`
			} `json:"blockedFollowees"`
			IgnoredTopics struct {
				IsFetching bool `json:"isFetching"`
				Paging     struct {
					PageNo   int `json:"pageNo"`
					PageSize int `json:"pageSize"`
				} `json:"paging"`
				Data []interface{} `json:"data"`
			} `json:"ignoredTopics"`
			RestrictedTopics interface{} `json:"restrictedTopics"`
			Laboratory       struct {
			} `json:"laboratory"`
		} `json:"settings"`
		Notification struct {
		} `json:"notification"`
		People struct {
			ProfileStatus struct {
			} `json:"profileStatus"`
			ActivitiesByUser struct {
			} `json:"activitiesByUser"`
			AnswersByUser struct {
			} `json:"answersByUser"`
			AnswersSortByVotesByUser struct {
			} `json:"answersSortByVotesByUser"`
			AnswersIncludedByUser struct {
			} `json:"answersIncludedByUser"`
			VotedAnswersByUser struct {
			} `json:"votedAnswersByUser"`
			ThankedAnswersByUser struct {
			} `json:"thankedAnswersByUser"`
			VoteAnswersByUser struct {
			} `json:"voteAnswersByUser"`
			ThankAnswersByUser struct {
			} `json:"thankAnswersByUser"`
			TopicAnswersByUser struct {
			} `json:"topicAnswersByUser"`
			ZvideosByUser struct {
			} `json:"zvideosByUser"`
			ArticlesByUser struct {
			} `json:"articlesByUser"`
			ArticlesSortByVotesByUser struct {
			} `json:"articlesSortByVotesByUser"`
			ArticlesIncludedByUser struct {
			} `json:"articlesIncludedByUser"`
			PinsByUser struct {
			} `json:"pinsByUser"`
			QuestionsByUser struct {
			} `json:"questionsByUser"`
			CommercialQuestionsByUser struct {
			} `json:"commercialQuestionsByUser"`
			FavlistsByUser struct {
			} `json:"favlistsByUser"`
			FollowingByUser struct {
			} `json:"followingByUser"`
			FollowersByUser struct {
			} `json:"followersByUser"`
			MutualsByUser struct {
			} `json:"mutualsByUser"`
			FollowingColumnsByUser struct {
			} `json:"followingColumnsByUser"`
			FollowingQuestionsByUser struct {
			} `json:"followingQuestionsByUser"`
			FollowingFavlistsByUser struct {
			} `json:"followingFavlistsByUser"`
			FollowingTopicsByUser struct {
			} `json:"followingTopicsByUser"`
			PublicationsByUser struct {
			} `json:"publicationsByUser"`
			ColumnsByUser struct {
			} `json:"columnsByUser"`
			AllFavlistsByUser struct {
			} `json:"allFavlistsByUser"`
			Brands          interface{} `json:"brands"`
			CreationsByUser struct {
			} `json:"creationsByUser"`
			CreationsSortByVotesByUser struct {
			} `json:"creationsSortByVotesByUser"`
			CreationsFeed struct {
			} `json:"creationsFeed"`
			Infinity struct {
			} `json:"infinity"`
			BatchUsers struct {
			} `json:"batchUsers"`
			ProfileInfinity interface{} `json:"profileInfinity"`
		} `json:"people"`
		Env struct {
			Ab struct {
				Config struct {
					Experiments []struct {
						ExpID                string `json:"expId"`
						ExpPrefix            string `json:"expPrefix"`
						IsDynamicallyUpdated bool   `json:"isDynamicallyUpdated"`
						IsRuntime            bool   `json:"isRuntime"`
						IncludeTriggerInfo   bool   `json:"includeTriggerInfo"`
					} `json:"experiments"`
					Params []struct {
						ID      string `json:"id"`
						Type    string `json:"type"`
						Value   string `json:"value"`
						ChainID string `json:"chainId,omitempty"`
						LayerID string `json:"layerId"`
						Key     int    `json:"key,omitempty"`
					} `json:"params"`
					Chains []struct {
						ChainID string `json:"chainId"`
					} `json:"chains"`
					EncodedParams string `json:"encodedParams"`
				} `json:"config"`
				Triggers struct {
				} `json:"triggers"`
			} `json:"ab"`
			UserAgent struct {
				Edge             bool   `json:"Edge"`
				IE               bool   `json:"IE"`
				Wechat           bool   `json:"Wechat"`
				Weibo            bool   `json:"Weibo"`
				QQ               bool   `json:"QQ"`
				MQQBrowser       bool   `json:"MQQBrowser"`
				Qzone            bool   `json:"Qzone"`
				Mobile           bool   `json:"Mobile"`
				Android          bool   `json:"Android"`
				IOS              bool   `json:"iOS"`
				IsAppleDevice    bool   `json:"isAppleDevice"`
				Zhihu            bool   `json:"Zhihu"`
				ZhihuHybrid      bool   `json:"ZhihuHybrid"`
				IsBot            bool   `json:"isBot"`
				Tablet           bool   `json:"Tablet"`
				UC               bool   `json:"UC"`
				Quark            bool   `json:"Quark"`
				Sogou            bool   `json:"Sogou"`
				Qihoo            bool   `json:"Qihoo"`
				Baidu            bool   `json:"Baidu"`
				BaiduApp         bool   `json:"BaiduApp"`
				Safari           bool   `json:"Safari"`
				GoogleBot        bool   `json:"GoogleBot"`
				AndroidDaily     bool   `json:"AndroidDaily"`
				IOSDaily         bool   `json:"iOSDaily"`
				WxMiniProgram    bool   `json:"WxMiniProgram"`
				BaiduMiniProgram bool   `json:"BaiduMiniProgram"`
				QQMiniProgram    bool   `json:"QQMiniProgram"`
				JDMiniProgram    bool   `json:"JDMiniProgram"`
				IsWebView        bool   `json:"isWebView"`
				IsMiniProgram    bool   `json:"isMiniProgram"`
				Origin           string `json:"origin"`
			} `json:"userAgent"`
			AppViewConfig struct {
			} `json:"appViewConfig"`
			Ctx struct {
				Path  string `json:"path"`
				Query struct {
				} `json:"query"`
				Href string `json:"href"`
				Host string `json:"host"`
			} `json:"ctx"`
			TrafficSource string `json:"trafficSource"`
			Edition       struct {
				Beijing      bool `json:"beijing"`
				Baidu        bool `json:"baidu"`
				Sogou        bool `json:"sogou"`
				BaiduBeijing bool `json:"baiduBeijing"`
				SogouBeijing bool `json:"sogouBeijing"`
				SogouInput   bool `json:"sogouInput"`
				BaiduSearch  bool `json:"baiduSearch"`
				GoogleSearch bool `json:"googleSearch"`
				Shenma       bool `json:"shenma"`
				MiniProgram  bool `json:"miniProgram"`
				Xiaomi       bool `json:"xiaomi"`
			} `json:"edition"`
			Theme          string `json:"theme"`
			AppHeaderTheme struct {
				Current string `json:"current"`
				Disable bool   `json:"disable"`
				Normal  struct {
					BgColor string `json:"bgColor"`
				} `json:"normal"`
				Custom struct {
					BgColor string `json:"bgColor"`
				} `json:"custom"`
			} `json:"appHeaderTheme"`
			EnableShortcut bool   `json:"enableShortcut"`
			Referer        string `json:"referer"`
			XUDID          string `json:"xUDId"`
			Mode           string `json:"mode"`
			Conf           struct {
			} `json:"conf"`
			XTrafficFreeOrigin string `json:"xTrafficFreeOrigin"`
			IPInfo             struct {
			} `json:"ipInfo"`
			Logged bool `json:"logged"`
			Vars   struct {
				PassThroughHeaders struct {
				} `json:"passThroughHeaders"`
			} `json:"vars"`
		} `json:"env"`
		Me struct {
			ColumnContributions []interface{} `json:"columnContributions"`
		} `json:"me"`
		Label struct {
			RecognizerLists struct {
			} `json:"recognizerLists"`
		} `json:"label"`
		Ecommerce struct {
		} `json:"ecommerce"`
		Favlists struct {
			Relations struct {
			} `json:"relations"`
		} `json:"favlists"`
		Question struct {
			Followers struct {
			} `json:"followers"`
			ConcernedFollowers struct {
			} `json:"concernedFollowers"`
			Answers struct {
			} `json:"answers"`
			HiddenAnswers struct {
			} `json:"hiddenAnswers"`
			UpdatedAnswers struct {
			} `json:"updatedAnswers"`
			AriaAnswers struct {
			} `json:"ariaAnswers"`
			CollapsedAnswers struct {
			} `json:"collapsedAnswers"`
			NotificationAnswers struct {
			} `json:"notificationAnswers"`
			InvitedQuestions struct {
				Total struct {
					Count     interface{}   `json:"count"`
					IsEnd     bool          `json:"isEnd"`
					IsLoading bool          `json:"isLoading"`
					Questions []interface{} `json:"questions"`
				} `json:"total"`
				Followees struct {
					Count     interface{}   `json:"count"`
					IsEnd     bool          `json:"isEnd"`
					IsLoading bool          `json:"isLoading"`
					Questions []interface{} `json:"questions"`
				} `json:"followees"`
			} `json:"invitedQuestions"`
			LaterQuestions struct {
				Count     interface{}   `json:"count"`
				IsEnd     bool          `json:"isEnd"`
				IsLoading bool          `json:"isLoading"`
				Questions []interface{} `json:"questions"`
			} `json:"laterQuestions"`
			WaitingQuestions struct {
				Recommend struct {
					IsEnd     bool          `json:"isEnd"`
					IsLoading bool          `json:"isLoading"`
					Questions []interface{} `json:"questions"`
				} `json:"recommend"`
				Invite struct {
					IsEnd     bool          `json:"isEnd"`
					IsLoading bool          `json:"isLoading"`
					Questions []interface{} `json:"questions"`
				} `json:"invite"`
				Newest struct {
					IsEnd     bool          `json:"isEnd"`
					IsLoading bool          `json:"isLoading"`
					Questions []interface{} `json:"questions"`
				} `json:"newest"`
				Hot struct {
					IsEnd     bool          `json:"isEnd"`
					IsLoading bool          `json:"isLoading"`
					Questions []interface{} `json:"questions"`
				} `json:"hot"`
			} `json:"waitingQuestions"`
			InvitationCandidates struct {
			} `json:"invitationCandidates"`
			Inviters struct {
			} `json:"inviters"`
			Invitees struct {
			} `json:"invitees"`
			SimilarQuestions struct {
			} `json:"similarQuestions"`
			QuestionBanners struct {
			} `json:"questionBanners"`
			RelatedCommodities struct {
			} `json:"relatedCommodities"`
			Bio struct {
			} `json:"bio"`
			Brand struct {
			} `json:"brand"`
			Permission struct {
			} `json:"permission"`
			Adverts struct {
			} `json:"adverts"`
			AdvancedStyle struct {
			} `json:"advancedStyle"`
			CommonAnswerCount int `json:"commonAnswerCount"`
			HiddenAnswerCount int `json:"hiddenAnswerCount"`
			TopicMeta         struct {
			} `json:"topicMeta"`
			BluestarRanklist struct {
			} `json:"bluestarRanklist"`
			RelatedSearch struct {
			} `json:"relatedSearch"`
			AutoInvitation struct {
			} `json:"autoInvitation"`
			SimpleConcernedFollowers struct {
			} `json:"simpleConcernedFollowers"`
			DraftStatus struct {
			} `json:"draftStatus"`
			Disclaimers struct {
			} `json:"disclaimers"`
			IsShowMobileSignInModal bool `json:"isShowMobileSignInModal"`
			ModalShowCondition      struct {
			} `json:"modalShowCondition"`
		} `json:"question"`
		Banner struct {
		} `json:"banner"`
		Comments struct {
			Pagination struct {
			} `json:"pagination"`
			Collapsed struct {
			} `json:"collapsed"`
			Reverse struct {
			} `json:"reverse"`
			Reviewing struct {
			} `json:"reviewing"`
			Conversation struct {
			} `json:"conversation"`
			Parent struct {
			} `json:"parent"`
		} `json:"comments"`
		CommentsV2 struct {
			Stickers                 []interface{} `json:"stickers"`
			CommentWithPicPermission struct {
			} `json:"commentWithPicPermission"`
			NotificationsComments struct {
			} `json:"notificationsComments"`
			Pagination struct {
			} `json:"pagination"`
			Collapsed struct {
			} `json:"collapsed"`
			Reverse struct {
			} `json:"reverse"`
			Reviewing struct {
			} `json:"reviewing"`
			Conversation struct {
			} `json:"conversation"`
			ConversationMore struct {
			} `json:"conversationMore"`
			Parent struct {
			} `json:"parent"`
		} `json:"commentsV2"`
		Answers struct {
			Voters struct {
			} `json:"voters"`
			CopyrightApplicants struct {
			} `json:"copyrightApplicants"`
			Favlists struct {
			} `json:"favlists"`
			NewAnswer struct {
			} `json:"newAnswer"`
			EntityWords struct {
			} `json:"entityWords"`
			ConcernedUpvoters struct {
			} `json:"concernedUpvoters"`
			SimpleConcernedUpvoters struct {
			} `json:"simpleConcernedUpvoters"`
			PaidContent struct {
			} `json:"paidContent"`
			Settings struct {
			} `json:"settings"`
		} `json:"answers"`
		Switches struct {
		} `json:"switches"`
		Video struct {
			Data struct {
			} `json:"data"`
			ShareVideoDetail struct {
			} `json:"shareVideoDetail"`
			Last struct {
			} `json:"last"`
		} `json:"video"`
		Sms struct {
			SupportedCountries []interface{} `json:"supportedCountries"`
		} `json:"sms"`
		Daily struct {
		} `json:"daily"`
		Explore struct {
			Recommendations struct {
			} `json:"recommendations"`
			Specials struct {
				Entities struct {
				} `json:"entities"`
				Order []interface{} `json:"order"`
			} `json:"specials"`
			Roundtables struct {
				Entities struct {
				} `json:"entities"`
				Order []interface{} `json:"order"`
			} `json:"roundtables"`
			Collections struct {
			} `json:"collections"`
			Columns struct {
			} `json:"columns"`
			Square struct {
				HotQuestionList []interface{} `json:"hotQuestionList"`
				PotentialList   []interface{} `json:"potentialList"`
			} `json:"square"`
		} `json:"explore"`
		RelatedReadings struct {
		} `json:"relatedReadings"`
		HotQuestions struct {
		} `json:"hotQuestions"`
		PushNotifications struct {
			Default struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				Ids        []interface{} `json:"ids"`
			} `json:"default"`
			Follow struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				Ids        []interface{} `json:"ids"`
			} `json:"follow"`
			VoteThank struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				Ids        []interface{} `json:"ids"`
			} `json:"vote_thank"`
			CurrentTab         string `json:"currentTab"`
			NotificationsCount struct {
				Default   int `json:"default"`
				Follow    int `json:"follow"`
				VoteThank int `json:"vote_thank"`
			} `json:"notificationsCount"`
		} `json:"pushNotifications"`
		Messages struct {
			Data struct {
			} `json:"data"`
			CurrentTab   string `json:"currentTab"`
			MessageCount int    `json:"messageCount"`
		} `json:"messages"`
		Pins struct {
			Reviewing struct {
			} `json:"reviewing"`
		} `json:"pins"`
		Pin struct {
			Pin      interface{}   `json:"pin"`
			Brief    interface{}   `json:"brief"`
			Author   interface{}   `json:"author"`
			Comments []interface{} `json:"comments"`
			Pins     struct {
			} `json:"pins"`
			Special struct {
			} `json:"special"`
			Daily struct {
			} `json:"daily"`
			Topic struct {
			} `json:"topic"`
			TopicFeeds struct {
				Data      []interface{} `json:"data"`
				IsLoading bool          `json:"isLoading"`
				Paging    struct {
				} `json:"paging"`
			} `json:"topicFeeds"`
		} `json:"pin"`
		Captcha struct {
			CaptchaNeeded            bool        `json:"captchaNeeded"`
			CaptchaValidated         bool        `json:"captchaValidated"`
			CaptchaBase64String      interface{} `json:"captchaBase64String"`
			CaptchaValidationMessage interface{} `json:"captchaValidationMessage"`
			LoginCaptchaExpires      bool        `json:"loginCaptchaExpires"`
		} `json:"captcha"`
		Login struct {
			LoginUnregisteredError    bool        `json:"loginUnregisteredError"`
			LoginBindWechatError      bool        `json:"loginBindWechatError"`
			LoginConfirmError         interface{} `json:"loginConfirmError"`
			SendDigitsError           interface{} `json:"sendDigitsError"`
			NeedSMSIdentify           bool        `json:"needSMSIdentify"`
			ValidateDigitsError       bool        `json:"validateDigitsError"`
			LoginConfirmSucceeded     interface{} `json:"loginConfirmSucceeded"`
			QrcodeLoginToken          string      `json:"qrcodeLoginToken"`
			QrcodeLoginScanStatus     int         `json:"qrcodeLoginScanStatus"`
			QrcodeLoginError          interface{} `json:"qrcodeLoginError"`
			QrcodeLoginReturnNewToken bool        `json:"qrcodeLoginReturnNewToken"`
		} `json:"login"`
		Register struct {
			RegisterValidateSucceeded interface{} `json:"registerValidateSucceeded"`
			RegisterValidateErrors    struct {
			} `json:"registerValidateErrors"`
			RegisterConfirmError     interface{} `json:"registerConfirmError"`
			SendDigitsError          interface{} `json:"sendDigitsError"`
			RegisterConfirmSucceeded interface{} `json:"registerConfirmSucceeded"`
		} `json:"register"`
		Topic struct {
			Bios struct {
			} `json:"bios"`
			Hot struct {
			} `json:"hot"`
			Newest struct {
			} `json:"newest"`
			Top struct {
			} `json:"top"`
			Unanswered struct {
			} `json:"unanswered"`
			Questions struct {
			} `json:"questions"`
			Followers struct {
			} `json:"followers"`
			Contributors struct {
			} `json:"contributors"`
			Parent struct {
			} `json:"parent"`
			Children struct {
			} `json:"children"`
			BestAnswerers struct {
			} `json:"bestAnswerers"`
			WikiMeta struct {
			} `json:"wikiMeta"`
			Index struct {
			} `json:"index"`
			Intro struct {
			} `json:"intro"`
			Meta struct {
			} `json:"meta"`
			Schema struct {
			} `json:"schema"`
			CreatorWall struct {
			} `json:"creatorWall"`
			WikiEditInfo struct {
			} `json:"wikiEditInfo"`
			CommittedWiki struct {
			} `json:"committedWiki"`
			LandingBasicData struct {
			} `json:"landingBasicData"`
			LandingExcellentItems   []interface{} `json:"landingExcellentItems"`
			LandingExcellentEditors []interface{} `json:"landingExcellentEditors"`
			LandingCatalog          []interface{} `json:"landingCatalog"`
			LandingEntries          struct {
			} `json:"landingEntries"`
		} `json:"topic"`
		Drama struct {
		} `json:"drama"`
		Topstory struct {
			Recommend struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				AfterID    int           `json:"afterId"`
				Items      []interface{} `json:"items"`
				Next       interface{}   `json:"next"`
			} `json:"recommend"`
			Follow struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				AfterID    int           `json:"afterId"`
				Items      []interface{} `json:"items"`
				Next       interface{}   `json:"next"`
			} `json:"follow"`
			Room struct {
				Meta struct {
				} `json:"meta"`
				IsFetching bool          `json:"isFetching"`
				AfterID    int           `json:"afterId"`
				Items      []interface{} `json:"items"`
				Next       interface{}   `json:"next"`
			} `json:"room"`
			FollowWonderful struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				AfterID    int           `json:"afterId"`
				Items      []interface{} `json:"items"`
				Next       interface{}   `json:"next"`
			} `json:"followWonderful"`
			Sidebar      interface{} `json:"sidebar"`
			Announcement struct {
			} `json:"announcement"`
			HotList []struct {
				Type         string `json:"type"`
				StyleType    string `json:"styleType"`
				ID           string `json:"id"`
				CardID       string `json:"cardId"`
				FeedSpecific struct {
					AnswerCount int `json:"answerCount"`
				} `json:"feedSpecific"`
				Target struct {
					TitleArea struct {
						Text string `json:"text"`
					} `json:"titleArea"`
					ExcerptArea struct {
						Text string `json:"text"`
					} `json:"excerptArea"`
					ImageArea struct {
						URL string `json:"url"`
					} `json:"imageArea"`
					MetricsArea struct {
						Text string `json:"text"`
					} `json:"metricsArea"`
					LabelArea struct {
						Type        string `json:"type"`
						Trend       int    `json:"trend"`
						NightColor  string `json:"nightColor"`
						NormalColor string `json:"normalColor"`
					} `json:"labelArea"`
					Link struct {
						URL string `json:"url"`
					} `json:"link"`
				} `json:"target"`
				AttachedInfo string `json:"attachedInfo"`
			} `json:"hotList"`
			GuestFeeds struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				AfterID    int           `json:"afterId"`
				Items      []interface{} `json:"items"`
				Next       interface{}   `json:"next"`
			} `json:"guestFeeds"`
			FollowExtra struct {
				IsNewUser   interface{}   `json:"isNewUser"`
				IsFetched   bool          `json:"isFetched"`
				FollowCount int           `json:"followCount"`
				Followers   []interface{} `json:"followers"`
			} `json:"followExtra"`
			HotDaily struct {
				Data   []interface{} `json:"data"`
				Paging struct {
				} `json:"paging"`
			} `json:"hotDaily"`
			HotHighlight struct {
				IsFetching bool          `json:"isFetching"`
				IsDrained  bool          `json:"isDrained"`
				Data       []interface{} `json:"data"`
				Paging     struct {
				} `json:"paging"`
			} `json:"hotHighlight"`
			Banner struct {
			} `json:"banner"`
			CommercialBanner struct {
				Show   bool `json:"show"`
				Banner struct {
				} `json:"banner"`
				TrackData struct {
				} `json:"trackData"`
			} `json:"commercialBanner"`
			Video struct {
				Items     []interface{} `json:"items"`
				Next      interface{}   `json:"next"`
				IsLoading bool          `json:"isLoading"`
				IsDrained bool          `json:"isDrained"`
			} `json:"video"`
		} `json:"topstory"`
		Search struct {
			RecommendSearch []interface{} `json:"recommendSearch"`
			TopSearch       struct {
			} `json:"topSearch"`
			SearchValue struct {
			} `json:"searchValue"`
			SuggestSearch struct {
			} `json:"suggestSearch"`
			AttachedInfo struct {
				GeneralByQuery struct {
				} `json:"generalByQuery"`
			} `json:"attachedInfo"`
			NextOffset struct {
				GeneralByQuery struct {
				} `json:"generalByQuery"`
			} `json:"nextOffset"`
			TopicReview struct {
			} `json:"topicReview"`
			Calendar struct {
			} `json:"calendar"`
			Scores interface{} `json:"scores"`
			Majors struct {
			} `json:"majors"`
			University struct {
			} `json:"university"`
			GeneralByQuery struct {
			} `json:"generalByQuery"`
			GeneralByQueryInADay struct {
			} `json:"generalByQueryInADay"`
			GeneralByQueryInAWeek struct {
			} `json:"generalByQueryInAWeek"`
			GeneralByQueryInThreeMonths struct {
			} `json:"generalByQueryInThreeMonths"`
			PeopleByQuery struct {
			} `json:"peopleByQuery"`
			ClubentityByQuery struct {
			} `json:"clubentityByQuery"`
			ClubPostByQuery struct {
			} `json:"clubPostByQuery"`
			TopicByQuery struct {
			} `json:"topicByQuery"`
			ZvideoByQuery struct {
			} `json:"zvideoByQuery"`
			ScholarByQuery struct {
			} `json:"scholarByQuery"`
			ColumnByQuery struct {
			} `json:"columnByQuery"`
			LiveByQuery struct {
			} `json:"liveByQuery"`
			AlbumByQuery struct {
			} `json:"albumByQuery"`
			EBookByQuery struct {
			} `json:"eBookByQuery"`
			KmGeneralByQuery struct {
			} `json:"kmGeneralByQuery"`
			KmCourseByQuery struct {
			} `json:"kmCourseByQuery"`
			CustomFilter struct {
				Keys []interface{} `json:"keys"`
				Tags []interface{} `json:"tags"`
			} `json:"customFilter"`
		} `json:"search"`
		Story struct {
		} `json:"story"`
		Spread struct {
			SpreadData struct {
			} `json:"spreadData"`
		} `json:"spread"`
		Gift struct {
			GotStatus    bool   `json:"gotStatus"`
			Participated bool   `json:"participated"`
			IsSuccess    bool   `json:"isSuccess"`
			PhoneNo      string `json:"phoneNo"`
		} `json:"gift"`
		InterestZone struct {
			InterestZone []interface{} `json:"interestZone"`
			FeedsList    struct {
			} `json:"feedsList"`
		} `json:"interestZone"`
		VideoTopic struct {
			TopData struct {
			} `json:"topData"`
			Data      []interface{} `json:"data"`
			IsLoading bool          `json:"isLoading"`
			IsDrained bool          `json:"isDrained"`
			Next      string        `json:"next"`
		} `json:"videoTopic"`
		ReadStatus struct {
		} `json:"readStatus"`
		News struct {
			Special struct {
				Topic struct {
				} `json:"topic"`
				BgItems       []interface{} `json:"bgItems"`
				InDepthModule struct {
				} `json:"inDepthModule"`
				PinItems   []interface{} `json:"pinItems"`
				ShowButton bool          `json:"showButton"`
			} `json:"special"`
		} `json:"news"`
		ThemeColors struct {
		} `json:"themeColors"`
		Creator struct {
			CurrentCreatorURLToken interface{} `json:"currentCreatorUrlToken"`
			HomeData               struct {
				RecommendQuestions []interface{} `json:"recommendQuestions"`
			} `json:"homeData"`
			Tools struct {
				Question struct {
					InvitationCount struct {
						QuestionFolloweeCount int `json:"questionFolloweeCount"`
						QuestionTotalCount    int `json:"questionTotalCount"`
					} `json:"invitationCount"`
					GoodatTopics []interface{} `json:"goodatTopics"`
				} `json:"question"`
				CustomPromotion struct {
					ItemLists struct {
					} `json:"itemLists"`
				} `json:"customPromotion"`
				Recommend struct {
					RecommendTimes struct {
					} `json:"recommendTimes"`
				} `json:"recommend"`
			} `json:"tools"`
			Explore struct {
				Academy struct {
					Tabs    []interface{} `json:"tabs"`
					Article struct {
					} `json:"article"`
				} `json:"academy"`
			} `json:"explore"`
			Rights       []interface{} `json:"rights"`
			NewRights    []interface{} `json:"newRights"`
			RightsStatus struct {
			} `json:"rightsStatus"`
			LevelUpperLimit int `json:"levelUpperLimit"`
			Account         struct {
				GrowthLevel struct {
				} `json:"growthLevel"`
			} `json:"account"`
			Mcn struct {
			} `json:"mcn"`
			ApplyStatus struct {
			} `json:"applyStatus"`
			VideoSupport struct {
			} `json:"videoSupport"`
			TextBenefit struct {
			} `json:"textBenefit"`
			McnManage struct {
			} `json:"mcnManage"`
			Tasks struct {
			} `json:"tasks"`
			NewTasks   []interface{} `json:"newTasks"`
			NewTaskDes []interface{} `json:"newTaskDes"`
			ScoreInfo  struct {
			} `json:"scoreInfo"`
			RecentlyCreated []interface{} `json:"recentlyCreated"`
			Analysis        struct {
				All struct {
				} `json:"all"`
				Answer struct {
				} `json:"answer"`
				Zvideo struct {
				} `json:"zvideo"`
				Article struct {
				} `json:"article"`
				Pin struct {
				} `json:"pin"`
				SingleContent struct {
				} `json:"singleContent"`
			} `json:"analysis"`
			Announcement struct {
			} `json:"announcement"`
			BannerList []interface{} `json:"bannerList"`
			School     struct {
				Tabs     []interface{} `json:"tabs"`
				Contents []interface{} `json:"contents"`
				Banner   interface{}   `json:"banner"`
				Entities struct {
				} `json:"entities"`
			} `json:"school"`
			CreatorsRecommendInfo struct {
			} `json:"creatorsRecommendInfo"`
			MenusShowControlByServer struct {
				BVipRecomend         bool `json:"bVipRecomend"`
				CreationRelationship bool `json:"creationRelationship"`
			} `json:"menusShowControlByServer"`
			Income struct {
				Aggregation struct {
				} `json:"aggregation"`
			} `json:"income"`
		} `json:"creator"`
		CreationRanking struct {
		} `json:"creationRanking"`
		CreatorSalt struct {
			RecommendQuestionList []interface{} `json:"recommendQuestionList"`
			BannerList            []interface{} `json:"bannerList"`
			ClaimBannerList       []interface{} `json:"claimBannerList"`
			Sites                 []interface{} `json:"sites"`
			Domains               struct {
			} `json:"domains"`
			HasRecored             bool          `json:"hasRecored"`
			HasClaim               bool          `json:"hasClaim"`
			HasContributedList     []interface{} `json:"hasContributedList"`
			NotContributedList     []interface{} `json:"notContributedList"`
			ContributesTotal       interface{}   `json:"contributesTotal"`
			PreviewPageTitle       string        `json:"previewPageTitle"`
			PreviewPageContent     string        `json:"previewPageContent"`
			RestContributionNumber string        `json:"restContributionNumber"`
		} `json:"creatorSalt"`
		Emoticons struct {
			EmoticonGroupList   []interface{} `json:"emoticonGroupList"`
			EmoticonGroupDetail struct {
			} `json:"emoticonGroupDetail"`
		} `json:"emoticons"`
		Authorize struct {
		} `json:"authorize"`
		Theater struct {
			Account struct {
				SaltBalance  int           `json:"saltBalance"`
				CoinBalance  int           `json:"coinBalance"`
				FishBalance  int           `json:"fishBalance"`
				GiftReserves []interface{} `json:"giftReserves"`
			} `json:"account"`
			Activity struct {
				Activities struct {
				} `json:"activities"`
			} `json:"activity"`
			SideFeeds struct {
			} `json:"sideFeeds"`
			Debates struct {
				Data            []interface{} `json:"data"`
				LastDebateApply struct {
				} `json:"lastDebateApply"`
			} `json:"debates"`
			Theaters struct {
			} `json:"theaters"`
			FinishInfos struct {
			} `json:"finishInfos"`
			Forecasts struct {
			} `json:"forecasts"`
			Messages struct {
			} `json:"messages"`
			ActorApply struct {
				IsAuditing bool `json:"isAuditing"`
			} `json:"actorApply"`
			DramaCategories struct {
				CategoriesData []interface{} `json:"categoriesData"`
			} `json:"dramaCategories"`
			DramaFeeds struct {
			} `json:"dramaFeeds"`
			DramaGifts struct {
				Data []interface{} `json:"data"`
				Tabs []interface{} `json:"tabs"`
			} `json:"dramaGifts"`
			GrowTask struct {
			} `json:"growTask"`
		} `json:"theater"`
		Ranking struct {
		} `json:"ranking"`
		KnowledgePlan struct {
			Lists struct {
			} `json:"lists"`
			AllCreationRankList struct {
			} `json:"allCreationRankList"`
			FeaturedQuestions struct {
			} `json:"featuredQuestions"`
		} `json:"knowledgePlan"`
		Topsearch struct {
		} `json:"topsearch"`
		Collections struct {
			Hot struct {
				Data   []interface{} `json:"data"`
				Paging struct {
				} `json:"paging"`
				IsLoading bool `json:"isLoading"`
			} `json:"hot"`
			CollectionFeeds struct {
			} `json:"collectionFeeds"`
		} `json:"collections"`
		SearchSpecial struct {
			Contents   []interface{} `json:"contents"`
			Navigators []interface{} `json:"navigators"`
		} `json:"searchSpecial"`
		Liveplus struct {
		} `json:"liveplus"`
		WallE struct {
			ProtectHistory struct {
				Total int `json:"total"`
				Pages struct {
				} `json:"pages"`
				Entities struct {
				} `json:"entities"`
			} `json:"protectHistory"`
		} `json:"wallE"`
		CommentManage struct {
			CommentList struct {
				Ids      []interface{} `json:"ids"`
				Entities struct {
				} `json:"entities"`
				NextOffset int    `json:"nextOffset"`
				URLToken   string `json:"urlToken"`
			} `json:"commentList"`
			SubCommentList struct {
				Ids      []interface{} `json:"ids"`
				Entities struct {
				} `json:"entities"`
				Paging struct {
					Next  string `json:"next"`
					IsEnd bool   `json:"isEnd"`
				} `json:"paging"`
			} `json:"subCommentList"`
		} `json:"commentManage"`
		ZhiPlus struct {
			PermissionStatus int `json:"permissionStatus"`
		} `json:"zhiPlus"`
		Zvideos struct {
			CampaignVideoList struct {
			} `json:"campaignVideoList"`
			Campaigns struct {
			} `json:"campaigns"`
			TagoreCategory  []interface{} `json:"tagoreCategory"`
			Recommendations struct {
			} `json:"recommendations"`
			Insertable struct {
			} `json:"insertable"`
			Recruit struct {
				Form struct {
					Platform      string `json:"platform"`
					Nickname      string `json:"nickname"`
					FollowerCount string `json:"followerCount"`
					Domain        string `json:"domain"`
					Contact       string `json:"contact"`
				} `json:"form"`
				Submited bool          `json:"submited"`
				Ranking  []interface{} `json:"ranking"`
			} `json:"recruit"`
			Club struct {
			} `json:"club"`
			QyActivityData struct {
			} `json:"qyActivityData"`
			TalkActivityData struct {
			} `json:"talkActivityData"`
			Party2022ActivityData struct {
			} `json:"party2022ActivityData"`
			BatchVideos struct {
			} `json:"batchVideos"`
			Contribution struct {
				SelectedContribution interface{} `json:"selectedContribution"`
				Campaign             interface{} `json:"campaign"`
				Configs              struct {
				} `json:"configs"`
				ContributionLists struct {
				} `json:"contributionLists"`
				RecommendQuestions struct {
					IsLoading bool `json:"isLoading"`
					Paging    struct {
						IsEnd   bool `json:"isEnd"`
						IsStart bool `json:"isStart"`
						Totals  int  `json:"totals"`
					} `json:"paging"`
					Data []interface{} `json:"data"`
				} `json:"recommendQuestions"`
				QuestionSearchResults struct {
					IsLoading bool `json:"isLoading"`
					Paging    struct {
						IsEnd   bool `json:"isEnd"`
						IsStart bool `json:"isStart"`
						Totals  int  `json:"totals"`
					} `json:"paging"`
					Data []interface{} `json:"data"`
				} `json:"questionSearchResults"`
			} `json:"contribution"`
			CreationReferences struct {
			} `json:"creationReferences"`
			ZvideoCollection struct {
			} `json:"zvideoCollection"`
			ZvideoGrant struct {
			} `json:"zvideoGrant"`
			CollectData struct {
				IsFetching bool          `json:"isFetching"`
				List       []interface{} `json:"list"`
			} `json:"collectData"`
			VideoSource struct {
				IsLoaded bool `json:"isLoaded"`
			} `json:"videoSource"`
		} `json:"zvideos"`
	} `json:"initialState"`
	SubAppName string `json:"subAppName"`
}
