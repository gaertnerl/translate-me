import { jsonRequestJsonResponse } from "./http"
import { routes } from "./routes"

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
    currentSentence = await jsonRequestJsonResponse(routes.nextSentence.method, routes.nextSentence.route, {})
}

async function submitTranslation(): Promise<Feedback> {
    if (currentSentence === null) throw Error("there is no sentence to translate")
    const translation: Sentence = { id: currentSentence.id, text: currentUserTranslation }
    const feedback: Feedback = await jsonRequestJsonResponse(routes.submitSentence.method, routes.submitSentence.route, translation)
    score += feedback.score
    return feedback
}

export { score, currentSentence, currentUserTranslation, submitTranslation, loadNextSentence }