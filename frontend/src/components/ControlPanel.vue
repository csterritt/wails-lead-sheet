<template>
  <div>
    <p>Chosen file: {{ fileName }}</p>

    <button class="btn btn-primary" @click="openFile">Open a file</button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

import { ChooseFile } from '../../wailsjs/go/main/App'

const fileName = ref('')

const openFile = async () => {
  const fileOpened = await ChooseFile()
  if (fileOpened == null || fileOpened.length === 0) {
    fileName.value = 'No file selected?'
  } else {
    fileName.value = JSON.stringify(fileOpened)
  }
}
</script>
