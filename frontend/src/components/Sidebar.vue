<script lang="ts" setup>

</script>

<template>
<v-layout>
    <v-navigation-drawer disable-resize-watcher floating permanent width="95">
    <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-view-dashboard"></v-list-item>
        <v-list-item prepend-icon="mdi-forum"></v-list-item>
        <v-list-item prepend-icon="mdi-cog"></v-list-item>
    </v-list>
    </v-navigation-drawer>
</v-layout>
</template>

<style scoped>
.v-navigation-drawer {
    border: 2px;
    border-color: white;
    border-radius: 5px;
    padding-top: 5px;
    background-color: rgba(42, 42, 42, 0.74);
}

.v-list-item {
    color: white;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}
</style>

