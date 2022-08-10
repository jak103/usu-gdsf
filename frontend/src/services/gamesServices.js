import axios from "axios";
import Game from "../models/game"

// Likely need to change this to be dynamic.
const root = "http://127.0.0.1:8080";

// GET
export function getAllGames() {
    return new Promise((resolve, reject) => {
        axios.get(`${root}/games`)
        .then(resp => {
            resolve({ data: resp.data.map(g => Game.populateFromObject(g)) })
        })
        .catch(error => {
            console.log(error)
            reject(error)
        })
    })
}

export function getGamesWithTags(tagsArr) {
    var tags = tagsArr.join('-')
    return new Promise((resolve, reject) => {
        axios.get(`${root}/game/tags`, { params: { tags }})
        .then(resp => {
            resolve({ data: resp.data.map(g => Game.populateFromObject(g)) })
        })
        .catch(error => {
            console.log(error)
            reject(error)
        })
    });
}

export function getGameInfo(id) {
    return new Promise((resolve, reject) => {
        axios.get(`${root}/game/${id}`)
            .then(resp => 
                resolve({ data: Game.populateFromObject(resp.data)}))
            .catch(error => {
                console.log(error)
                reject(error);
            })
    });
}

export function getAllTags() {
    return axios.get(`${root}/games/tags`)
}

export function getMostPopularGame() {
    return new Promise((resolve, reject) => {
        axios.get(`${root}/most_popular`)
            .then(resp => resolve({ data: Game.populateFromObject(resp.data) }))
            .catch(error => {
                console.log(error)
                reject(error);
            })
    })
}

// POST

// PUSH

// DELETE