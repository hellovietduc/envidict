const searchInput = document.getElementById('searchInput')
const searchBtn = document.getElementById('searchBtn')
const notiDiv = document.getElementById('noti')
const suggestionsDiv = document.getElementById('suggestions')
const resultsDiv = document.getElementById('results')
const serverUrl = 'http://localhost:8080'
const debounceTimeout = 200

async function search(e, input) {
    clearNoti()
    try {
        let searchTerm = input || searchInput.value
        searchTerm = searchTerm.trim()
        if (!searchTerm) {
            return showNoti('Please enter a search term.')
        }

        let url = `${serverUrl}/search?t=${searchTerm}`
        if (input) url += `&e=true`

        const res = await fetch(url)
        if (!res.ok) {
            return showNoti("Not found. Please try again later.")
        }

        const json = await res.json()
        if (!json.match) {
            showSuggestions(json.related)
            clearResult()
            clearRelated()
            return
        }

        showResult(json.match)
        if (json.related) {
            showRelated(json.related, searchTerm)
        }

        clearSuggestions()
    } catch (err) {
        console.error(err)
        showNoti("Something's wrong. Please try again later.")
    }
}

function showNoti(msg) {
    notiDiv.innerHTML = `<h3>${msg}</h3>`
}

function clearNoti() {
    notiDiv.innerHTML = ''
}

function showResult(item) {
    resultsDiv.innerHTML = `
        <h2 class="word">${item.word}</h2>
        <div class="word-pronunciation">${item.pronunciation}</div>
        ${item.definitions.map(def => `
            <div class="word-definition">
                ${def.kind ? `<div class="word-kind">${def.kind}</div>` : ''}
                ${def.descriptions.map(desc => `
                    <div class="word-description">
                        ${desc.meaning ? `<div class="word-meaning">${desc.meaning}</div>` : ''}
                        ${desc.example ? `<div class="word-example">Example: ${desc.example}</div>` : ''}
                    </div>
            `).join('')}
            </div>
        `).join('')}
    `
}

function clearResult() {
    resultsDiv.innerHTML = ''
}

function showRelated(items, skipped) {
    resultsDiv.innerHTML += `
        <div class="related-words">
            <h3>Related words:</h3>
            ${items.map(e => `
                ${e.word === skipped ? '' : `<span class="related-word" onclick="search(null, '${e.word}')">${e.word}</span>`}
            `).join('')}
        </div>
    `
}

function clearRelated() {
    resultsDiv.innerHTML = ''
}

function showSuggestions(items) {
    suggestionsDiv.innerHTML = `
        ${items.map(e => `
            <span class="suggested-word" onclick="search(null, '${e.word}')">${e.word}</span>
        `).join('')}
    `
}

function clearSuggestions() {
    suggestionsDiv.innerHTML = ''
}

let timeout
searchInput.addEventListener('keydown', () => {
    clearTimeout(timeout)
    timeout = setTimeout(search, debounceTimeout)
})
searchBtn.addEventListener('click', search)
