import type {
    Commit, Dispatch,
    ActionContext as Ctx
} from 'vuex'

export interface Context {
    dispatch: Dispatch
    commit: Commit
}

export interface ContextWithState<S> extends Ctx<S, any> {
    state: S
}

export type ActionContext<S, R> = Ctx<S, R>