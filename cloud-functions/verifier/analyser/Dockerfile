# ---- Builder Stage ----
FROM python:3.12-slim AS builder

# Install build tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy only requirements to leverage Docker cache
COPY requirements.txt .

# Create a virtual environment
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

# Install dependencies into the virtual environment
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application's code
COPY . .

# ---- Final Stage ----
FROM python:3.12-slim AS final

# Install runtime dependencies for OpenCV
RUN apt-get update && apt-get install -y --no-install-recommends \
    libgl1-mesa-glx \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the virtual environment from the builder stage
COPY --from=builder /opt/venv /opt/venv

# Copy application code from the builder stage
# Be specific about what you copy if not everything from /app in builder is needed
COPY --from=builder /app .

# Make port available
ENV PORT=8080
EXPOSE $PORT

# Define environment variable for the Functions Framework
ENV GOOGLE_FUNCTION_TARGET=handler
ENV GOOGLE_FUNCTION_SIGNATURE_TYPE=http

# Set PATH to include venv
ENV PATH="/opt/venv/bin:$PATH"

# Run functions-framework
CMD ["functions-framework", "--target=handler", "--host=0.0.0.0"]