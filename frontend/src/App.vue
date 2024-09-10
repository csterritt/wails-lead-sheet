<template>
  <div class="h-full flex flex-col flex-grow">
    <h3 class="text-2xl">Lead Sheet App</h3>

    <control-panel></control-panel>

    <file-view></file-view>
  </div>
</template>

<script lang="ts" setup>
import ControlPanel from './components/ControlPanel.vue'
import FileView from './components/FileView.vue'
</script>
