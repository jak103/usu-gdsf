export default class Game {
    constructor(
        id = 1,
        name = "Minecraft",
        developer = "Mojang Studios",
        rating = 3.5,
        timesPlayed = 5,
        imagePath = "../src/assets/minecraft.png",
        description = 'Explore your own unique world, survive the night, and create anything you can imagine!',
        tags = ['Survival', 'Sandbox']
    ){
        this.id = id;
        this.name = name;
        this.developer = developer;
        this.rating = rating;
        this.timesPlayed = timesPlayed;
        this.imagePath = imagePath;
        this.description = description;
        this.tags = tags;
    }
}