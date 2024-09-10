<template>
  <div>
    <div v-if="store.loading" class="italic">Loading...</div>

    <div v-else>
      <p v-if="store.errorMessage.trim().length === 0">
        Chosen file: {{ store.currentFileName }}
      </p>

      <p v-else>Error: {{ store.errorMessage }}</p>
    </div>

    <button class="btn btn-primary" @click="store.retrieveFile">
      Open a file
    </button>
  </div>
</template>

<script setup lang="ts">
import { useContentStore } from '../store'

const store = useContentStore()
</script>
