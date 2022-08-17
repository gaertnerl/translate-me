import { createScore0Element, createScore100Element, createScore500Element, elements } from "./lib/dom";
import { loadNextSentence, submitTranslation, score as gameScore, currentSentence, updateCurrentUserTranslation, Score, score } from "./lib/game";


function renderCurrentSentence() {
    const originalText = elements.originalText
    if (originalText === null) throw Error("cannot render current sentence, element is null")
    originalText.innerHTML = currentSentence.Text
}

function addScoreElement(element: HTMLElement) {
    const scoreList = elements.scoreList
    if (scoreList === null) throw Error("cannot update scores, element not found")
    const cnt = scoreList.childElementCount
    if (cnt >= 5) scoreList.removeChild(scoreList.lastChild)
    elements.scoreList.prepend(element)
}

function renderLastScore(score: Score) {
    let element
    if (score === 0) element = createScore0Element()
    if (score === 100) element = createScore100Element()
    if (score === 500) element = createScore500Element()
    addScoreElement(element)
}

function timeout(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function renderTotalScore(lastScore: number) {

    function calcCounterRate(i, total): Number {
        return 10 * (100 + Math.abs(i - (total / 2))) / total
    }

    const score = elements.totalScore
    if (score === null) throw Error("cannot render total score, element is null")
    if (!gameScore) score.innerHTML = "---"
    else {
        // should take 1000 ms in total
        const gameScoreCP = gameScore
        let scoreBefore = gameScoreCP - lastScore
        for (let i = scoreBefore; i <= gameScoreCP; i++) {
            const counterRate = calcCounterRate(i - scoreBefore, lastScore)
            await timeout(counterRate)
            score.innerHTML = String(i)
        }
    }
}


async function submitTranslationButtonClickHandler() {
    const score = await submitTranslation()
    renderLastScore(score.Score)
    renderTotalScore(score.Score)
    await loadNextSentence()
    renderCurrentSentence()
}

async function startGame() {

    const submitTranslationButton = elements.submitTranslationButton
    if (submitTranslationButton === null) throw Error("submit translation button not found")
    submitTranslationButton.addEventListener("click", submitTranslationButtonClickHandler)

    const userTranslationTextarea = elements.userTranslationTextarea
    if (userTranslationTextarea === null) throw Error("user translation textarea is null")
    userTranslationTextarea.addEventListener("input", function (event) {
        updateCurrentUserTranslation((<{ value: string }><unknown>event.target).value)
    })

    await loadNextSentence()
    renderCurrentSentence()
}

startGame()
