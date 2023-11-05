module.exports = {
    extends: [
        'eslint:recommended',
        'plugin:@typescript-eslint/recommended',
        // 'plugin:@typescript-eslint/recommended-react-checked',
        'prettier',
        // "eslint-plugin-import"
    ],
    parser: '@typescript-eslint/parser',
    parserOptions: {
        project: ['./tsconfig.json','./tsconfig.node.json', './packages/*/tsconfig.json'],
    },
    plugins: ['@typescript-eslint'],
    // preferConst: true,
    rules: {
        'no-console': 0,
    },
};
