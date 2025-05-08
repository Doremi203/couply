import { coverageConfigDefaults, defineConfig } from 'vitest/config';

export default defineConfig({
  // ...
  test: {
    coverage: {
      // 👇 Add this
      exclude: [
        ...coverageConfigDefaults.exclude,
        '**/.storybook/**',
        // 👇 This pattern must align with the `stories` property of your `.storybook/main.ts` config
        '**/*.stories.*',
        // 👇 This pattern must align with the output directory of `storybook build`
        '**/storybook-static/**',
      ],
    },
  },
});
