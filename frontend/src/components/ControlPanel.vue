<template>
  <div class="mb-2">
    <h3 class="text-2xl">Lead Sheet App</h3>

    <div v-if="store.loading" class="italic">Loading...</div>

    <div v-else>
      <p v-if="store.errorMessage.trim().length === 0">
        Chosen file: {{ store.currentFileName }}
      </p>

      <p v-else>Error: {{ store.errorMessage }}</p>
    </div>

    <div class="flex flex-row items-center space-x-8">
      <button class="btn btn-primary" @click="store.retrieveFile">
        Open a file
      </button>

      <template v-if="store.fileLoaded">
        <div class="flex flex-row items-center space-x-2 text-xl">
          <span class="font-bold">Key:</span>
          <select
            class="select select-primary w-full max-w-xs"
            v-model="store.currentKey"
          >
            <option disabled>-</option>
            <option value="Ab">Ab</option>
            <option value="A">A</option>
            <option value="A#">A#</option>

            <option value="Bb">Bb</option>
            <option value="B">B</option>

            <option value="C">C</option>
            <option value="C#">C#</option>

            <option value="Db">Db</option>
            <option value="D">D</option>
            <option value="D#">D#</option>

            <option value="Eb">Eb</option>
            <option value="E">E</option>

            <option value="F">F</option>
            <option value="F#">F#</option>

            <option value="Gb">Gb</option>
            <option value="G">G</option>
            <option value="G#">G#</option>
          </select>
        </div>

        <div class="flex flex-row items-center space-x-2 text-xl">
          <span class="font-bold">Transpose:</span>
          <button class="btn btn-sm btn-primary" @click="store.transposeUp">
            Up
          </button>

          <button class="btn btn-sm btn-primary" @click="store.transposeDown">
            Down
          </button>
        </div>

        <button class="btn btn-sm btn-primary" :disabled="!store.keyChosen">
          Change to NNS
        </button>

        <button class="btn btn-sm btn-primary" @click="store.exportToClipboard">
          Export to clipboard
        </button>
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useContentStore } from '../store'

const store = useContentStore()
</script>
