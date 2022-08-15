import { jsonRequestJsonResponse } from "./http"
import { routes } from "./routes"

type Score = 0 | 100 | 500

type Sentence = {
    Text: string
    Id: number
}

type Feedback = {
    Score: Score
}

let score: number = 0
let currentSentence: Sentence | null = null
let currentUserTranslation: string = ""

function updateCurrentUserTranslation(s: string) {
    currentUserTranslation = s
}

async function loadNextSentence(): Promise<void> {
    currentSentence = await jsonRequestJsonResponse(routes.nextSentence.method, routes.nextSentence.route, {})
}

async function submitTranslation(): Promise<Feedback> {
    if (currentSentence === null) throw Error("there is no sentence to translate")
    const translation: Sentence = { Id: currentSentence.Id, Text: currentUserTranslation }
    const feedback: Feedback = await jsonRequestJsonResponse(routes.submitSentence.method, routes.submitSentence.route, translation)
    score += feedback.Score
    return feedback
}

export { score, currentSentence, currentUserTranslation, submitTranslation, loadNextSentence, updateCurrentUserTranslation, Score }