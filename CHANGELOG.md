# Changelog

All notable changes to this project will be documented in this file.

## [1.5.3](https://github.com/equinor/radix-cicd-canary/compare/v1.5.2..v1.5.3) - 2026-02-16

### 🐛 Bug Fixes

- *(deps)* Update Go version to 1.25.7 (#267) - ([4c53c47](https://github.com/equinor/radix-cicd-canary/commit/4c53c47c0eb247e3a0f2db091ab0649b9b88c944)) by @nilsgstrabo in [#267](https://github.com/equinor/radix-cicd-canary/pull/267)


### 📚 Documentation

- Update catalog-info.yaml (#266) - ([2b4fcc5](https://github.com/equinor/radix-cicd-canary/commit/2b4fcc5e65a6a1f940d30a1eec46baf206bba026)) by @emirgens in [#266](https://github.com/equinor/radix-cicd-canary/pull/266)


## [1.5.2](https://github.com/equinor/radix-cicd-canary/compare/v1.5.1..v1.5.2) - 2026-01-09

### 🐛 Bug Fixes

- *(deps)* Update to go 1.25 - ([41d4f10](https://github.com/equinor/radix-cicd-canary/commit/41d4f106c2eb6056f4ea998f409952a08f4e6278)) by @nilsgstrabo in [#261](https://github.com/equinor/radix-cicd-canary/pull/261)


### 📚 Documentation

- Add catalog-info.yaml config file (#260) - ([8e3888b](https://github.com/equinor/radix-cicd-canary/commit/8e3888be9994c50a2e6bb3ff289bb5f60f160caf)) by @emirgens in [#260](https://github.com/equinor/radix-cicd-canary/pull/260)


### ⚙️ Miscellaneous Tasks

- Configure concurrency for prepare-release workflow to prevent concurrent runs - ([ecd4dbe](https://github.com/equinor/radix-cicd-canary/commit/ecd4dbe2ca6e0d8dd9a8f9e44be0fdd756633db5)) by @nilsgstrabo in [#258](https://github.com/equinor/radix-cicd-canary/pull/258)


## [1.5.1](https://github.com/equinor/radix-cicd-canary/compare/v1.5.0..v1.5.1) - 2025-08-26

### 🐛 Bug Fixes

- Regenerate radix-api client (#251) - ([ed914de](https://github.com/equinor/radix-cicd-canary/commit/ed914defa41bd20b18613b98aee2f48f6b197ef6)) by @nilsgstrabo in [#251](https://github.com/equinor/radix-cicd-canary/pull/251)


### 📚 Documentation

- Update development and release documentation - ([69499b6](https://github.com/equinor/radix-cicd-canary/commit/69499b65b40c9c5f429218a440ff66e4c3d85109)) by @nilsgstrabo in [#248](https://github.com/equinor/radix-cicd-canary/pull/248)


### ⚙️ Miscellaneous Tasks

- Remove deprecated "build & push" workflow (#249) - ([4d4ba45](https://github.com/equinor/radix-cicd-canary/commit/4d4ba45b3652b7241091f8a1fa206d8b600e6258)) by @nilsgstrabo in [#249](https://github.com/equinor/radix-cicd-canary/pull/249)

- Use workflow token when uploading chart artifact to release (#250) - ([85961c8](https://github.com/equinor/radix-cicd-canary/commit/85961c819fb8ef63405e7982f58dbb62aa7f2553)) by @nilsgstrabo in [#250](https://github.com/equinor/radix-cicd-canary/pull/250)

- Bump docker/build-push-action to v6 - ([ec50704](https://github.com/equinor/radix-cicd-canary/commit/ec50704d86b50aff5cd554410f1851468303bd09)) by @nilsgstrabo in [#254](https://github.com/equinor/radix-cicd-canary/pull/254)


## [1.5.0](https://github.com/equinor/radix-cicd-canary/compare/v1.4.4..v1.5.0) - 2025-08-15

### 🚀 Features

- *(ci)* Refactor release and deploy workflow - ([b94f135](https://github.com/equinor/radix-cicd-canary/commit/b94f13593b01f3e5761ba9fac75c735eeae53025)) by @nilsgstrabo in [#241](https://github.com/equinor/radix-cicd-canary/pull/241)


### ⚙️ Miscellaneous Tasks

- Bump release workflows version - ([7742c44](https://github.com/equinor/radix-cicd-canary/commit/7742c446e1f4f0ca8ae9bdefa4654795fe1aa80e)) by @nilsgstrabo in [#242](https://github.com/equinor/radix-cicd-canary/pull/242)

- Fix repository name in deploy workflow (#245) - ([0eb8f7b](https://github.com/equinor/radix-cicd-canary/commit/0eb8f7bc3fb0cd2f7286c57176049aeed084d056)) by @nilsgstrabo in [#245](https://github.com/equinor/radix-cicd-canary/pull/245)


## New Contributors ❤️

* @github-actions[bot] made their first contribution in [#244](https://github.com/equinor/radix-cicd-canary/pull/244)
## [1.4.4] - 2025-07-30

### 🚀 Features

- Refactor tests to match different response codes (#217) - ([7bfe071](https://github.com/equinor/radix-cicd-canary/commit/7bfe0716ca9c528e692e58cf0a621d85792126e9)) by @Richard87 in [#217](https://github.com/equinor/radix-cicd-canary/pull/217)

- Change sysId to appId for consistency (#237) - ([5a5da27](https://github.com/equinor/radix-cicd-canary/commit/5a5da277fb258a4c680364310a8cc1d2257f3e29)) by @Richard87 in [#237](https://github.com/equinor/radix-cicd-canary/pull/237)


### 🐛 Bug Fixes

- PrivateImageHubs return Forbidden instead of BadRequest in Dev (#215) - ([7ea2310](https://github.com/equinor/radix-cicd-canary/commit/7ea23106ac749b3d828573750f12bbd8f8a0ca41)) by @Richard87 in [#215](https://github.com/equinor/radix-cicd-canary/pull/215)

- Delete non existant apps (#234) - ([5e05e69](https://github.com/equinor/radix-cicd-canary/commit/5e05e69fbe7121d3eb0cc8422258bc1028b66832)) by @Richard87 in [#234](https://github.com/equinor/radix-cicd-canary/pull/234)


### ⚙️ Miscellaneous Tasks

- Update go dependencies (#229) - ([3a5785e](https://github.com/equinor/radix-cicd-canary/commit/3a5785e2f89c1583c11b6c37b443531c2d5a6f12)) by @nilsgstrabo in [#229](https://github.com/equinor/radix-cicd-canary/pull/229)


## New Contributors ❤️

* @Richard87 made their first contribution in [#237](https://github.com/equinor/radix-cicd-canary/pull/237)
* @satr made their first contribution in [#232](https://github.com/equinor/radix-cicd-canary/pull/232)
* @nilsgstrabo made their first contribution in [#229](https://github.com/equinor/radix-cicd-canary/pull/229)
* @emirgens made their first contribution in [#208](https://github.com/equinor/radix-cicd-canary/pull/208)
* @dependabot[bot] made their first contribution in [#200](https://github.com/equinor/radix-cicd-canary/pull/200)
* @magnus-longva-bouvet made their first contribution in [#171](https://github.com/equinor/radix-cicd-canary/pull/171)
* @anneliawa made their first contribution in [#164](https://github.com/equinor/radix-cicd-canary/pull/164)
* @oterno made their first contribution in [#117](https://github.com/equinor/radix-cicd-canary/pull/117)
* @ made their first contribution
* @sveinpj made their first contribution in [#89](https://github.com/equinor/radix-cicd-canary/pull/89)
* @JoakimHagen made their first contribution in [#53](https://github.com/equinor/radix-cicd-canary/pull/53)
* @ingeknudsen made their first contribution
* @thezultimate made their first contribution in [#48](https://github.com/equinor/radix-cicd-canary/pull/48)
* @keaaa made their first contribution in [#36](https://github.com/equinor/radix-cicd-canary/pull/36)
* @nemzes made their first contribution in [#19](https://github.com/equinor/radix-cicd-canary/pull/19)
* @kjellerik made their first contribution
<!-- generated by git-cliff -->
