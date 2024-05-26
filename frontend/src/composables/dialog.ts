import { ref } from "vue"

const visible = ref(false)
const draggable = ref(false)
const closable = ref(false)

const setVisible = (val: boolean) => {
    visible.value = val
}

const setDraggable = (val: boolean) => {
    draggable.value = val
}

const setClosable = (val: boolean ) => {
    closable.value = val
}

export const useDialog = () => ({
    visible,
    draggable,
    closable,
    setVisible,
    setDraggable,
    setClosable
})