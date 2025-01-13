## Prerequisites

- Go 1.19+
- Node.js 16+
- pnpm (or npm/yarn)

## Getting Started

### Backend

1. **Install Go dependencies:**
   ```sh
   go mod tidy
   ```

2. **Run the application:**
   ```sh
   make start
   ```

### Frontend

1. **Navigate to the frontend directory:**
   ```sh
   cd frontend
   ```

2. **Install frontend dependencies:**
   ```sh
   pnpm install
   ```

3. **Start the development server:**
   ```sh
   pnpm run dev
   ```

### Building the Project

1. **Build the frontend:**
   ```sh
   make build_frontend
   ```

2. **Build the entire project:**
   ```sh
   make build
   ```

### Cleaning the Project

1. **Clean the build directories:**
   ```sh
   make clean
   ```

## Configuration

- **Port:** The application listens on port 9090 by default. You can change this by passing the `--port` flag when running the application.

## Development Mode

In development mode, the application proxies requests to the Vite development server running on `http://localhost:5173`.

## Production Mode

In production mode, the application serves the built frontend assets and uses a Vite manifest to load the correct CSS and JS files.

## License

This project is licensed under the MIT License.
```
MIT License

Copyright (c) 2025 Michał Mania

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```