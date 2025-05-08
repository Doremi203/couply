# Storybook Coverage

This project is set up to track and report code coverage for Storybook stories. This helps ensure that components are properly tested and documented with stories.

## How It Works

The coverage setup uses:

- `@storybook/addon-coverage`: Instruments your code to track coverage
- `@storybook/test-runner`: Runs your stories as tests
- `nyc`: Generates coverage reports

## Running Coverage Tests

To run the coverage tests and generate a report, follow these steps:

### 1. Start Storybook

First, start the Storybook development server:

```bash
npm run storybook
```

### 2. Run Tests with Coverage

In a separate terminal, run the Storybook tests with coverage:

```bash
npm run test-storybook:coverage
```

This will run all your stories as tests and collect coverage data.

### 3. Generate Coverage Report

After the tests complete, generate an HTML coverage report:

```bash
npm run generate-coverage-report
```

### All-in-One Command

You can also run the entire process with a single command:

```bash
npm run storybook-coverage
```

This will run the tests with coverage and generate the report.

## Viewing the Coverage Report

After generating the report, open `coverage/storybook-html/index.html` in your browser to view the coverage results.

The report shows:

- Overall coverage percentage
- File-by-file breakdown
- Line-by-line coverage highlighting

## Understanding Coverage Metrics

The coverage report includes several metrics:

- **Statements**: Percentage of statements executed
- **Branches**: Percentage of control flow branches executed
- **Functions**: Percentage of functions called
- **Lines**: Percentage of executable lines executed

## Story Coverage Analysis

In addition to code coverage, you can also analyze which components have stories and which don't:

```bash
npm run analyze-story-coverage
```

This will generate:

- A console report showing components with and without stories
- An HTML report at `coverage/story-coverage-report.html`
- A JSON report at `coverage/story-coverage-report.json`

This analysis helps you identify components that need stories to improve your documentation and testing.

## Improving Coverage

To improve coverage:

1. Create stories for components that don't have them (use the story coverage analysis to identify these)
2. Ensure stories exercise different component states and props
3. Add interaction tests to cover component behavior

## Troubleshooting

### Missing Coverage Data

If you see an error like:

```
ENOENT: no such file or directory, scandir '/path/to/project/.nyc_output'
```

Or if you get a message about a "dummy report", it means the coverage data wasn't generated properly. Make sure:

1. Storybook is running when you run the tests
2. You run `npm run test-storybook:coverage` before `npm run generate-coverage-report`
3. The test runner completed successfully

The script will create a dummy report if no coverage data is found, which includes instructions on how to generate the real data.

### Other Common Issues

- **Missing modules**: Make sure all dependencies are installed
- **No stories found**: Check that components have properly exported stories
- **Coverage addon not working**: Verify that the coverage addon is properly configured in `.storybook/main.ts`

## Current Coverage Status

As of the last analysis, the project has:

- Total components: 66
- Components with stories: 63
- Components without stories: 3
- Coverage percentage: 95.45%

Components that need stories:

- src/features/filters/components/FiltersDrawer/styled/CustomSlider.tsx
- src/pages/LikesPage/components/MatchesSection/MatchesSection.tsx
- src/pages/HomePage/components/Page/HomePage.tsx
