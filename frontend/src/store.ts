import { computed, Ref, ref } from 'vue'
import { defineStore } from 'pinia'

import { ChooseFile, RetrieveFileContents } from '../wailsjs/go/main/App'
import { LogPrint } from '../wailsjs/runtime'

type Line = {
  LineNumber: number
  Text: string
  Type: string
}

type Content = {
  Lines: Line[]
}

export const useContentStore = defineStore('counter', () => {
  const currentFileName = ref('')
  const currentFileContent: Ref<Content> = ref({ Lines: [] })
  const currentKey: Ref<string> = ref('-')
  const errorMessage = ref('')
  const fileLoaded = ref(false)
  const loading = ref(false)

  const retrieveFile = async () => {
    currentKey.value = '-'
    const fileOpened = await ChooseFile()
    if (fileOpened == null || fileOpened.length === 0) {
      currentFileName.value = 'No file selected?'
      fileLoaded.value = false
    } else {
      loading.value = true
      currentFileName.value = fileOpened
      let content: any = null
      currentFileContent.value = { Lines: [] }
      try {
        content = await RetrieveFileContents(currentFileName.value)
        currentFileContent.value = content
        fileLoaded.value = true
      } catch (err: any) {
        errorMessage.value = err.toString()
        fileLoaded.value = false
        LogPrint(
          `error caught during file open: ${JSON.stringify(errorMessage.value, null, 2)}`
        )
      }

      loading.value = false
    }
  }

  const lineClass = computed(() => (lineNumber: number) => {
    let res = `flex space-x-2`

    switch (currentFileContent.value.Lines[lineNumber].Type) {
      case 'Section':
        res += ` bg-cyan-100`
        break
      case 'Chords':
        res += ` bg-pink-200`
        break
      case 'Lyrics':
        res += ` bg-yellow-100`
        break
    }

    return res
  })

  const keyChosen = computed(() => currentKey.value !== '-')

  return {
    currentFileName,
    currentFileContent,
    currentKey,
    errorMessage,
    fileLoaded,
    keyChosen,
    lineClass,
    loading,
    retrieveFile,
  }
})
