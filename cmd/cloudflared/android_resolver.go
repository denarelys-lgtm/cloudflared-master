//go:build android

package main

import "net"

// Android uses Bionic's resolver through cgo. The pure-Go resolver reads
// resolver configuration files that are not authoritative on Android and can
// fall back to a loopback DNS address such as [::1]:53. Keep Go's cgo resolver
// selected so hostname lookups use Android's active Network/DNS configuration
// (including Private DNS and VPN-provided DNS).
func init() {
    net.DefaultResolver = &net.Resolver{
        PreferGo:     false,
        StrictErrors: false,
    }
}
