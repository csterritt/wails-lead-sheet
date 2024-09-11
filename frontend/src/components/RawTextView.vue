<template>
  <div v-if="store.currentFileContent == null" class="italic">
    No content yet (open a file)
  </div>

  <div v-else class="font-monoslab">
    <div v-for="line in store.currentFileContent.Lines" :key="line.LineNumber">
      <div :class="lineClass(line.LineNumber)">
        <span class="w-6">{{ line.LineNumber + 1 }}</span>
        <pre>{{ line.Text }}</pre>
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
