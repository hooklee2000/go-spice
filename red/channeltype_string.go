// Code generated by "stringer -type=ChannelType"; DO NOT EDIT.

package red

import "fmt"

const _ChannelType_name = "ChannelMainChannelDisplayChannelInputsChannelCursorChannelPlaybackChannelRecord"

var _ChannelType_index = [...]uint8{0, 11, 25, 38, 51, 66, 79}

func (i ChannelType) String() string {
	i -= 1
	if i >= ChannelType(len(_ChannelType_index)-1) {
		return fmt.Sprintf("ChannelType(%d)", i+1)
	}
	return _ChannelType_name[_ChannelType_index[i]:_ChannelType_index[i+1]]
}
