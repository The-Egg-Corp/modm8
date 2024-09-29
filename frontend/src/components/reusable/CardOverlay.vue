<script lang="ts" setup>
interface Props {
    dialogStyle?: any
    cardStyle?: any
    class?: string
}

const props = withDefaults(defineProps<Props>(), {
    dialogStyle: {
        "margin-left": "10px",
        "width": 'auto',
        "min-width": '45rem',
        "max-width": '65rem'
    },
    cardStyle: {
        "box-shadow": 'none',
        "overflow": 'hidden',
        "margin": '-15px',
        "background": 'none'
    }
})
</script>

<template>
<Dialog modal
    :class="props.class"
    :style="props.dialogStyle"
    :block-scroll="true"
    :dismissable-mask="true"
    :show-header="false"
>
    <Card :style="props.cardStyle">
        <template #content>
            <div class="card-content">
                <slot name="cardContent"></slot>
            </div>
        </template>
    </Card>

    <slot name="dialogContent"></slot>
</Dialog>
</template>

<style scoped>
.card-content {
    justify-content: center;
}

:deep(.card-content .header) {
    user-select: none;
    font-size: 42px;
    font-weight: 490;
    margin-top: 15px;
    margin-bottom: 5px;
    -webkit-font-smoothing: antialiased;
}
</style>