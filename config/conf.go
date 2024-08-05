package config

type Config struct {
	AppID       string `json:"appid"`
	AppSecret   string `json:"appsecret"`
	AccessToken string `json:"access_token"`
}

var Conf = Config{
	AppID:     "wx7dc67faba9359f5e",
	AppSecret: "9c73b9258d648318c84f5e8ea0fc26b6",
}

const UPLOAD_MEDIA_URL = "http://file.api.weixin.qq.com/cgi-bin"
const API_BASE_URL_PREFIX = "https://api.weixin.qq.com" //以下API接口URL需要使用此前缀
const QRCODE_IMG_URL = "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket="
const API_URL_PREFIX = "https://api.weixin.qq.com/cgi-bin"
const OAUTH_PREFIX = "https://open.weixin.qq.com/connect/oauth2"
const MSGTYPE_TEXT = "text"
const MSGTYPE_IMAGE = "image"
const MSGTYPE_LOCATION = "location"
const MSGTYPE_LINK = "link"
const MSGTYPE_EVENT = "event"
const MSGTYPE_MUSIC = "music"
const MSGTYPE_NEWS = "news"
const MSGTYPE_VOICE = "voice"
const MSGTYPE_VIDEO = "video"
const EVENT_SUBSCRIBE = "subscribe"                 //订阅
const EVENT_UNSUBSCRIBE = "unsubscribe"             //取消订阅
const EVENT_SCAN = "SCAN"                           //扫描带参数二维码
const EVENT_LOCATION = "LOCATION"                   //上报地理位置
const EVENT_MENU_VIEW = "VIEW"                      //菜单 - 点击菜单跳转链接
const EVENT_MENU_CLICK = "CLICK"                    //菜单 - 点击菜单拉取消息
const EVENT_MENU_SCAN_PUSH = "scancode_push"        //菜单 - 扫码推事件(客户端跳URL)
const EVENT_MENU_SCAN_WAITMSG = "scancode_waitmsg"  //菜单 - 扫码推事件(客户端不跳URL)
const EVENT_MENU_PIC_SYS = "pic_sysphoto"           //菜单 - 弹出系统拍照发图
const EVENT_MENU_PIC_PHOTO = "pic_photo_or_album"   //菜单 - 弹出拍照或者相册发图
const EVENT_MENU_PIC_WEIXIN = "pic_weixin"          //菜单 - 弹出微信相册发图器
const EVENT_MENU_LOCATION = "location_select"       //菜单 - 弹出地理位置选择器
const EVENT_SEND_MASS = "MASSSENDJOBFINISH"         //发送结果 - 高级群发完成
const EVENT_SEND_TEMPLATE = "TEMPLATESENDJOBFINISH" //发送结果 - 模板消息发送结果
const EVENT_KF_SEESION_CREATE = "kfcreatesession"   //多客服 - 接入会话
const EVENT_KF_SEESION_CLOSE = "kfclosesession"     //多客服 - 关闭会话
const EVENT_KF_SEESION_SWITCH = "kfswitchsession"   //多客服 - 转接会话
const EVENT_CARD_PASS = "card_pass_check"           //卡券 - 审核通过
const EVENT_CARD_NOTPASS = "card_not_pass_check"    //卡券 - 审核未通过
const EVENT_CARD_USER_GET = "user_get_card"         //卡券 - 用户领取卡券
const EVENT_CARD_USER_DEL = "user_del_card"         //卡券 - 用户删除卡券
const AUTH_URL = "/token?grant_type=client_credential&"
const MENU_CREATE_URL = "/menu/create?"
const MENU_GET_URL = "/menu/get?"
const MENU_DELETE_URL = "/menu/delete?"
const GET_TICKET_URL = "/ticket/getticket?"
const CALLBACKSERVER_GET_URL = "/getcallbackip?"
const QRCODE_CREATE_URL = "/qrcode/create?"
const QR_SCENE = 0
const QR_LIMIT_SCENE = 1
const SHORT_URL = "/shorturl?"
const USER_GET_URL = "/user/get?"
const USER_INFO_URL = "/user/info?"
const USER_UPDATEREMARK_URL = "/user/info/updateremark?"
const GROUP_GET_URL = "/groups/get?"
const USER_GROUP_URL = "/groups/getid?"
const GROUP_CREATE_URL = "/groups/create?"
const GROUP_UPDATE_URL = "/groups/update?"
const GROUP_MEMBER_UPDATE_URL = "/groups/members/update?"
const GROUP_MEMBER_BATCHUPDATE_URL = "/groups/members/batchupdate?"
const CUSTOM_SEND_URL = "/message/custom/send?"
const MEDIA_UPLOADNEWS_URL = "/media/uploadnews?"
const MASS_SEND_URL = "/message/mass/send?"
const TEMPLATE_SET_INDUSTRY_URL = "/template/get_industry??"
const TEMPLATE_ADD_TPL_URL = "/template/api_add_template?"
const GET_All_PRIVATE_TEMPLATE_URL = "/template/get_all_private_template?"
const DEL_PRIVATE_TEMPLATE_URL = "/template/del_private_template?"
const TEMPLATE_SEND_URL = "/message/template/send?"
const MASS_SEND_GROUP_URL = "/message/mass/sendall?"
const MASS_DELETE_URL = "/message/mass/delete?"
const MASS_PREVIEW_URL = "/message/mass/preview?"
const MASS_QUERY_URL = "/message/mass/get?"
const MEDIA_UPLOAD_URL = "/media/upload?"
const MEDIA_GET_URL = "/media/get?"
const MEDIA_VIDEO_UPLOAD = "/media/uploadvideo?"
const OAUTH_AUTHORIZE_URL = "/authorize?"

