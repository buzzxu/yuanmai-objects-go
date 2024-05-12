package main

// OAuthType represents the type of OAuth provider.
type OAuthType int32

const (
	// UNKNOWN represents an unknown OAuth provider.
	UNKNOWN OAuthType = 0

	// WEIXIN represents the WeChat public account OAuth provider.
	WEIXIN OAuthType = 1

	// WEIXIN_MINIAPP represents the WeChat mini program OAuth provider.
	WEIXIN_MINIAPP OAuthType = 2

	// QQ represents the QQ OAuth provider.
	QQ OAuthType = 3

	// ALIPAY represents the Alipay OAuth provider.
	ALIPAY OAuthType = 4

	// TAOBAO represents the Taobao OAuth provider.
	TAOBAO OAuthType = 5

	// WEIBO represents the Weibo OAuth provider.
	WEIBO OAuthType = 6

	// JPUSH represents the JPush OAuth provider.
	JPUSH OAuthType = 7

	// APPLE represents the Apple OAuth provider.
	APPLE OAuthType = 8

	// ALLINPAY represents the Allinpay OAuth provider.
	ALLINPAY OAuthType = 9

	// ALLINPAY_SUB represents the Allinpay sub account OAuth provider.
	ALLINPAY_SUB OAuthType = 901

	// WEIXIN_MOBILE represents the WeChat mobile OAuth provider.
	WEIXIN_MOBILE OAuthType = 11

	// WEIXIN_H5 represents the WeChat H5 OAuth provider.
	WEIXIN_H5 OAuthType = 12

	// YUNXIN represents the NetEase Yunxin OAuth provider.
	YUNXIN OAuthType = 13
)

// String returns the string representation of the OAuth type.
func (t OAuthType) String() string {
	switch t {
	case UNKNOWN:
		return "UNKNOWN"
	case WEIXIN:
		return "wechat"
	case WEIXIN_MINIAPP:
		return "wechat_miniapp"
	case QQ:
		return "qq"
	case ALIPAY:
		return "alipay"
	case TAOBAO:
		return "taobao"
	case WEIBO:
		return "weibo"
	case JPUSH:
		return "jpush"
	case APPLE:
		return "apple"
	case ALLINPAY:
		return "allinpay"
	case ALLINPAY_SUB:
		return "allinpay_sub"
	case WEIXIN_MOBILE:
		return "wechat_mobile"
	case WEIXIN_H5:
		return "wechat_h5"
	case YUNXIN:
		return "yunxin"
	}
	return "UNKNOWN"
}
