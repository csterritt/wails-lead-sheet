<template>
  <div class="flex flex-grow flex-row relative overflow-auto">
    <div class="flex-grow w-1/2">
      <h3 class="text-2xl flex-grow">Raw text</h3>

      <raw-text-view
        :file-content="store.currentFileContent"
        message="No content yet (open a file)"
      ></raw-text-view>
    </div>

    <div class="divider divider-horizontal"></div>

    <div class="flex-grow w-1/2">
      <h3 class="text-2xl">Processed</h3>

      <raw-text-view
        :file-content="store.processedFileContent"
        message=""
      ></raw-text-view>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useContentStore } from '../store'

import RawTextView from './RawTextView.vue'

const store = useContentStore()
</script>
