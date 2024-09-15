# modm8 <img align="right" width="128" height="128" src="./frontend/src/assets/images/appicon.png">

Lightweight and intuitive mod manager. Currently supports [Thunderstore](https://thunderstore.io/).\
Built with [Wails](https://wails.io) for an Electron-like experience but without the bloat.

## Installation
No installer or executable currently exists yet, please build from source instead. 
1. Clone this repository.
2. Install [Wails](https://wails.io/docs/gettingstarted/installation) and run one of the following:
   - `wails build` - Generates an executable in **/build/bin**.
   - `wails dev` - Opens a preview of the app with DevTools enabled.

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

  - Useful buttons and insightful info layed out in a sleek and compact manner.
  - Flexible query matching ensures you can search for a game without being exact.
  - Filter by your favourite or installed games - no more tedious searching.
  - Switch between List and Grid layouts for a better viewing experience.

  <img src="./screenshots/game-selection-grid.png"/>
</details>

<details>
  <summary>Selected Game Screen</summary>

  <img src="./screenshots/selected-game.png"/>
</details>

<details>
  <summary>Config Editor</summary>

  <img src="./screenshots/config-editor.png"/>
</details>

<details>
  <summary>Settings</summary>

  <img src="./screenshots/settings.png"/>
</details>

## Contributing
Pull requests and issues are very much welcome and are essential to keeping this project working and evolving!\
All I ask is you follow some simple guidelines to maintain a healthy, productive environment.

- Please be polite and keep critisicm constructive toward the project and other contributors.
- Provide as much detail as possible so the issue or idea can be easily understood.
- Bumping/rushing an issue does not help and only creates unnecessary notifications, please be patient.
- Promptly address any CI failures, a maintainer is not obliged to do this for you.

A lot of time and effort went into this project as a single person.\
If you find it useful or like where it's headed, please consider sponsoring it. 💛

## Contact
Feel free to join the [discord](https://discord.gg/psBXpXF2JZ) for support, suggestions or anything mod related!
