// Code generated by "stringer -type=AuthMethod"; DO NOT EDIT.

package spice

import "fmt"

const _AuthMethod_name = "AuthMethodSpiceAuthMethodSASL"

var _AuthMethod_index = [...]uint8{0, 15, 29}

func (i AuthMethod) String() string {
	i -= 1
	if i >= AuthMethod(len(_AuthMethod_index)-1) {
		return fmt.Sprintf("AuthMethod(%d)", i+1)
	}
	return _AuthMethod_name[_AuthMethod_index[i]:_AuthMethod_index[i+1]]
}
