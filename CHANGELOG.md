# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.1] - 2021-07-10

### Fixed

- crashes related to path handling on Windows

## [0.1.0] - 2021-07-09

### Added

#### For users
- Works on MacOS, GNU/Linux and Windows, tested on:
- Manjaro Linux Omara 21.0.7 (with kernel 5.12.9-1-MANJARO)
- MacOS Catalina 10.15.7
- Windows 10 20H2 (Build 19042.1083) (Please use the new Windows Terminal or something else that support ANSI escape sequences)
- a `use` commands to download & install themes
- works with any remote git repository
- works with any URL poiting to a .zip file
- shorthand for github repositories: OWNER/REPO
- a `init` command to add a `ffcss.yaml` manifest in your current repository
- basic for now, [a smarter version is planned](https://github.com/ewen-lbh/ffcss/issues/20)
- a `cache clear` command to clear the cache of downloaded repositories
- 8 themes available out-of-the-box (use them by typing their name only, it works)
- [chameleons-beauty](https://github.com/Godiesc/Chameleons-Beauty) by [Godiesc](https://github.com/Godiesc)
- [fxcompact](https://github.com/dannycolin/fx-compact-mode) by [dannycolin](https://github.com/dannycolin)
- [lepton](https://github.com/black7375/Firefox-UI-Fix) by [black7375](https://github.com/black7375)
- [materialfox](https://github.com/muckSponge/MaterialFox) by [muckSponge](https://github.com/muckSponge)
- [modoki](https://github.com/soup-bowl/Modoki-FirefoxCSS) by [soup-bowl](https://github.com/soup-bowl)
- [simplerentfox](https://github.com/MiguelRAvila/SimplerentFox) by [MiguelRAvila](https://github.com/MiguelRAvila)
- [verticaltabs](https://github.com/ranmaru22/firefox-vertical-tabs) by [ranmaru22](https://github.com/ranmaru22)
#### For theme makers
- a mechanism to handle theme variants:
- Variants can be declared in the same manifest file under the `variants` entry to override other entries
- per-OS paths: the {{ variant }} and {{ os }} placeholders get replaced with their values
- the value {{ os }} gets replaced with can be customized in the manifest file under the `os` entry, use `null` to mark an OS as incompatible
- Support for helper addons:
- Declare URLs to open after installation under the `addons` manifest entry (I plan to auto-install them in the future)
- Easy way to write about:config changes without writing a `user.js` file:
- Use the `config` manifest entry
- If you also use a `user.js`, you can combine both, they'll be written as a single `user.js` file
- Support for custom assets:
- Use the `assets` manifest entry to list out your assets
- Supports glob patterns
- If you store them under a `chrome` directory, you can use `copy from: chrome/` so that they don't get copied to `<profile directory>/chrome/chrome`

[Unreleased]: https://github.com/ewen-lbh/ffcss/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/ewen-lbh/ffcss/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/ewen-lbh/ffcss/releases/tag/v0.1.0

[//]: # (C3-2-DKAC:GGH:Rewen-lbh/ffcss:Tv{t})