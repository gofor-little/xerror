# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v1.0.0 - 2021-10-13
### Added
* Added CodeQL action.

## v0.5.1 - 2021-09-04
### Fixed
* Fixed character scape issues with ```Error.MarshalJSON```.

## v0.5.0 - 2021-08-21
### Added
* Added Go 1.17 support.
* Added ```Wrapf``` function.

## v0.4.1 - 2021-08-07
### Fixed
* Fixed a string escaping bug in ```Error.MarshalJSON```.

## v0.4.0 - 2021-06-28
### Added
* Added ```Wrap``` function.

### Changed
* **BREAKING**: Changed ```New``` function parameters.

## v0.3.0 - 2021-05-24
### Added
* Implemented the ```errors.Wrap``` interface on the Error struct.
* Implemented the ```error.Unwrap``` interface on the Error struct.

## v0.2.3 - 2021-05-20
### Fixed
* Fixed a failing test.

## v0.2.2 - 2021-05-20
### Changed
* Implemented the ```json.Marshaler``` interface on the Error struct.

## v0.2.1 - 2021-05-20
### Changed
* Updated Dependabot configuration file.

## v0.2.0 - 2021-05-17
### Added
* Added ```Newf``` function.

## v0.1.1 - 2021-02-20
### Changed
* Updated the readme.

## v0.1.0 - 2021-02-20
### Added
* Added the ```Error``` struct.
* Implemented the ```errors.Error``` interface on the Error struct.
* Added ```New``` function.
* Added GitHub CI action.
* Added GitHub Stale action.
* Added a license.
* Added a readme.
