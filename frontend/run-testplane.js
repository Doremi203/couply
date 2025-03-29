// #!/usr/bin/env node

// import { spawn } from 'child_process';
// import { fileURLToPath } from 'url';
// import { dirname, resolve } from 'path';
// import { existsSync } from 'fs';

// const __filename = fileURLToPath(import.meta.url);
// const __dirname = dirname(__filename);

// // Parse command line arguments
// const args = process.argv.slice(2);
// const guiMode = args.includes('--html-reporter-gui');
// const updateRefs = args.includes('--update-refs');

// // Filter out the --html-reporter-gui flag and add the correct flags
// const filteredArgs = args.filter(arg => arg !== '--html-reporter-gui');

// // Add GUI-specific flags
// if (guiMode) {
//   filteredArgs.push('--html-reporter-path', 'testplane-report');
//   filteredArgs.push('--html-reporter-gui');
//   console.log('Running Testplane in GUI mode...');
// } else if (updateRefs) {
//   console.log('Running Testplane with reference updates...');
// } else {
//   console.log('Running Testplane tests...');
// }

// // Check if ts-node is installed
// const tsNodePath = resolve(__dirname, 'node_modules/.bin/ts-node');
// const hasTsNode = existsSync(tsNodePath);

// // Run testplane with the correct options
// const testplaneProcess = spawn(hasTsNode ? tsNodePath : 'node', [
//   ...(hasTsNode ? ['--project', resolve(__dirname, 'testplane-tests/tsconfig.json')] : []),
//   resolve(__dirname, 'node_modules/testplane/bin/testplane'),
//   '--config',
//   resolve(__dirname, 'testplane.config.cjs'),
//   ...filteredArgs
// ], {
//   stdio: 'inherit',
//   env: {
//     ...process.env,
//     TS_NODE_PROJECT: resolve(__dirname, 'testplane-tests/tsconfig.json'),
//     NODE_OPTIONS: '--loader ts-node/esm'
//   }
// });

// testplaneProcess.on('close', (code) => {
//   process.exit(code);
// });