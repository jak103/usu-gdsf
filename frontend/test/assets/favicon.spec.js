import fs from 'fs'
var base = process.cwd
const path = base + "../../src/assets/favicon.ico"

//var fav = new File("../../src/assets/favicon.ico")
console.log()

describe('favicon available', () => {
    const output = fs.existsSync(path);
    test('file exists', () => {
        expect(output).toBe(true);
    });
})
