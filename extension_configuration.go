package helix

import "fmt"

// SegmentType A segment configuration type
type ExtensionSegmentType string

// Types of segments datastores for the configuration service
const (
	BroadcasterSegment ExtensionSegmentType = "broadcaster"
	DeveloperSegment   ExtensionSegmentType = "developer"
	GlobalSegment      ExtensionSegmentType = "global"
)

func (s ExtensionSegmentType) String() string {
	return string(s)
}

type ExtensionConfigurationParams struct {
	Segment       ExtensionSegmentType `json:"segment"`
	ExtensionId   string               `json:"extension-id"`
	BroadcasterID string               `json:"broadcaster_id,omitempty"` // populated if segment is of type 'developer' || 'broadcaster'
	Version       string               `json:"version"`
	Content       string               `json:"content"`
}

type ExtensionConfigurationSegment struct {
	Segment ExtensionSegmentType `json:"segment"`
	Version string               `json:"version"`
	Content string               `json:"content"`
}

type ExtensionGetConfigurationParams struct {
	ExtensionID   string                 `query:"extension_id"`
	BroadcasterID string                 `query:"broadcaster_id"`
	Segment       []ExtensionSegmentType `query:"segment"`
}

type ExtensionSetRequiredConfigurationParams struct {
	ExtensionID           string `json:"extension_id"`
	ExtensionVersion      string `json:"extension_version"`
	RequiredConfiguration string `json:"required_configuration"`
}

type ExtensionSetRequiredConfigurationResponse struct {
	ResponseCommon
}

type ExtensionGetConfigurationSegmentResponse struct {
	ResponseCommon
	Data ManyExtensionConfigurationSegments
}

type ManyExtensionConfigurationSegments struct {
	Segments []ExtensionConfigurationSegment `json:"data"`
}

type ExtensionSetConfigurationResponse struct {
	ResponseCommon
}

// https://dev.twitch.tv/docs/extensions/reference/#set-extension-configuration-segment
func (c *Client) SetExtensionSegmentConfig(params *ExtensionConfigurationParams) (*ExtensionSetConfigurationResponse, error) {
	if params.BroadcasterID != "" {
		switch params.Segment {
		case DeveloperSegment, BroadcasterSegment:
		default:
			return nil, fmt.Errorf("error: developer or broadcaster extension configuration segment type must be provided")
		}
	}

	resp, err := c.putAsJSON("/extensions/configurations", &ManyPolls{}, params)
	if err != nil {
		return nil, err
	}

	setExtCnfgResp := &ExtensionSetConfigurationResponse{}
	resp.HydrateResponseCommon(&setExtCnfgResp.ResponseCommon)

	return setExtCnfgResp, nil
}

func (c *Client) GetExtensionConfigurationSegment(params *ExtensionGetConfigurationParams) (*ExtensionGetConfigurationSegmentResponse, error) {
	resp, err := c.get("/extensions/configurations", &ManyExtensionConfigurationSegments{}, params)
	if err != nil {
		return nil, err
	}

	extCfgSegResp := &ExtensionGetConfigurationSegmentResponse{}
	resp.HydrateResponseCommon(&extCfgSegResp.ResponseCommon)
	extCfgSegResp.Data.Segments = resp.Data.(*ManyExtensionConfigurationSegments).Segments

	return extCfgSegResp, nil
}

func (c *Client) SetExtensionRequiredConfiguration(params *ExtensionSetRequiredConfigurationParams) (*ExtensionSetRequiredConfigurationResponse, error) {

	resp, err := c.putAsJSON("/extensions/configurations/required_configuration", &ExtensionSetRequiredConfigurationResponse{}, params)
	if err != nil {
		return nil, err
	}

	extReqCfgResp := &ExtensionSetRequiredConfigurationResponse{}
	resp.HydrateResponseCommon(&extReqCfgResp.ResponseCommon)

	return extReqCfgResp, nil
}