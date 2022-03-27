package model

import (
	"net/http"
	"time"
)

type Item struct {
	Name string
	Link string
	Description string
	Badge string
}
var M = make(map[string][]Item)
var DFClient = &http.Client{Timeout: time.Second*10}
var Num = 20
var GithubNum = 10

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