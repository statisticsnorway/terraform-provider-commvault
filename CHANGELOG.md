# Changelog

## [0.1.0](https://github.com/statisticsnorway/terraform-provider-commvault/compare/v0.0.4...v0.1.0) (2025-09-03)


### Features

* support multiple bucket_contents for GCP subclient ([#18](https://github.com/statisticsnorway/terraform-provider-commvault/issues/18)) ([1be3929](https://github.com/statisticsnorway/terraform-provider-commvault/commit/1be39296d31290ea3af2952a48ebd2eefe881962))


### Bug Fixes

* **apiclient:** Add docs and update parameter name ([908d4be](https://github.com/statisticsnorway/terraform-provider-commvault/commit/908d4bed3c0b4699535d02e10e35c086cef06abb))
* **apiclient:** variable should be passed by pointer ([908d4be](https://github.com/statisticsnorway/terraform-provider-commvault/commit/908d4bed3c0b4699535d02e10e35c086cef06abb))
* Make project for bucket configuration required ([908d4be](https://github.com/statisticsnorway/terraform-provider-commvault/commit/908d4bed3c0b4699535d02e10e35c086cef06abb))
* Set subclient ID ([908d4be](https://github.com/statisticsnorway/terraform-provider-commvault/commit/908d4bed3c0b4699535d02e10e35c086cef06abb))

## [0.0.4](https://github.com/statisticsnorway/terraform-provider-commvault/compare/v0.0.3...v0.0.4) (2025-09-02)


### Bug Fixes

* Add information to logging ([7511ef8](https://github.com/statisticsnorway/terraform-provider-commvault/commit/7511ef864d1c108c5ab6319830093ab6c5e5c1b1))

## [0.0.3](https://github.com/statisticsnorway/terraform-provider-commvault/compare/v0.0.2...v0.0.3) (2025-09-02)


### Miscellaneous Chores

* release 0.0.3 ([91ffeb4](https://github.com/statisticsnorway/terraform-provider-commvault/commit/91ffeb4b650e72f2dad472ea38f1f52d01a9fbd7))

## [0.0.2](https://github.com/statisticsnorway/terraform-provider-commvault/compare/v0.0.1...v0.0.2) (2025-09-02)


### Miscellaneous Chores

* release 0.0.2 ([67aea7e](https://github.com/statisticsnorway/terraform-provider-commvault/commit/67aea7e0672edf42214997d5869e052f4dc35810))

## 0.0.1 (2025-09-01)


### Features

* Add workflow for release of provider ([fb3ce93](https://github.com/statisticsnorway/terraform-provider-commvault/commit/fb3ce93a646f364aa8e878f3bbd54ce4f2f0b464))
* **apiclient:** Add methods to delete client ([2856839](https://github.com/statisticsnorway/terraform-provider-commvault/commit/2856839f6160adbe61b6883aca12428e92ca8c0e))
* **apiclient:** Add type for Client get by id ([fd8352e](https://github.com/statisticsnorway/terraform-provider-commvault/commit/fd8352e89b7b86165f7a922b149ccf75b3ea461c))
* **apiclient:** Generate client for sp36 (v11.36) and refactor apiexplorer and sp36 to own modules ([032fbd5](https://github.com/statisticsnorway/terraform-provider-commvault/commit/032fbd594812b3f430a33e08076d45bc24c2280f))
* **apiclient:** Implement create subclient ([b9e4e37](https://github.com/statisticsnorway/terraform-provider-commvault/commit/b9e4e3796fb4bdcfb839e24f4caa213c4869d5f2))
* **apiclient:** Implement endpoints for login, client and subclient ([8bece60](https://github.com/statisticsnorway/terraform-provider-commvault/commit/8bece60a03588512a0158800e53772a941505d21))
* **apiclient:** Implement login and make client accessible as providerdata ([2176d02](https://github.com/statisticsnorway/terraform-provider-commvault/commit/2176d02d542546b1e17e3ceb0edc2c362023a69c))
* Increase timeout ([d8b1c1c](https://github.com/statisticsnorway/terraform-provider-commvault/commit/d8b1c1caf230a381b5d793642c9d034023423ab0))
* Use new apiclient ([33c1906](https://github.com/statisticsnorway/terraform-provider-commvault/commit/33c1906b8ef702223f436e89424aa09fde29254c))


### Bug Fixes

* **apiclient:** Delete method should indeed use http method delete ([a7599f2](https://github.com/statisticsnorway/terraform-provider-commvault/commit/a7599f2813b01f9e7c1a45925900cc46f091d889))
* **apiclient:** Fix build after generation ([dca30ce](https://github.com/statisticsnorway/terraform-provider-commvault/commit/dca30ced96512663209ce1720a140f01d80d5246))
* **apiclient:** Fix build after generation ([a7715c1](https://github.com/statisticsnorway/terraform-provider-commvault/commit/a7715c1431ffa1f70db0ca9ffcea52dc2882b983))
* **apiclient:** login model should have defined json names for username and password ([f3e76b0](https://github.com/statisticsnorway/terraform-provider-commvault/commit/f3e76b071bb4b655df58188cab6e6eab4a118319))


### Miscellaneous Chores

* release 0.0.1 ([6da4de1](https://github.com/statisticsnorway/terraform-provider-commvault/commit/6da4de1eef39c06d511bfa763a38a2b765ba803e))
