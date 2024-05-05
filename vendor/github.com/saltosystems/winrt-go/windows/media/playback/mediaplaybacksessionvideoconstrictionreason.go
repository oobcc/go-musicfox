// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package playback

type MediaPlaybackSessionVideoConstrictionReason int32

const SignatureMediaPlaybackSessionVideoConstrictionReason string = "enum(Windows.Media.Playback.MediaPlaybackSessionVideoConstrictionReason;i4)"

const (
	MediaPlaybackSessionVideoConstrictionReasonNone                      MediaPlaybackSessionVideoConstrictionReason = 0
	MediaPlaybackSessionVideoConstrictionReasonVirtualMachine            MediaPlaybackSessionVideoConstrictionReason = 1
	MediaPlaybackSessionVideoConstrictionReasonUnsupportedDisplayAdapter MediaPlaybackSessionVideoConstrictionReason = 2
	MediaPlaybackSessionVideoConstrictionReasonUnsignedDriver            MediaPlaybackSessionVideoConstrictionReason = 3
	MediaPlaybackSessionVideoConstrictionReasonFrameServerEnabled        MediaPlaybackSessionVideoConstrictionReason = 4
	MediaPlaybackSessionVideoConstrictionReasonOutputProtectionFailed    MediaPlaybackSessionVideoConstrictionReason = 5
	MediaPlaybackSessionVideoConstrictionReasonUnknown                   MediaPlaybackSessionVideoConstrictionReason = 6
)