package helix

import "github.com/stretchr/testify/mock"

type MockClient struct {
	mock.Mock
}

type Service interface {
	GetAuthorizationURL(params *AuthorizationURLParams) string
	RequestAppAccessToken(scopes []string) (*AppAccessTokenResponse, error)
	RequestUserAccessToken(code string) (*UserAccessTokenResponse, error)
	RefreshUserAccessToken(refreshToken string) (*RefreshTokenResponse, error)
	RevokeUserAccessToken(accessToken string) (*RevokeAccessTokenResponse, error)
	ValidateToken(accessToken string) (bool, *ValidateTokenResponse, error)
	StartCommercial(params *StartCommercialParams) (*StartCommercialResponse, error)
	GetExtensionAnalytics(params *ExtensionAnalyticsParams) (*ExtensionAnalyticsResponse, error)
	GetGameAnalytics(params *GameAnalyticsParams) (*GameAnalyticsResponse, error)
	GetCheermotes(params *CheermotesParams) (*CheermotesResponse, error)
	GetBitsLeaderboard(params *BitsLeaderboardParams) (*BitsLeaderboardResponse, error)

	SearchCategories(params *SearchCategoriesParams) (*SearchCategoriesResponse, error)

	SearchChannels(params *SearchChannelsParams) (*SearchChannelsResponse, error)
	GetChannelInformation(params *GetChannelInformationParams) (*GetChannelInformationResponse, error)
	EditChannelInformation(params *EditChannelInformationParams) (*EditChannelInformationResponse, error)
	GetChannelFollows(params *GetChannelFollowsParams) (*GetChannelFollowersResponse, error)
	GetFollowedChannels(params *GetFollowedChannelParams) (*GetFollowedChannelResponse, error)

	GetChannelEditors(params *ChannelEditorsParams) (*ChannelEditorsResponse, error)

	CreateCustomReward(params *ChannelCustomRewardsParams) (*ChannelCustomRewardResponse, error)
	UpdateCustomReward(params *UpdateChannelCustomRewardsParams) (*ChannelCustomRewardResponse, error)
	DeleteCustomRewards(params *DeleteCustomRewardsParams) (*DeleteCustomRewardsResponse, error)
	GetCustomRewards(params *GetCustomRewardsParams) (*ChannelCustomRewardResponse, error)
	UpdateChannelCustomRewardsRedemptionStatus(params *UpdateChannelCustomRewardsRedemptionStatusParams) (*ChannelCustomRewardsRedemptionResponse, error)
}
