import { ref } from 'vue'
import { defineStore } from 'pinia'

import { ChooseFile, RetrieveFileContents } from '../wailsjs/go/main/App'
import { LogPrint } from '../wailsjs/runtime'

export const useContentStore = defineStore('counter', () => {
  const currentFileName = ref('')
  const currentFileContent = ref('')
  const errorMessage = ref('')
  const loading = ref(false)

  const retrieveFile = async () => {
    const fileOpened = await ChooseFile()
    if (fileOpened == null || fileOpened.length === 0) {
      currentFileName.value = 'No file selected?'
    } else {
      loading.value = true
      currentFileName.value = fileOpened
      let content = ''
      let error = ''
      currentFileContent.value = ''
      try {
        content = await RetrieveFileContents(currentFileName.value)
        currentFileContent.value = content
      } catch (err: any) {
        error = err.toString()
        LogPrint(
          `error caught during file open: ${JSON.stringify(error, null, 2)}`
        )
      }

      loading.value = false
    }
  }

  return {
    currentFileName,
    currentFileContent,
    errorMessage,
    loading,
    retrieveFile,
  }
})
