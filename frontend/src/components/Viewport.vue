<script lang="ts" setup>

</script>

<template>
<div class="viewport">
    <div class="content-area"></div>
</div>
</template>
  
<style scoped>
.viewport {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.content-area {
    display: flex;
    width: 100vw; /* Full viewport width */
    height: 100vh; /* Full viewport height to center content */
    justify-content: center; /* Center horizontally */
    align-items: flex-start; /* Center vertically */
    padding-top: 30px;
    padding-right: 65px;
    margin-left: 75px;
    border: rgba(216, 216, 216, 0.912) solid 1px;
    border-radius: 8px;
}
</style>