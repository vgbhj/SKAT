# Build frontend
FROM node:20-alpine AS frontend-builder
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend .
RUN npm run build

# Build backend
FROM golang:1.25-alpine AS backend-builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

ENV GOPROXY=direct
ENV GOSUMDB=off

COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend .
# Create the dist directory explicitly
RUN mkdir -p frontend/dist
COPY --from=frontend-builder /app/dist ./frontend/dist
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=backend-builder /app/main .
# Create the directory structure
RUN mkdir -p frontend/dist
COPY --from=backend-builder /app/frontend/dist ./frontend/dist
COPY backend/conf ./conf

EXPOSE 8080
CMD ["./main"]