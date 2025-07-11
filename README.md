# modm8 <img align="right" width="128" height="128" src="./frontend/src/assets/images/appicon.png">
A simple yet powerful mod manager. Fast ✅ Intuitive ✅ Ad-free ✅ Multiple languages ✅\
Built with [Wails](https://wails.io) for an Electron-like experience but without the bloat.

> [!WARNING]
> **CURRENT STATUS:** Not yet ready.\
> This project is heavily **WIP** and as such, missing functionality or bugs will be commonplace until there is a stable release (v1).

## Features
- Multiple themes + dark mode by default.
- Go as the backend for concurrency and responsiveness.
- No browser bundled, uses native browser to render content.
- App settings can be easily changed. No need to go through unnecessary menus.
- Familiar frontend. Similar to existing apps like `r2modman`.
- Supports multiple languages with the help of `vue-i18n`.
  - Currently implemented: **English**, **German**, **French**, **Spanish**, **Italian**
  - Language PRs are greatly appreciated as some keys may have been wrongly translated!

## Showcase
<details>
  <summary>Game Selection</summary>
  
  <img src="./screenshots/game-selection-grid.png"/>
</details>

<details>
  <summary>Selected Game Screen</summary>

  <img src="./screenshots/selected-game.png"/>
  <img src="./screenshots/this-profile.png"/>
</details>

<details>
  <summary>Config Editor</summary>

  <img src="./screenshots/config-editor-selection.png"/>
  <img src="./screenshots/config-editor.png"/>
</details>

<details>
  <summary>Settings</summary>

  <img src="./screenshots/settings.png"/>
</details>

## Installing
There are currently no official releases. You can download dev builds under the latest [Crossplatform Build](https://github.com/The-Egg-Corp/modm8/actions/workflows/build.yml) action.

## Building from source
1. Clone this repository.
2. Install dependencies ([Node 20+](https://nodejs.org/en/download), [Bun](https://bun.sh), [Go](https://go.dev/doc/install), [Wails](https://wails.io/docs/gettingstarted/installation)).
3. Run one of the following commands:
   - `wails build` - Generates an executable in **/build/bin**.
   - `wails dev` - Opens a preview of the app with DevTools enabled.

## Contributing
Pull requests and issues are very much welcome and are essential to keeping this project working and evolving!\
All I ask is you follow a few simple guidelines to maintain a respectful and productive environment.

- Please be polite and keep critisicm toward the project and other contributors constructive.
- Provide as much detail as possible so the issue or idea can be easily understood.
- Bumping/rushing an issue does not help and only creates unnecessary notifications, please be patient.

To get started, see the **Building from source** section above.

## Contact
Feel free to join the [discord](https://discord.gg/psBXpXF2JZ) to contact me! Here you will get to:

- See first-hand updates about the project and its progress.
- Chat about the project's development or anything mod related.
- Suggest ideas and get support if you are stuck.
- View the **TODO** list to see what needs fixing/implementing and if it's being currently worked on.
