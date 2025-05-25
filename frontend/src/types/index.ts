export declare type Nullable<T = void> = T | null | undefined
export declare type ObjValues<T extends Record<any, any>> = T[keyof T]

export declare  type Prettify<T> = {
  [K in keyof T]: T[K];
} & {};

export * from './settings'
export * from './primevue'
export * from './game'
export * from './thunderstore'
export * from './profile'