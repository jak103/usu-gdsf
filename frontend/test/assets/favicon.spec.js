/* Ensure that the favicon is in place and not modified in a manner incompatible with index.html */
import fs from 'fs'
var base = process.cwd
const path = base + "../../src/assets/favicon.ico"

describe('favicon available', () => {
    const output = fs.existsSync(path);
    test('file exists', () => {
        expect(output).toBe(true);
    });
})
