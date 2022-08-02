module.exports = {
  "moduleFileExtensions": [
    "js",
    "json",
    "vue"
  ],
  transform: {
    '.*\\.js$':'babel-jest',
    ".*\\.(vue)$": "@vue/vue3-jest"
  },
  moduleNameMapper: {
    "@/(.*)": "<rootDir>/src/$1",
  },
  testEnvironment: 'jsdom',
  transformIgnorePatterns: ['<rootDir>/node_modules/'],
  testEnvironmentOptions: {
    customExportConditions: ["node", "node-addons"],
 },
}