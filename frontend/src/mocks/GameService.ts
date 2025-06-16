import type { GameContainer, ThunderstoreGame, NexusGame } from "@types"

export function makeThunderstoreGame(game: ThunderstoreGame): GameContainer {
  return { platform: 'THUNDERSTORE', value: game }
}

export function makeNexusGame(game: NexusGame): GameContainer {
  return { platform: 'NEXUS_MODS', value: game }
}

export const mockGameList: GameContainer[] = [
  // makeNexusGame({
  //   ID: 1093894374,
  //   title: 'Lethal Company',
  //   domainName: 'lethal-company'
  // }),
  makeThunderstoreGame({
    title: "Lethal Company",
    identifier: 'lethal-company',
    image: "LethalCompany.png",
    path: "E:/SteamLibrary/steamapps/common/Lethal Company",
    aliases: ["LC", "LethalCompany"],
    steamID: 1966720
  }),
  makeThunderstoreGame({
    title: "Risk of Rain 2",
    identifier: 'riskofrain2',
    image: "RiskOfRain2.jpg",
    path: "H:\\Program Files (x86)\\Steam\\steamapps\\common\\Risk of Rain 2",
    aliases: ["ror2"],
    steamID: 248820
  }),
  makeThunderstoreGame({
    title: "Content Warning",
    identifier: 'content-warning',
    image: "ContentWarning.png",
    path: "H:\\Program Files (x86)\\Steam\\steamapps\\common\\Content Warning",
    aliases: ["CW", "ContentWarning"],
    steamID: 0
  }),
  makeThunderstoreGame({
    title: "Palworld",
    identifier: 'palworld',
    image: "Palworld.png",
    path: "H:\\Program Files (x86)\\Steam\\steamapps\\common\\Palworld",
    aliases: ["Palworld", "Pal"],
    steamID: 0
  }),
  makeThunderstoreGame({
    title: "Valheim",
    identifier: 'valheim',
    image: "Valheim.jpg",
    steamID: 0
  }),
  makeThunderstoreGame({
    title: "Subnautica",
    identifier: 'subnautica',
    image: 'Subnautica.png',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: "Dredge",
    identifier: 'dredge',
    image: 'Dredge.png',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: 'Techtonica',
    identifier: 'techtonica',
    image: 'techtonica.png',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: 'Placeholder',
    identifier: 'tf2',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: 'Placeholder',
    identifier: 'tf3',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: 'Placeholder',
    identifier: 'tf4',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: 'Placeholder',
    identifier: 'tf5',
    steamID: 0
  }),
  makeThunderstoreGame({
    title: 'Placeholder',
    identifier: 'tf6',
    steamID: 0
  })
]