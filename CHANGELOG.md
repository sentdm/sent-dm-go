# Changelog

## [0.33.0](https://github.com/sentdm/sent-dm-go/compare/v0.32.0...v0.33.0) (2026-09-06)


### Features

* **api:** sync OpenAPI spec from production ([60ebd92](https://github.com/sentdm/sent-dm-go/commit/60ebd925ca6cfa0cab410556c9c254cb6b82c82a))
* **api:** sync OpenAPI spec from production ([f6185e5](https://github.com/sentdm/sent-dm-go/commit/f6185e59aaf68b5a0b274cb41bdbdedaf77c54fb))

## [0.32.0](https://github.com/sentdm/sent-dm-go/compare/v0.31.0...v0.32.0) (2026-08-31)


### Features

* **api:** sync OpenAPI spec from production ([66073c8](https://github.com/sentdm/sent-dm-go/commit/66073c8deaa3f3276c982f44004c797c4601f241))
* **api:** sync OpenAPI spec from production ([3e03c72](https://github.com/sentdm/sent-dm-go/commit/3e03c726bbeb4d0a9bf282ed97e70a52ca3822c6))

## [0.31.0](https://github.com/sentdm/sent-dm-go/compare/v0.30.0...v0.31.0) (2026-08-17)


### Highlights

Webhook payloads are now typed. The events Sent POSTs to your endpoint — `MessageEvent`, `InboundMessageEvent` and `TemplateEvent`, each with its own payload type — are generated types you can deserialize into, instead of a shape you had to hand-write from the docs.

The webhook delivery log is typed too. `event_data` on `GET /v3/webhooks/{id}/events` returns the exact envelope that was delivered, and now describes itself as one of those three rather than an opaque object.

Also in this release:

- `csp_id` on the brand object is deprecated and will be removed in a later release. It identifies the Campaign Service Provider that registered the brand, which is Sent, so the value is the same for every account. There is no replacement. Your own TCR identifiers, `tcr_brand_id` and `universal_ein`, are unaffected.
- Corrected descriptions for blocked sends, which now name the cases that gate a send before any delivery attempt: insufficient balance, a template not approved for sending, and free-form content with no open conversation.
- `campaign.volume` documents what an omitted value does. Leave it out and the campaign registers as standard, the higher-fee tier, with no error.

### Features

* **api:** sync OpenAPI spec from production ([8984ed2](https://github.com/sentdm/sent-dm-go/commit/8984ed242dbe2bf5f4b69d79b27b7a1922d424f7))
* **api:** sync OpenAPI spec from production ([c6ea4a7](https://github.com/sentdm/sent-dm-go/commit/c6ea4a7d760911a5cf7ea155036d2f0301350a8e))
* **api:** sync OpenAPI spec from production ([1277e78](https://github.com/sentdm/sent-dm-go/commit/1277e789a770bf4f302765d33edd24f943744c0d))


### Chores

* add eager seal-dispatch workflow ([4b9a291](https://github.com/sentdm/sent-dm-go/commit/4b9a29111504023c050963bb225c449f090e7232))

## [0.30.0](https://github.com/sentdm/sent-dm-go/compare/v0.29.0...v0.30.0) (2026-08-08)


### Features

* **api:** sync OpenAPI spec from production ([633eea0](https://github.com/sentdm/sent-dm-go/commit/633eea0aef26ed5150ff8ef1cff565ff9ea0ef89))


### Chores

* mark GitHub Releases as stable releases (prerelease: false) ([385ba77](https://github.com/sentdm/sent-dm-go/commit/385ba77d8244388a9fee7e11b0d726eb76ec6a3a))

## [0.29.0](https://github.com/sentdm/sent-dm-go/compare/v0.28.0...v0.29.0) (2026-07-07)


### Features

* enable release-please releases and back-sync trigger ([80b99f1](https://github.com/sentdm/sent-dm-go/commit/80b99f161e1db4edcf5bfe1560806dc6d887ceb2))
* initial stlc build ([7ece6f0](https://github.com/sentdm/sent-dm-go/commit/7ece6f0ebc4f0d8cc6bd847dcecf5b496d3df3aa))


### Chores

* add promote, back-sync, and trunk-lock workflows ([67a9fef](https://github.com/sentdm/sent-dm-go/commit/67a9fef051a09f96433d2a0be0a3f92701d20524))
* add release back-sync trigger workflow ([3acc2f7](https://github.com/sentdm/sent-dm-go/commit/3acc2f7ee0bf99c217448a7d639d51aaf288f5ff))

## 0.28.0 (2026-07-02)

Full Changelog: [v0.27.0...v0.28.0](https://github.com/sentdm/sent-dm-go/compare/v0.27.0...v0.28.0)

### Features

* **api:** api update ([5a0ff90](https://github.com/sentdm/sent-dm-go/commit/5a0ff9055eed066c1f5e6bd617a270f377b56d33))

## 0.27.0 (2026-06-30)

Full Changelog: [v0.26.0...v0.27.0](https://github.com/sentdm/sent-dm-go/compare/v0.26.0...v0.27.0)

### Features

* **api:** api update ([c78bb83](https://github.com/sentdm/sent-dm-go/commit/c78bb83a27e52df83f99e3bc9435c9404b2eea6b))
* **api:** api update ([c122082](https://github.com/sentdm/sent-dm-go/commit/c122082b2af934160f2e55800e0e5bda8d870ff5))

## 0.26.0 (2026-05-21)

Full Changelog: [v0.25.0...v0.26.0](https://github.com/sentdm/sent-dm-go/compare/v0.25.0...v0.26.0)

### Features

* **api:** api update ([01ebf2d](https://github.com/sentdm/sent-dm-go/commit/01ebf2d0f1d0eae2554674cb87fbbb2117106d01))
* **api:** api update ([4643ffa](https://github.com/sentdm/sent-dm-go/commit/4643ffabbfc4b6e5cdf2323b3251494476d01abb))

## 0.25.0 (2026-05-14)

Full Changelog: [v0.24.0...v0.25.0](https://github.com/sentdm/sent-dm-go/compare/v0.24.0...v0.25.0)

### Features

* **api:** manual updates ([d4684fc](https://github.com/sentdm/sent-dm-go/commit/d4684fc190b7d7b6e44a6c654b61c1dec85b3541))

## 0.24.0 (2026-05-14)

Full Changelog: [v0.23.0...v0.24.0](https://github.com/sentdm/sent-dm-go/compare/v0.23.0...v0.24.0)

### Features

* **api:** manual updates ([1bcde22](https://github.com/sentdm/sent-dm-go/commit/1bcde22513488432e3d610b7c0b527b7c39ccfec))

## 0.23.0 (2026-05-14)

Full Changelog: [v0.22.2...v0.23.0](https://github.com/sentdm/sent-dm-go/compare/v0.22.2...v0.23.0)

### Features

* **api:** api update ([bc87961](https://github.com/sentdm/sent-dm-go/commit/bc879613d047e25034ca8e0ae2683a7420d285ed))
* **client:** optimize json encoder for internal types ([9fc2d0e](https://github.com/sentdm/sent-dm-go/commit/9fc2d0e67692fbd5a9111b3ee51c1ad48b67261f))

## 0.22.2 (2026-05-08)

Full Changelog: [v0.22.1...v0.22.2](https://github.com/sentdm/sent-dm-go/compare/v0.22.1...v0.22.2)

### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([231d7a2](https://github.com/sentdm/sent-dm-go/commit/231d7a26180e5cf850d66da96bcca7312e569780))


### Chores

* redact api-key headers in debug logs ([d0e58f4](https://github.com/sentdm/sent-dm-go/commit/d0e58f4e70f282ed6a03df7747e9fa67cb08d642))

## 0.22.1 (2026-05-01)

Full Changelog: [v0.22.0...v0.22.1](https://github.com/sentdm/sent-dm-go/compare/v0.22.0...v0.22.1)

### Chores

* avoid embedding reflect.Type for dead code elimination ([7b390c2](https://github.com/sentdm/sent-dm-go/commit/7b390c2c05ae260d8c7babcaf58acb3389bb5cfd))

## 0.22.0 (2026-04-29)

Full Changelog: [v0.21.0...v0.22.0](https://github.com/sentdm/sent-dm-go/compare/v0.21.0...v0.22.0)

### Features

* **api:** api update ([9e2ab18](https://github.com/sentdm/sent-dm-go/commit/9e2ab18b9e3e11c111e28c2ba1abf13cf378b071))

## 0.21.0 (2026-04-29)

Full Changelog: [v0.20.0...v0.21.0](https://github.com/sentdm/sent-dm-go/compare/v0.20.0...v0.21.0)

### Features

* **api:** manual updates ([4b590a5](https://github.com/sentdm/sent-dm-go/commit/4b590a5906789ea65944e364958a85d243429289))

## 0.20.0 (2026-04-28)

Full Changelog: [v0.19.1...v0.20.0](https://github.com/sentdm/sent-dm-go/compare/v0.19.1...v0.20.0)

### Features

* **go:** add default http client with timeout ([2238e68](https://github.com/sentdm/sent-dm-go/commit/2238e6893398eeb5051783a6e240c29c812f267b))
* support setting headers via env ([490ce5e](https://github.com/sentdm/sent-dm-go/commit/490ce5e4023b39cfa2ba12ba02cec009193c7b87))

## 0.19.1 (2026-04-24)

Full Changelog: [v0.19.0...v0.19.1](https://github.com/sentdm/sent-dm-go/compare/v0.19.0...v0.19.1)

### Chores

* **internal:** more robust bootstrap script ([f65ebd5](https://github.com/sentdm/sent-dm-go/commit/f65ebd5d1a3a8d85979e169cb99467c674bdb1a6))

## 0.19.0 (2026-04-21)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/sentdm/sent-dm-go/compare/v0.18.0...v0.19.0)

### Features

* **api:** api update ([ec527fa](https://github.com/sentdm/sent-dm-go/commit/ec527fa032b50185b07540589680c44c18bee1c7))

## 0.18.0 (2026-04-20)

Full Changelog: [v0.17.1...v0.18.0](https://github.com/sentdm/sent-dm-go/compare/v0.17.1...v0.18.0)

### Features

* **api:** api update ([593668a](https://github.com/sentdm/sent-dm-go/commit/593668a68ec590d11a495e7ea59e2c2383bf8344))

## 0.17.1 (2026-04-10)

Full Changelog: [v0.17.0...v0.17.1](https://github.com/sentdm/sent-dm-go/compare/v0.17.0...v0.17.1)

### Bug Fixes

* better respect format tags from the spec ([96daf5d](https://github.com/sentdm/sent-dm-go/commit/96daf5d37bcea5bad2a7466afd8f495efc3c54d2))

## 0.17.0 (2026-04-07)

Full Changelog: [v0.16.1...v0.17.0](https://github.com/sentdm/sent-dm-go/compare/v0.16.1...v0.17.0)

### Features

* **api:** api update ([30971d3](https://github.com/sentdm/sent-dm-go/commit/30971d32454720aef8a73c25711b6b327a50b9cf))

## 0.16.1 (2026-04-03)

Full Changelog: [v0.16.0...v0.16.1](https://github.com/sentdm/sent-dm-go/compare/v0.16.0...v0.16.1)

### Bug Fixes

* fix issue with unmarshaling in some cases ([1ec2c95](https://github.com/sentdm/sent-dm-go/commit/1ec2c958bfb0cda0d2d8198dbbbf5c3de9620049))

## 0.16.0 (2026-03-31)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/sentdm/sent-dm-go/compare/v0.15.0...v0.16.0)

### Features

* **api:** manual updates ([527c9d2](https://github.com/sentdm/sent-dm-go/commit/527c9d225764d95d569ba569241426327b45af8e))

## 0.15.0 (2026-03-31)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/sentdm/sent-dm-go/compare/v0.14.0...v0.15.0)

### Features

* **internal:** support comma format in multipart form encoding ([31ab267](https://github.com/sentdm/sent-dm-go/commit/31ab267c9557a3d841f32467a5b54b418b05e607))


### Bug Fixes

* prevent duplicate ? in query params ([ff23ab0](https://github.com/sentdm/sent-dm-go/commit/ff23ab045caee21dd34532bb104e08d272255215))


### Chores

* **ci:** support opting out of skipping builds on metadata-only commits ([791f9ba](https://github.com/sentdm/sent-dm-go/commit/791f9bae36c571500651c8d043bee3ba0b573020))
* **client:** fix multipart serialisation of Default() fields ([affa363](https://github.com/sentdm/sent-dm-go/commit/affa363479d8daf2163106ad44cd789e9031db84))
* **internal:** support default value struct tag ([173c9e3](https://github.com/sentdm/sent-dm-go/commit/173c9e31857bf467fdddc03f5c336802935ecfc2))
* remove unnecessary error check for url parsing ([af61338](https://github.com/sentdm/sent-dm-go/commit/af613384ef47cd7f2423fe44a99f83018ff5737f))
* update docs for api:"required" ([c63067e](https://github.com/sentdm/sent-dm-go/commit/c63067ed8e482d483a460192f46ab7fd59f107c5))

## 0.14.0 (2026-03-25)

Full Changelog: [v0.13.1...v0.14.0](https://github.com/sentdm/sent-dm-go/compare/v0.13.1...v0.14.0)

### Features

* **api:** api update ([aabf001](https://github.com/sentdm/sent-dm-go/commit/aabf001530b2c77d36cf2371176cfbfbc97c9bb8))
* **api:** api update ([8b17379](https://github.com/sentdm/sent-dm-go/commit/8b1737970cd68ef554e8a722eeac6720ffad15d7))


### Chores

* **ci:** skip lint on metadata-only changes ([2bb5272](https://github.com/sentdm/sent-dm-go/commit/2bb527290da0338f9db1768b67fd1a2244d50f71))
* **internal:** update gitignore ([d670de7](https://github.com/sentdm/sent-dm-go/commit/d670de7fb5f184f7801c8f0552fb698b0b085e9a))

## 0.13.1 (2026-03-17)

Full Changelog: [v0.13.0...v0.13.1](https://github.com/sentdm/sent-dm-go/compare/v0.13.0...v0.13.1)

### Chores

* **internal:** tweak CI branches ([e9fc6ce](https://github.com/sentdm/sent-dm-go/commit/e9fc6ce70eed087349799728898ae36f90f8e1bb))

## 0.13.0 (2026-03-16)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/sentdm/sent-dm-go/compare/v0.12.0...v0.13.0)

### Features

* **api:** api update ([50aa024](https://github.com/sentdm/sent-dm-go/commit/50aa02426dbb1f2a8627ec1f695ea9dee55a997f))

## 0.12.0 (2026-03-12)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/sentdm/sent-dm-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** manual updates ([2604ce4](https://github.com/sentdm/sent-dm-go/commit/2604ce4df61fd602fb0b43c366fc44b63ea17052))

## 0.11.0 (2026-03-12)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/sentdm/sent-dm-go/compare/v0.10.0...v0.11.0)

### Features

* **api:** api update ([e536251](https://github.com/sentdm/sent-dm-go/commit/e536251e4a753b2c60a96f1494955772697113ff))
* **api:** manual updates ([5e2c19f](https://github.com/sentdm/sent-dm-go/commit/5e2c19fcbc41abf0ffc8f5bd2ef0e04873ae19ac))
* **api:** manual updates ([2c1aa55](https://github.com/sentdm/sent-dm-go/commit/2c1aa55d55abbab2f2c71d52de7b4be9b98d0364))

## 0.10.0 (2026-03-11)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/sentdm/sent-dm-go/compare/v0.9.0...v0.10.0)

### Features

* **api:** manual updates ([8ce333f](https://github.com/sentdm/sent-dm-go/commit/8ce333f5fae1e4a8605cbc37b043d8fd76574bb1))

## 0.9.0 (2026-03-11)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/sentdm/sent-dm-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** manual updates ([4e19ad9](https://github.com/sentdm/sent-dm-go/commit/4e19ad90bf4af16acda50bfcb8096cebb27b2334))

## 0.8.0 (2026-03-11)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/sentdm/sent-dm-go/compare/v0.7.1...v0.8.0)

### Features

* **api:** api update ([7fb43ac](https://github.com/sentdm/sent-dm-go/commit/7fb43acc7c6828df877957cfafbfb54cafb2ff68))

## 0.7.1 (2026-03-11)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/sentdm/sent-dm-go/compare/v0.7.0...v0.7.1)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([83a8163](https://github.com/sentdm/sent-dm-go/commit/83a81631e176b618770ea4dceb0d05aa65ebcbab))
* **internal:** codegen related update ([6f63ea8](https://github.com/sentdm/sent-dm-go/commit/6f63ea8fe011659f0605bc89dd1cba8da2493fc5))
* **internal:** codegen related update ([1375476](https://github.com/sentdm/sent-dm-go/commit/1375476e06cb90a49e1cf465993f8309b9c9b413))
* **internal:** minor cleanup ([ca9b532](https://github.com/sentdm/sent-dm-go/commit/ca9b53213f8724170635d07619bc37fd7c634b1e))
* **internal:** move custom custom `json` tags to `api` ([a208d39](https://github.com/sentdm/sent-dm-go/commit/a208d390304ab21a90998b8082d1eaf7f27d9372))
* **internal:** use explicit returns ([72af3b8](https://github.com/sentdm/sent-dm-go/commit/72af3b8d6a07bedef6725b68cccee54b124180dc))
* **internal:** use explicit returns in more places ([ee1e9e2](https://github.com/sentdm/sent-dm-go/commit/ee1e9e2745a0384271e568a8068382a5f8a2de10))

## 0.7.0 (2026-02-18)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/sentdm/sent-dm-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** manual updates ([33cd4ee](https://github.com/sentdm/sent-dm-go/commit/33cd4ee9efbd5462c93c27c296f0eab2a564a51a))
* **api:** manual updates ([57f856a](https://github.com/sentdm/sent-dm-go/commit/57f856ae119a9b57a27e0b407ed9bd8b03c3cb1f))
* **api:** manual updates ([d293a96](https://github.com/sentdm/sent-dm-go/commit/d293a9608ea8b172f299a74e75d17f407816009d))

## 0.6.0 (2026-02-16)

Full Changelog: [v0.5.1...v0.6.0](https://github.com/sentdm/sent-dm-go/compare/v0.5.1...v0.6.0)

### Features

* **api:** manual updates ([10d0e1f](https://github.com/sentdm/sent-dm-go/commit/10d0e1f165e55da6ec1f20192ac285eda4d4a3d8))

## 0.5.1 (2026-02-12)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/sentdm/sent-dm-go/compare/v0.5.0...v0.5.1)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([c527f7a](https://github.com/sentdm/sent-dm-go/commit/c527f7ad8426e2c56e6b1f944082cb4b98b1dbf9))

## 0.5.0 (2026-02-10)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/sentdm/sent-dm-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([1386117](https://github.com/sentdm/sent-dm-go/commit/1386117d2fde12190983969f919dcda184d67c22))

## 0.4.0 (2026-01-28)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/sentdm/sent-dm-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** manual updates ([04eabb9](https://github.com/sentdm/sent-dm-go/commit/04eabb9fb199243424fee896d3f4cbb95df7026d))

## 0.3.0 (2026-01-28)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/sentdm/sent-dm-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** manual updates ([14ddd5f](https://github.com/sentdm/sent-dm-go/commit/14ddd5f8c26122b8f8af9eb889247ab0af5338c0))

## 0.2.0 (2026-01-27)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/sentdm/sent-dm-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([2f905ed](https://github.com/sentdm/sent-dm-go/commit/2f905ed9f898bb3314c0bf202e28f968712c7bb0))

## 0.1.0 (2026-01-27)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/sentdm/sent-dm-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** api update ([0816c8d](https://github.com/sentdm/sent-dm-go/commit/0816c8d4582a84419cbb3be0a7931468743f62e4))
* **api:** api update ([8e2ef1b](https://github.com/sentdm/sent-dm-go/commit/8e2ef1b89a943ab0ac325668141b88b3b5e5b47e))
* **api:** api update ([212d55d](https://github.com/sentdm/sent-dm-go/commit/212d55d461806e15bb3e1551041d9ef3d3f1ff52))
* **api:** manual updates ([c444367](https://github.com/sentdm/sent-dm-go/commit/c444367795b119538539dbf69d0cc1e6ebc028e0))
* **api:** manual updates ([98c6921](https://github.com/sentdm/sent-dm-go/commit/98c69212bf92ee1b23a53a3934695eb0e8be4dc5))
* **api:** manual updates ([38facfd](https://github.com/sentdm/sent-dm-go/commit/38facfde356477bd2b50f49489d271f669c2b810))
* **api:** manual updates ([4c5d6a7](https://github.com/sentdm/sent-dm-go/commit/4c5d6a78ec1be6e0a32d47742ce9c7e195957666))
* **api:** manual updates ([cc1e5ee](https://github.com/sentdm/sent-dm-go/commit/cc1e5ee19f78a3af3b600c5a57f84fef442270bc))
* **api:** manual updates ([dd4ef3b](https://github.com/sentdm/sent-dm-go/commit/dd4ef3b8bbe468cd7ed840c4a2fdd61cc3a8e6e7))


### Chores

* sync repo ([4efe203](https://github.com/sentdm/sent-dm-go/commit/4efe2030975587260c4dcbfd018f2e580756e4ed))
* update SDK settings ([385a430](https://github.com/sentdm/sent-dm-go/commit/385a430b8e20c69aa28364bb1a2dc2433b9de54b))
* update SDK settings ([7c5068a](https://github.com/sentdm/sent-dm-go/commit/7c5068aa62fa67da9049af7d1591057980fe6442))
* update SDK settings ([d3488d9](https://github.com/sentdm/sent-dm-go/commit/d3488d9f100193f9f6893ed72a45e9ca88480442))
