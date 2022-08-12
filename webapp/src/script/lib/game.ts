import { jsonRequestJsonResponse } from "./http"

type Sentence = {
    text: string
    id: number
}

type Feedback = {
    score: number
}

let score: number = 0
let currentSentence: Sentence | null = null
let currentUserTranslation: string = ""

async function loadNextSentence(): Promise<void> {
    currentSentence = await jsonRequestJsonResponse("POST", "/text/next", {})
}

async function submitTranslation(): Promise<Feedback> {
    if (currentSentence === null) throw Error("there is no sentence to translate")
    const translation: Sentence = { id: currentSentence.id, text: currentUserTranslation }
    const feedback: Feedback = await jsonRequestJsonResponse("POST", "/sentence/translate", translation)
    score += feedback.score
    return feedback
}

export { score, currentSentence, currentUserTranslation, submitTranslation, loadNextSentence }