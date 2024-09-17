<template>
  <div
    v-if="fileContent == null || fileContent.Lines.length === 0"
    class="italic"
  >
    {{ message }}
  </div>

  <div v-else class="font-monoslab">
    <div v-for="line in fileContent.Lines" :key="line.LineNumber">
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

const props = defineProps({
  fileContent: { type: Object, required: true },
  message: { type: String, required: true },
})

const store = useContentStore()

const { lineClass } = storeToRefs(store)
</script>
