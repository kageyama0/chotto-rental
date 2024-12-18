package auth_service

import (
	"bytes"
	"net"
)

func validateDeviceInfo(stored DeviceInfo, current DeviceInfo) bool {
	// 基本的な検証ポイント
	if stored.UserAgent != current.UserAgent {
		return false
	}

	// IPアドレスの検証
	// 完全一致を求めるのは厳しすぎる可能性があるため、
	// モバイル回線などを考慮してプライベートIPの場合のみ厳密にチェック
	if isPrivateIP(net.ParseIP(stored.IP)) && isPrivateIP(net.ParseIP(current.IP)) {
		if stored.IP != current.IP {
			return false
		}
	}

	return true
}

// isPrivateIP checks if an IP address is private
func isPrivateIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	// プライベートIPの範囲をチェック
	privateIPRanges := []struct {
		start net.IP
		end   net.IP
	}{
		{
			net.ParseIP("10.0.0.0"),
			net.ParseIP("10.255.255.255"),
		},
		{
			net.ParseIP("172.16.0.0"),
			net.ParseIP("172.31.255.255"),
		},
		{
			net.ParseIP("192.168.0.0"),
			net.ParseIP("192.168.255.255"),
		},
	}

	for _, r := range privateIPRanges {
		if bytes.Compare(ip, r.start) >= 0 && bytes.Compare(ip, r.end) <= 0 {
			return true
		}
	}

	return false
}

// より厳密な検証が必要な場合は、以下のようなオプションを追加できます
// type DeviceInfoValidationOptions struct {
//     ValidateIP      bool
//     ValidateUA      bool
//     IPMatchRequired bool
//     AllowedIPRange  []string
// }

// func validateDeviceInfoWithOptions(stored DeviceInfo, current DeviceInfo, opts DeviceInfoValidationOptions) (bool, error) {
//     if opts.ValidateUA && stored.UserAgent != current.UserAgent {
//         return false, errors.New("user agent mismatch")
//     }

//     if opts.ValidateIP {
//         storedIP := net.ParseIP(stored.IP)
//         currentIP := net.ParseIP(current.IP)

//         if storedIP == nil || currentIP == nil {
//             return false, errors.New("invalid IP format")
//         }

//         if opts.IPMatchRequired {
//             if !storedIP.Equal(currentIP) {
//                 return false, errors.New("IP address mismatch")
//             }
//         } else if isPrivateIP(storedIP) && isPrivateIP(currentIP) {
//             // プライベートIPの場合は厳密にチェック
//             if !storedIP.Equal(currentIP) {
//                 return false, errors.New("private IP address mismatch")
//             }
//         }
//     }

//     return true, nil
// }
