<script lang="ts" setup>

</script>

<template>
<!-- <v-layout>
    <v-navigation-drawer disable-resize-watcher floating permanent width="90">
        <v-list>
            <v-list-item class="item-content">
                <v-btn flat icon="mdi-view-dashboard"></v-btn>
            </v-list-item>
            <v-list-item class="item-content">
                <v-btn flat icon="mdi-cube"></v-btn>
            </v-list-item>
        </v-list>
        <v-list-item class="item-content bottom-pad">
            <v-btn flat icon="mdi-cog"></v-btn>
        </v-list-item>
    </v-navigation-drawer>
</v-layout> -->
</template>

<style scoped>
.v-navigation-drawer {
    background-color: rgba(255, 255, 255, 0.09);
    justify-items: center;
}

.item-content :deep(.v-list-item__content) {
    justify-self: center;
}

.bottom-pad {
    position: absolute;
    bottom: 15px;
    margin-left: auto;
    margin-right: auto;
    left: 0;
    right: 0;
}

.v-list-item .v-btn {
    color: rgba(223, 223, 223, 0.962);
    border-color: rgba(225, 225, 225, 0.421);
    border-width: 1px;
    border-radius: 8px;
    width: 55px;
    height: 50px;
}
</style>