module.exports = {
  moduleFileExtensions: [
    "js",
    "json",
    "vue",
 ],

  transform: {
      "^.+\\.js$": "babel-jest",
      ".*\\.(vue)$": "@vue/vue3-jest"
  },

  transformIgnorePatterns: ['<rootDir>/node_modules/'],

  modulePaths: [
    '<rootDir>/src/',
    '<rootDir>/node_modules'
  ],

  moduleDirectories: [
    "src",
    "node_modules"
  ],

  moduleNameMapper: {
     "@/(.*)": "<rootDir>/test",
   },

  testEnvironment: 'jsdom',

  testEnvironmentOptions: {
    customExportConditions: ['node', 'node-addons']
  },

  preset: '@vue/cli-plugin-unit-jest/presets/no-babel',
  testMatch: [
    "**/test/**/*.spec.[jt]s?(x)",
  ]
}
