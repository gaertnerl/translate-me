const userTranslationId = "user-translation"
const originalText = "original-text"
const submitTranslationButtonId = "submit-translation-button"
const scoreClass = "score"
const score500Class = "score-500"
const score100Class = "score-100"
const score0Class = "score-0"
const scoreListId = "score-list"
const totalScoreId = "total-score-num"


function createScore500Element(): HTMLElement | null {
    const div = document.createElement('div');
    div.classList.add(scoreClass);
    div.classList.add(score500Class);
    const text = document.createTextNode("Perfect! 🤩");
    div.appendChild(text);
    return div
}

function createScore100Element(): HTMLElement | null {
    const div = document.createElement('div');
    div.classList.add(scoreClass);
    div.classList.add(score100Class);
    const text = document.createTextNode("Good 😊");
    div.appendChild(text);
    return div
}

function createScore0Element(): HTMLElement | null {
    const div = document.createElement('div');
    div.classList.add(scoreClass);
    div.classList.add(score0Class);
    const text = document.createTextNode("Wrong 🙃");
    div.appendChild(text);
    return div
}

function getScoreListElement(): HTMLElement | null {
    return document.getElementById(scoreListId)
}

function getSubmitTranslationButton(): HTMLElement | null {
    return document.getElementById(submitTranslationButtonId)
}

function getOriginalText(): HTMLElement | null {
    return document.getElementById(originalText)
}

function getUserTranslationTextarea(): HTMLElement | null {
    return document.getElementById(userTranslationId)
}

function getTotalScoreElement(): HTMLElement | null {
    return document.getElementById(totalScoreId)
}

const elements: { [key: string]: (HTMLElement | null) } = {
    originalText: getOriginalText(),
    submitTranslationButton: getSubmitTranslationButton(),
    userTranslationTextarea: getUserTranslationTextarea(),
    scoreList: getScoreListElement(),
    totalScore: getTotalScoreElement()
}

export { elements, createScore500Element, createScore100Element, createScore0Element }