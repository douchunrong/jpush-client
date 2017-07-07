package push

const (
	// all audiences, all platforms
	ALL = "all"

	PLATFORM_IOS     = "ios"
	PLATFORM_ANDROID = "android"
	PLATFORM_WP      = "winphone"

	IosNotificationMaxSize = 2000

	MaxTimeToLive      = 10 * 24 * 60 * 60 // 10 天
	MaxBigPushDuration = 1400
)
