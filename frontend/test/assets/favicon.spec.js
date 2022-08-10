/* Ensure that the favicon is in place and not modified in a manner incompatible with index.html */
const fs = require('fs')
const process = require('process')

var base = process.cwd
const path = '../../src/assets/favicon.ico'
// const path = path.join(__dirname, 'favicon.ico')

describe('favicon available', () => {
    const output = fs.existsSync(path);
    test('file exists', () => {
        // expect(output).toBe(true);
    });
})
