export default class Game {
    constructor(
        Id = 1,
        Name = "Minecraft",
        Developer = "Mojang Studios",
        Rating = 3.5,
        TimesPlayed = 5,
        ImagePath = "../src/assets/minecraft.png",
        Description = 'Explore your own unique world, survive the night, and create anything you can imagine!',
        Tags = ['Survival', 'Sandbox'],
        CreationDate = new Date(),
        Version = "1.0.0",
        Downloads = 1234,
        DownloadLink = "TEST LINK"
    ){
        this.Id = Id;
        this.Name = Name;
        this.Developer = Developer;
        this.Rating = Rating;
        this.TimesPlayed = TimesPlayed;
        this.ImagePath = ImagePath;
        this.Description = Description;
        this.Tags = Tags;
        this.CreationDate = CreationDate;
        this.Version = Version;
        this.Downloads = Downloads;
        this.DownloadLink = DownloadLink;
    }

    static populateFromObject(gameObj) {
        return new Game(
            gameObj.Id ? gameObj.Id : "",
            gameObj.Name ? gameObj.Name : "",
            gameObj.Developer ? gameObj.Developer : "",
            gameObj.Rating ? gameObj.Rating : 0,
            gameObj.TimesPlayed ? gameObj.TimesPlayed : 0,
            gameObj.ImagePath ? gameObj.ImagePath : "",
            gameObj.Description ? gameObj.Description : "",
            gameObj.Tags ? gameObj.Tags : [],
            gameObj.CreationDate ? gameObj.CreationDate : new Date(),
            gameObj.Version ? gameObj.Version : "0.0.0",
            gameObj.Downloads ? gameObj.Downloads : 0,
            gameObj.DownloadLink ? gameObj.DownloadLink : ""
        )
    }
}