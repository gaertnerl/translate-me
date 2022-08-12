import { elements } from "./lib/dom";
import { loadNextSentence, submitTranslation, score } from "./lib/game";

async function submitTranslationButtonClickHandler() {
    await submitTranslation()
    if (elements.scoreElement !== null) elements.scoreElement.innerHTML = String(score)
    await loadNextSentence()
}


loadNextSentence()
elements.submitTranslationButton?.addEventListener("click", submitTranslationButtonClickHandler)

const testjajaja = 10


