import { computed, Ref, ref } from 'vue'
import { defineStore } from 'pinia'

import {
  ChooseFile,
  ExportToClipboard,
  RetrieveFileContents,
  TransposeDownOneStep,
  TransposeUpOneStep,
} from '../wailsjs/go/main/App'
import { LogPrint } from '../wailsjs/runtime'
import { parser } from '../wailsjs/go/models'

type LetterRun = {
  Type: string
  Letters: string
  OriginalLetters: string
  Chord: string
  TransposedLetters: string
}

type Line = {
  LineNumber: number
  Text: string
  Parts: LetterRun[]
  Type: string
}

const processTransposedLines = (
  inputContent: parser.ParsedContent
): { Lines: Line[] } => {
  const content = inputContent as { Lines: Line[] }
  const res: { Lines: Line[] } = { Lines: [] }
  for (let lineIndex = 0; lineIndex < content.Lines.length; lineIndex += 1) {
    const line = content.Lines[lineIndex]
    if (line.Type === 'Chords') {
      const newLine = {
        LineNumber: line.LineNumber,
        Text: '',
        Parts: line.Parts,
        Type: line.Type,
      }

      for (let partIndex = 0; partIndex < line.Parts.length; partIndex += 1) {
        const part = line.Parts[partIndex]
        if (part.Type === 'ChordRun' && part.TransposedLetters !== '') {
          newLine.Text += part.TransposedLetters
        } else {
          newLine.Text += part.Letters
        }

        res.Lines[lineIndex] = newLine
      }
    } else {
      res.Lines[lineIndex] = line
    }
  }

  return res
}

export const useContentStore = defineStore('counter', () => {
  const currentFileName = ref('')
  const currentFileContent: Ref<parser.ParsedContent> = ref({ Lines: [] })
  const processedFileContent: Ref<parser.ParsedContent> = ref({ Lines: [] })
  const currentKey: Ref<string> = ref('-')
  const errorMessage = ref('')
  const fileLoaded = ref(false)
  const loading = ref(false)

  const lineClass = computed(() => (lineNumber: number) => {
    let res = `flex space-x-2`

    switch (
      currentFileContent.value != null &&
      (currentFileContent.value as { Lines: Line[] }).Lines[lineNumber].Type
    ) {
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

  const retrieveFile = async () => {
    currentKey.value = '-'
    const fileOpened = await ChooseFile()
    if (fileOpened == null || fileOpened.length === 0) {
      currentFileName.value = 'No file selected?'
      fileLoaded.value = false
    } else {
      loading.value = true
      currentFileName.value = fileOpened
      let content: parser.ParsedContent = parser.ParsedContent.createFrom({
        Lines: [],
      })
      currentFileContent.value = { Lines: [] }
      processedFileContent.value = { Lines: [] }
      try {
        content = await RetrieveFileContents(currentFileName.value)
        currentFileContent.value = content
        processedFileContent.value = content
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

  const transposeUp = async () => {
    const res = await TransposeUpOneStep(processedFileContent.value)
    processedFileContent.value = processTransposedLines(res)
  }

  const transposeDown = async () => {
    const res = await TransposeDownOneStep(processedFileContent.value)
    processedFileContent.value = processTransposedLines(res)
  }

  const exportToClipboard = async () => {
    const err = await ExportToClipboard(processedFileContent.value)
    if (err != '') {
      LogPrint(
        `error caught during export to clipboard: ${JSON.stringify(err, null, 2)}`
      )
      errorMessage.value = err
    }
  }

  return {
    currentFileName,
    currentFileContent,
    currentKey,
    errorMessage,
    exportToClipboard,
    fileLoaded,
    keyChosen,
    lineClass,
    loading,
    processedFileContent,
    retrieveFile,
    transposeDown,
    transposeUp,
  }
})
