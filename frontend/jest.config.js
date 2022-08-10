module.exports = {
  collectCoverage: true,
  collectCoverageFrom: ["src/**/*.{js,vue}"],
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
    "@/(.*)": "<rootDir>/src/$1",
  },
  testEnvironment: 'jsdom',
  testEnvironmentOptions: {
    customExportConditions: ['node', 'node-addons']
  }
}
