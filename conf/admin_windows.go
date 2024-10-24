/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2021 WireGuard LLC. All Rights Reserved.
 */

package conf

import "golang.org/x/sys/windows/registry"

const adminRegKey = `Software\AmneziaWG`

var adminKey registry.Key

func openAdminKey() (registry.Key, error) {
	if adminKey != 0 {
		return adminKey, nil
	}
	var err error
	adminKey, err = registry.OpenKey(registry.LOCAL_MACHINE, adminRegKey, registry.QUERY_VALUE|registry.WOW64_64KEY)
	if err != nil {
		return 0, err
	}
	return adminKey, nil
}

func AdminBool(name string) bool {
	key, err := openAdminKey()
	if err != nil {
		// PORTABLE: return true
		return true
	}
	val, _, err := key.GetIntegerValue(name)
	if err != nil {
		// PORTABLE: return true
		return true
	}
	return val != 0
}
