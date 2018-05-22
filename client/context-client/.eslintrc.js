module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    // https://github.com/vuejs/eslint-plugin-vue#priority-a-essential-error-prevention
    'plugin:vue/recommended',
    'plugin:prettier/recommended'
  ],
  // required to lint *.vue files
  plugins: [
    'vue',
    'prettier'
  ],
  // add your custom rules here
  rules: {
    'prettier/prettier': [
      'error',
      {
        'singleQuote': true,
        'trailingComma': 'es5',
      }
    ],
    'vue/no-parsing-error': [2, {
      'x-invalid-end-tag': false
    }]
  }
}
