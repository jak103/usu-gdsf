import axios from "axios";

// Likely need to change this to be dynamic.
const root = "http://127.0.0.1:8080";

// GET
export function getAllGames() {
    return axios.get(`${root}/games`)
}

export function getGamesWithTags(tagsArr) {
    var tags = tagsArr.join('-')
    return axios.get(`${root}/game/tags`, { params: { tags }})
}

export function getGameInfo(id) {
    return axios.get(`${root}/info/${id}`);
}

export function getAllTags() {
    return axios.get(`${root}/games/tags`)
}

// POST

// PUSH

// DELETE