# Build stage
FROM node:22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy the source code
COPY . .

# Generate Prisma Client
RUN npx prisma generate

# Build the application
RUN npm run build


# Runtime stage
FROM node:22-alpine

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install only production dependencies
RUN npm ci --omit=dev

# Copy built application from builder stage
COPY --from=builder /app/dist ./dist

# Copy generated Prisma Client
COPY --from=builder /app/generated ./generated

# Expose the port the app runs on
EXPOSE 3000

# Command to run the application
CMD ["node", "dist/src/main.js"]