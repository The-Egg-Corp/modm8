import { Ref, ref } from "vue"

export interface Dialog extends DialogState {
    setVisible: (val?: boolean) => void
    setDraggable: (val?: boolean) => void
    setClosable: (val?: boolean) => void
}

export interface DialogState {
    visible: Ref<boolean>
    draggable: Ref<boolean>
    closable: Ref<boolean>
}

const dialogs: Record<string, DialogState> = {}

export const openDialogCount = () => Object.keys(dialogs).reduce((amt, key) => amt + (dialogs[key].visible.value ? 1 : 0), 0)
export const useDialog = (id: string): Dialog => {
    // Initialize the dialog state if it doesn't exist.
    if (!dialogs[id]) {
        dialogs[id] = {
            visible: ref(false),
            draggable: ref(false),
            closable: ref(false)
        }
    }

    const setVisible = (val = true) => {
        dialogs[id].visible.value = val
    }
    
    const setDraggable = (val = true) => {
        dialogs[id].draggable.value = val
    }
    
    const setClosable = (val = true) => {
        dialogs[id].closable.value = val
    }
    
    return {
        visible: dialogs[id].visible,
        draggable: dialogs[id].draggable,
        closable: dialogs[id].closable,
        setVisible,
        setDraggable,
        setClosable
    }
}