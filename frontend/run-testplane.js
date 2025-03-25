#!/usr/bin/env node

import { spawn } from 'child_process';
import { fileURLToPath } from 'url';
import { dirname, resolve } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// Run testplane with the correct options
const args = process.argv.slice(2);
const testplaneProcess = spawn('node', [
  resolve(__dirname, 'node_modules/testplane/bin/testplane'),
  '--config',
  resolve(__dirname, 'testplane.config.cjs'),
  ...args
], {
  stdio: 'inherit',
  env: {
    ...process.env,
    TS_NODE_PROJECT: resolve(__dirname, 'testplane-tests/tsconfig.json')
  }
});

testplaneProcess.on('close', (code) => {
  process.exit(code);
});