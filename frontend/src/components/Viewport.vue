<script lang="ts" setup>

</script>

<template>
<div class="viewport">
    <div class="content-area">
        <h1>Settings</h1>
    </div>
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
    margin-top: 35px;
    margin-left: 150px;
    background-color: rgba(0, 0, 0, 0);
    border: rgba(216, 216, 216, 0.912) solid 1px;
    border-radius: 8px;
}

.content-area h1 {
    margin: 0;
    font-size: 42px;
}
</style>