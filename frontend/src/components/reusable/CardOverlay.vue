<script lang="ts" setup>
interface Props {
    dialogContentStyle?: any
    dialogStyle?: any
    cardStyle?: any
    class?: string
}

const props = withDefaults(defineProps<Props>(), {
    dialogStyle: {
        "margin-left": "40px",
        "width": 'auto',
        "min-width": '40rem',
        "max-width": '65rem',
        "box-shadow": '0px 0px 15px rgb(0, 0, 0, 0.9)'
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
    :content-style="props.dialogContentStyle"
    :block-scroll="true"
    :dismissable-mask="false"
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
    margin-top: 10px;
    margin-bottom: 10px;
    -webkit-font-smoothing: antialiased;
    text-wrap: nowrap;
}
</style>