// /多客服相关地址
const CUSTOM_SERVICE_GET_RECORD = "/customservice/getrecord?"
const CUSTOM_SERVICE_GET_KFLIST = "/customservice/getkflist?"
const CUSTOM_SERVICE_GET_ONLINEKFLIST = "/customservice/getonlinekflist?"
const OAUTH_TOKEN_URL = "/sns/oauth2/access_token?"
const OAUTH_REFRESH_URL = "/sns/oauth2/refresh_token?"
const OAUTH_USERINFO_URL = "/sns/userinfo?"
const OAUTH_AUTH_URL = "/sns/auth?"

// /多客服相关地址
const CUSTOM_SESSION_CREATE = "/customservice/kfsession/create?"
const CUSTOM_SESSION_CLOSE = "/customservice/kfsession/close?"
const CUSTOM_SESSION_SWITCH = "/customservice/kfsession/switch?"
const CUSTOM_SESSION_GET = "/customservice/kfsession/getsession?"
const CUSTOM_SESSION_GET_LIST = "/customservice/kfsession/getsessionlist?"
const CUSTOM_SESSION_GET_WAIT = "/customservice/kfsession/getwaitcase?"
const CS_KF_ACCOUNT_ADD_URL = "/customservice/kfaccount/add?"
const CS_KF_ACCOUNT_UPDATE_URL = "/customservice/kfaccount/update?"
const CS_KF_ACCOUNT_DEL_URL = "/customservice/kfaccount/del?"
const CS_KF_ACCOUNT_UPLOAD_HEADIMG_URL = "/customservice/kfaccount/uploadheadimg?"

// /卡券相关地址
const CARD_CREATE = "/card/create?"
const CARD_DELETE = "/card/delete?"
const CARD_UPDATE = "/card/update?"
const CARD_GET = "/card/get?"
const CARD_BATCHGET = "/card/batchget?"
const CARD_MODIFY_STOCK = "/card/modifystock?"
const CARD_LOCATION_BATCHADD = "/card/location/batchadd?"
const CARD_LOCATION_BATCHGET = "/card/location/batchget?"
const CARD_GETCOLORS = "/card/getcolors?"
const CARD_QRCODE_CREATE = "/card/qrcode/create?"
const CARD_CODE_CONSUME = "/card/code/consume?"
const CARD_CODE_DECRYPT = "/card/code/decrypt?"
const CARD_CODE_GET = "/card/code/get?"
const CARD_CODE_UPDATE = "/card/code/update?"
const CARD_CODE_UNAVAILABLE = "/card/code/unavailable?"
const CARD_TESTWHILELIST_SET = "/card/testwhitelist/set?"
const CARD_MEMBERCARD_ACTIVATE = "/card/membercard/activate?"        //激活会员卡
const CARD_MEMBERCARD_UPDATEUSER = "/card/membercard/updateuser?"    //更新会员卡
const CARD_MOVIETICKET_UPDATEUSER = "/card/movieticket/updateuser?"  //更新电影票(未加方法)
const CARD_BOARDINGPASS_CHECKIN = "/card/boardingpass/checkin?"      //飞机票-在线选座(未加方法)
const CARD_LUCKYMONEY_UPDATE = "/card/luckymoney/updateuserbalance?" //更新红包金额
const SEMANTIC_API_URL = "/semantic/semproxy/search?"                //语义理解
// 订阅通知
const SUB_BIZSEND = "/message/subscribe/bizsend?" //发送订阅通知
