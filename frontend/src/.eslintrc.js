// src/.eslintrc.js
module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: [
    "eslint:recommended",
    'plugin:react/recommended',
    'airbnb', // 确保这里正确配置
    'plugin:prettier/recommended', // 确保这里正确配置
    "plugin:jsx-a11y/recommended",
  ],
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  plugins: [
    'react',
    'prettier', // 确保这里正确配置
    'jsx-a11y',
  ],
  rules: {
    'prettier/prettier': 'error', // 确保这里正确配置
    'react/jsx-filename-extension': [1, { extensions: ['.js', '.jsx'] }],
    'import/no-extraneous-dependencies': ['error', { devDependencies: true }],
    'react/jsx-wrap-multilines': 0,
    'react/jsx-one-expression-per-line': 0,
    'react/jsx-curly-newline': 0,
    'react/jsx-indent': [2, 2],
    'react/jsx-indent-props': [2, 2],
    'react/jsx-closing-tag-location': 0,
    'react/jsx-max-props-per-line': 0,
    'react/jsx-first-prop-new-line': 0,
    'react/jsx-max-depth': 0,
    'react/jsx-no-bind': 0,
    'react/jsx-no-duplicate-props': 2,
    'react/jsx-no-literals': 0,
    'react/jsx-no-target-blank': 2,
    'react/jsx-no-undef': 2,
    'react/jsx-pascal-case': 2,
    'react/jsx-sort-default-props': 0,
    'react/jsx-sort-prop-types': 0,
    'react/jsx-sort-props': 0,
    'react/jsx-uses-react': 2,
    'react/jsx-uses-vars': 2,
    'react/jsx-wrap-multilines': 0,
    'react/no-array-index-key': 2,
    'react/no-danger': 2,
    'react/no-deprecated': 2,
    'react/no-direct-mutation-state': 2,
    'react/no-is-mounted': 2,
    'react/no-multi-comp': 2,
    'react/no-render-return-value': 2,
    'react/no-set-state': 0,
    'react/no-string-refs': 2,
    'react/no-unknown-property': 2,
    'react/prefer-es6-class': 2,
    'react/prefer-stateless-function': 2,
    'react/prop-types': 0,
    'react/react-in-jsx-scope': 2,
    'react/require-default-props': 0,
    'react/require-render-return': 2,
    'react/self-closing-comp': 2,
    'react/sort-comp': 2,
    'react/style-prop-object': 2,
    'react/void-dom-elements-no-children': 2,

    // 添加或修改此规则
    'react/function-component-definition': [
      2,
      {
        namedComponents: 'arrow-function',
        unnamedComponents: 'arrow-function',
      },
    ],

    // 确保 prettier/prettier 规则正确配置
    'prettier/prettier': [
      'error',
      {
        endOfLine: 'auto', // 或者设置为 'lf' 如果你想强制使用 LF 结尾
      },
    ],
  },
};
