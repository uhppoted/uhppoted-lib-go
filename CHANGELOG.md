# CHANGELOG

## [0.9.1](https://github.com/uhppoted/uhppoted-lib-go/releases/tag/v0.9.1) - 2026-10-12

### Added
1. _SetFirstCard_ API function and low-level _io_ function.
2. Add first-card privilege to `get-card` and `put-card` APIs.

### Updated
1. Updated to Go v1.27.
2. Updated to _modern_ Go with 'go fix'.
3. Retries when ephemeral UDP port collides with broadcast port.
4. Migrated from _github.com/uhppoted/uhppoted-lib-go_ to _codeberg.org/uhppoted/uhppoted-lib-go_.
5. Added retries on _send-to-self_ to UDP and TCP transports.


## [0.9.0](https://codeberg.org/uhppoted/uhppoted-lib-go/releases/tag/v0.9.0) - 2026-01-27

1. Initial release.
