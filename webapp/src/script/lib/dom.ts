const userTranslationId = "user-translation"
const originalText = "original-text"
const submitTranslationButtonId = "submit-translation-button"
const scoreId = "score"

function getSubmitTranslationButton(): HTMLElement | null {
    return document.getElementById(submitTranslationButtonId)
}

function getOriginalText(): HTMLElement | null {
    return document.getElementById(originalText)
}

function getUserTranslationTextarea(): HTMLElement | null {
    return document.getElementById(userTranslationId)
}

function getScoreElement(): HTMLElement | null {
    return document.getElementById(scoreId)
}

const elements: { [key: string]: HTMLElement | null } = {
    originalText: getOriginalText(),
    submitTranslationButton: getSubmitTranslationButton(),
    userTranslationTextarea: getUserTranslationTextarea(),
    scoreElement: getScoreElement()
}

export { elements }