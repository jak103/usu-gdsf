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
}