<script lang="ts" setup>

</script>

<template>
<v-layout>
    <v-navigation-drawer disable-resize-watcher floating permanent width="95">
    <v-list density="compact">
        <v-list-item>
            <v-btn flat flex icon="mdi-view-dashboard"></v-btn>
        </v-list-item>
        <v-list-item>
            <v-btn flat icon="mdi-palette"></v-btn>
        </v-list-item>
        <v-list-item>
            <v-btn flat icon="mdi-cog"></v-btn>
        </v-list-item>
    </v-list>
    </v-navigation-drawer>
</v-layout>
</template>

<style scoped>
.v-navigation-drawer {
    background-color: rgba(255, 255, 255, 0.09);
}

.v-list-item {
    background-color: rgb(39, 41, 50);
}

.v-list-item .v-btn {
    color: rgba(223, 223, 223, 0.962);
    display: inline-flex;
    border-radius: 8px;
    border-color: rgba(173, 58, 255, 0.421);
    border-width: 0.1px;
    width: 50px;
    height: 45px;
}
</style>

