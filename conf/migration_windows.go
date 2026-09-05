/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019-2021 WireGuard LLC. All Rights Reserved.
 */

package conf

// MigrationCallback is kept for API compatibility with the upstream Windows client.
type MigrationCallback func(name, oldPath, newPath string)

// MigrateUnencryptedConfigs is intentionally a no-op in the portable build.
// Portable tunnel configurations are stored as plain .conf files and must not
// be migrated to DPAPI-encrypted .conf.dpapi files.
func MigrateUnencryptedConfigs(migrated MigrationCallback) {}
