# AmneziaWG Windows — portable storage fork

Technical dependency used by the portable AmneziaWG Windows client.

Upstream: [amnezia-vpn/amneziawg-windows](https://github.com/amnezia-vpn/amneziawg-windows)

## Purpose

The upstream Windows library stores saved tunnel configurations using Windows DPAPI (`.conf.dpapi`). That is appropriate for an installed application, but binds saved configurations to a particular Windows account/machine.

This fork keeps the upstream codebase current and applies only a small portable-storage change:

- saved tunnel configurations use `.conf`;
- configurations are stored without DPAPI encryption;
- `PathIsEncrypted` always reports `false`;
- upstream ACL/file handling is preserved unchanged.

The main portable client separately presets the root data directory to `Data\` beside `amneziawg.exe` and disables automatic migration back to DPAPI.

## Security warning

Tunnel configuration files contain private keys and are stored unencrypted. Protect the portable directory like passwords and do not place it on untrusted/shared storage.

## Legacy branch

The previous 2024 portable implementation is preserved in the `legacy-portable-2024` branch.

That older implementation included broader changes which are intentionally not carried into the current portable patch, including permissive file/directory modes and forced admin-related behavior.

## Maintenance

The portable diff should stay intentionally small. To update this fork:

1. sync `master` with the corresponding upstream `amneziawg-windows` revision;
2. reapply only the plain-`.conf` storage change in `conf/store.go`;
3. build/test together with the matching portable client revision.

## License

The upstream project is MIT-licensed. Original copyright and license notices are retained.
