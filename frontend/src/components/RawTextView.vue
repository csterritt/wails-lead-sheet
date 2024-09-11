<template>
  <div v-if="store.currentFileContent == null" class="italic">
    No content yet (open a file)
  </div>

  <div v-else>
    <div v-for="line in store.currentFileContent.Lines" :key="line.LineNumber">
      <div :class="lineClass(line.LineNumber)">
        <span class="w-3">{{ line.LineNumber + 1 }}</span>
        <span class="w-3">{{ line.Type.substring(0, 1) }}</span>
        <span>{{ line.Text }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'

import { useContentStore } from '../store'

const store = useContentStore()

const { lineClass } = storeToRefs(store)
</script>
